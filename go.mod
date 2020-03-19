module github.com/242617/synapse-crawler

go 1.14

require (
	github.com/242617/synapse-core v0.0.0-20200319055251-42db5ba93c8e
	github.com/getsentry/sentry-go v0.5.1
	github.com/golang/protobuf v1.3.5 // indirect
	github.com/mattn/go-isatty v0.0.12
	github.com/pkg/errors v0.9.1
	github.com/rs/zerolog v1.18.0
	google.golang.org/grpc v1.28.0
	gopkg.in/yaml.v2 v2.2.8
)

replace github.com/242617/synapse-core => ../synapse-core