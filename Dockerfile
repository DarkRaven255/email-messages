#
# Dockerfile
#
FROM golang:alpine as builder

WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod download

COPY cmd /app/cmd
COPY app /app/app
COPY config /app/config
COPY delivery /app/delivery
COPY repository /app/repository
COPY service /app/service
COPY domain /app/domain
COPY utils /app/utils

RUN go build -o bin/app cmd/emailmessages/main.go

FROM alpine
COPY --from=builder /app/bin/app /bin/
WORKDIR /app

CMD ["app"]