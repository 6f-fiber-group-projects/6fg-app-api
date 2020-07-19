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
COPY ./src/migration /app/migration

CMD ["app/main"]