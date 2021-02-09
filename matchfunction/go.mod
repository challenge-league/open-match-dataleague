module open-match.dev/open-match/tutorials/matchmaker101/matchfunction

go 1.14

require (
	github.com/challenge-league/nakama-go/commands v0.0.0-00010101000000-000000000000
	github.com/gofrs/uuid v3.3.0+incompatible
	github.com/golang/protobuf v1.4.1
	github.com/hako/durafmt v0.0.0-20200710122514-c0fb7b4da026 // indirect
	github.com/heroiclabs/nakama/v2/apigrpc v0.0.0-00010101000000-000000000000 // indirect
	google.golang.org/grpc v1.27.1
	google.golang.org/protobuf v1.25.0
	open-match.dev/open-match v1.1.0
)

replace (
	github.com/challenge-league/nakama-go/commands => ./nakama-go/commands
	github.com/challenge-league/nakama-go/context => ./nakama-go/context
	github.com/heroiclabs/nakama/v2/apigrpc => ./apigrpc
)
