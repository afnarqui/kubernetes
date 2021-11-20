FROM golang:1.16-buster AS build

WORKDIR /app

COPY ./src/main.go ./

RUN CGO_ENABLED=0 GOOS=linux GOPROXY=https://proxy.golang.org go build -o app ./main.go

##
## Deploy
##
FROM alpine:latest

WORKDIR /app

COPY --from=build /app/app .

CMD ["./app"]