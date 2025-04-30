#!/bin/sh
reflex -r '\.go$' -s -- sh -c 'go run /app/cmd/server/main.go'