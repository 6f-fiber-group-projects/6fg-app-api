FROM golang:latest as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /go/src/github.com/6f-fiber-group-projects/6fg-app-api
COPY ./src .
RUN go build -o main *.go

# runtime image
FROM alpine
COPY --from=builder /go/src/github.com/6f-fiber-group-projects/6fg-app-api/main /app/main
COPY --from=builder /go/src/github.com/6f-fiber-group-projects/6fg-app-api/src/migration /app/migration
COPY --from=builder /go/bin/goose /usr/local/bin/goose

CMD ["app/main"]