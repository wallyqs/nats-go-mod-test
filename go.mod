module github.com/wallyqs/nats-go-mod-test

go 1.16

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/nats-io/nats-server/v2 v2.2.1 // indirect
	github.com/nats-io/nats.go v1.10.1-0.20210401234754-e0d6a6af9882
)

replace github.com/nats-io/nats.go => github.com/wallyqs/nats.go v1.11.50
