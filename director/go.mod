module open-match.dev/open-match/tutorials/matchmaker101/director

go 1.14

require (
	github.com/bwmarrin/discordgo v0.21.1 // indirect
	github.com/challenge-league/nakama-go/commands v0.0.0-00010101000000-000000000000
	github.com/challenge-league/nakama-go/context v0.0.0-00010101000000-000000000000
	github.com/gofrs/uuid v3.3.0+incompatible // indirect
	github.com/golang/protobuf v1.4.1
	github.com/hako/durafmt v0.0.0-20200710122514-c0fb7b4da026 // indirect
	github.com/heroiclabs/nakama-common v1.5.1
	github.com/heroiclabs/nakama/v2/apigrpc v0.0.0-00010101000000-000000000000 // indirect
	github.com/micro/go-micro/v2 v2.9.1
	github.com/sqs/goreturns v0.0.0-20181028201513-538ac6014518 // indirect
	google.golang.org/grpc v1.27.1
	open-match.dev/open-match v1.1.0
)

replace (
	github.com/challenge-league/nakama-go/commands => ./nakama-go/commands
	github.com/challenge-league/nakama-go/context => ./nakama-go/context
	github.com/heroiclabs/nakama/v2/apigrpc => ./apigrpc
)
