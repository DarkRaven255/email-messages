#
# Makefile
#

build:
  cd cmd/emailmessages && go build -o ../../builds/emailmessages

run:
  go run cmd/emailmessages/main.go

test:
  go test ./...