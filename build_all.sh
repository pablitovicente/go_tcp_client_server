#!/usr/bin/env bash
go build -o cmd/client/client cmd/client/client.go
go build -o cmd/multiple_clients/multiple_clients cmd/multiple_clients/multiple_clients.go
go build -o cmd/server/server cmd/server/server.go