//go:generate protoc -I checker/ checker.proto --go_out=plugins=grpc:checker
package deckard
