// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"sync"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/heroiclabs/nakama-common/api"

	"github.com/golang/protobuf/ptypes/any"
	"google.golang.org/grpc"
	"open-match.dev/open-match/pkg/pb"
)

// The Director in this tutorial continously polls Open Match for the Match
// Profiles and makes random assignments for the Tickets in the returned matches.

const (
	// The endpoint for the Open Match Backend service.
	omBackendEndpoint = "open-match-backend.open-match.svc.cluster.local:50505"
	// The Host and Port for the Match Function service endpoint.
	functionHostName       = "matchfunction.dataleague.svc.cluster.local"
	functionPort     int32 = 50502
)

func Marshal(v interface{}) []byte {
	resultJSON, err := json.Marshal(v)
	if err != nil {
		log.Infof("Error %+v", err)
	}
	return resultJSON
}

func MarshalIndent(v interface{}) string {
	resultJSON, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		log.Error(err)
	}
	return string(resultJSON)
}

func main() {
	// Connect to Open Match Backend.
	conn, err := grpc.Dial(omBackendEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Infof("Error %+v", err)
		log.Fatalf("Failed to connect to Open Match Backend, got %s", err.Error())
	}

	defer conn.Close()
	be := pb.NewBackendServiceClient(conn)

	// Generate the profiles to fetch matches for.
	profiles := generateProfiles()
	log.Infof("Fetching matches for %v profiles", len(profiles))

	for range time.Tick(time.Second * 3) {
		// Fetch matches for each profile and make random assignments for Tickets in
		// the matches returned.
		var wg sync.WaitGroup
		for _, p := range profiles {
			wg.Add(1)
			go func(wg *sync.WaitGroup, p *pb.MatchProfile) {
				defer wg.Done()
				matches, err := fetch(be, p)
				if err != nil {
					log.Infof("Error %+v", err)
					log.Infof("Failed to fetch matches for profile %v, got %s", p.GetName(), err.Error())
					return
				}

				log.Infof("Generated %v matches for profile %v", len(matches), p.GetName())
				if err := assign(be, matches); err != nil {
					log.Error(err)
					log.Infof("Failed to assign servers to matches, got %s", err.Error())
					return
				}
			}(&wg, p)
		}

		wg.Wait()
	}
}

func fetch(be pb.BackendServiceClient, p *pb.MatchProfile) ([]*pb.Match, error) {
	req := &pb.FetchMatchesRequest{
		Config: &pb.FunctionConfig{
			Host: functionHostName,
			Port: functionPort,
			Type: pb.FunctionConfig_GRPC,
		},
		Profile: p,
	}

	stream, err := be.FetchMatches(context.Background(), req)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	var result []*pb.Match
	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return nil, err
		}

		result = append(result, resp.GetMatch())
	}

	return result, nil
}

func assign(be pb.BackendServiceClient, matches []*pb.Match) error {
	for _, match := range matches {
		ticketIDs := []string{}
		for _, t := range match.GetTickets() {
			ticketIDs = append(ticketIDs, t.Id)
		}
		conn := fmt.Sprintf("discord match %s", match.MatchId)

		req := &pb.AssignTicketsRequest{
			Assignments: []*pb.AssignmentGroup{
				{
					TicketIds: ticketIDs,
					Assignment: &pb.Assignment{
						Connection: conn,
						Extensions: map[string]*any.Any{},
					},
				},
			},
		}

		if _, err := be.AssignTickets(context.Background(), req); err != nil {
			log.Fatalf("Failed to assign servers to matches, got %s	", err.Error())
			return fmt.Errorf("AssignTickets failed for match %v, got %w", match.MatchId, err)
		}

		nakamaCtx := NewNakamaContext()
		result, err := nakamaCtx.Client.RpcFunc(nakamaCtx.Ctx, &api.Rpc{Id: "MatchCreate", Payload: string(Marshal(match))})
		if err != nil {
			log.Error(err)
			return err
		}
		log.Infof("%+v\n", MarshalIndent(result))
	}

	return nil
}
