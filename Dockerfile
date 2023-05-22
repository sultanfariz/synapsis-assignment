FROM golang:1.20.4-alpine AS builder

RUN apk add build-base

RUN mkdir /repo
ADD . /repo
WORKDIR /repo

RUN go clean --modcache
RUN go build ./app/main.go

FROM alpine:3.6

WORKDIR /root/

COPY --from=builder /repo/.env .
COPY --from=builder /repo/main .

EXPOSE 8080

CMD [ "./main" ]
