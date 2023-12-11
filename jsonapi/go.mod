module jsonapi

go 1.20

replace (
	auth => ../auth
	core => ../core
	logger => ../logger
	pkg => ../pkg
)

require (
	auth v0.0.0-00010101000000-000000000000
	core v0.0.0-00010101000000-000000000000
	github.com/caarlos0/env/v7 v7.1.0
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/schema v1.2.0
	github.com/swaggo/http-swagger v1.3.4
	google.golang.org/grpc v1.58.2
	google.golang.org/protobuf v1.31.0
	logger v0.0.0-00010101000000-000000000000
	pkg v0.0.0-20231025211159-75bb09ada85b
)

require (
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.6 // indirect
	github.com/go-openapi/swag v0.19.15 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/mailru/easyjson v0.7.6 // indirect
	github.com/swaggo/files v0.0.0-20220610200504-28940afbdbfe // indirect
	github.com/swaggo/swag v1.16.1 // indirect
	github.com/tidwall/gjson v1.14.2 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	github.com/tidwall/pretty v1.2.0 // indirect
	github.com/tidwall/sjson v1.2.5 // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
	golang.org/x/tools v0.7.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231002182017-d307bd883b97 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)
