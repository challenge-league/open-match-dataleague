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
	nakamaCommands "github.com/challenge-league/nakama-go/commands"
	"open-match.dev/open-match/pkg/pb"
)

// generateProfiles generates test profiles for the matchmaker101 tutorial.
func generateProfiles() []*pb.MatchProfile {
	var profiles []*pb.MatchProfile
	for _, mode := range nakamaCommands.MATCH_MAKER_MODES {
		for duration := 1; duration <= 48; duration++ {
			profiles = append(profiles, &pb.MatchProfile{
				Name: mode,
				Pools: []*pb.Pool{
					{
						Name: mode,
						TagPresentFilters: []*pb.TagPresentFilter{
							{
								Tag: mode,
							},
						},
						DoubleRangeFilters: []*pb.DoubleRangeFilter{
							&pb.DoubleRangeFilter{
								DoubleArg: nakamaCommands.SEARCH_MIN_DURATION,
								Min:       float64(duration),
								Max:       float64(duration),
							},
							&pb.DoubleRangeFilter{
								DoubleArg: nakamaCommands.SEARCH_MAX_DURATION,
								Min:       float64(duration),
								Max:       float64(duration),
							},
						},
					},
				},
			},
			)
		}
	}
	return profiles
}
