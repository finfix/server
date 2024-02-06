module tgBot

go 1.20

replace (
	logger => ../logger
	pkg => ../pkg
)

require (
	github.com/caarlos0/env/v7 v7.1.0
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
	gopkg.in/telebot.v3 v3.2.1
	logger v0.0.0-00010101000000-000000000000
	pkg v0.0.0-20231025211159-75bb09ada85b
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
)