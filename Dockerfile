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

ENV PORT=8080
ENV DB_CLUSTER1=192.168.0.14
ENV DB_USERNAME=cassandra
ENV DB_PASSWORD=cassandra
ENV DB_KEYSPACE=em
ENV EMAIL_HOST=smtp.gmail.com
ENV EMAIL_PORT=587
ENV EMAIL_LOGIN=kdulembabitbucket@gmail.com
ENV EMAIL_PASSWORD=Mas3lko#12

CMD ["app"]