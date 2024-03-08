#! /bin/bash

protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/coreSwitch/*.proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/dataNode/*.proto
protoc --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. proto/nameNode/*.proto