FROM golang:1.25-alpine AS build

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -ldflags="-w -s" ./...

FROM alpine:latest AS run

WORKDIR /app

COPY --from=build /app/smurfsatwork.org ./smurfsatwork.org
COPY --from=build /app/templates ./templates

EXPOSE 8080

CMD ["./smurfsatwork.org"]
