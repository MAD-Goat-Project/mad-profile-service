FROM golang:1.20-alpine as build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/profile

FROM alpine:latest as production

WORKDIR /app
COPY --from=build /app/profile .

ENV GIN_MODE=release

EXPOSE 8080

CMD ["./profile"]