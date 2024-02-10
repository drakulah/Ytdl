@echo off
go build -ldflags "-s -w -X main.version=1.0.0"