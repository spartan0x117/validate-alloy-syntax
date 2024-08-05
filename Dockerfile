FROM golang:1.22.5-alpine3.20 as builder

WORKDIR /app

COPY ./ ./

RUN go build -o /bin/app main.go

FROM alpine:3.20 as app

COPY --from=builder /bin/app /bin/app

ENTRYPOINT ["app"]
