module auth

go 1.20

replace (
	core => ../core
	logger => ../logger
	pkg => ../pkg
)

require (
	core v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env/v7 v7.1.0
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
	logger v0.0.0-00010101000000-000000000000
	pkg v0.0.0-20231025211159-75bb09ada85b
)

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.4.3 // indirect
	github.com/jmoiron/sqlx v1.3.5 // indirect
	golang.org/x/crypto v0.13.0 // indirect
	golang.org/x/net v0.15.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
)
