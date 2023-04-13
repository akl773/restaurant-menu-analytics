FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app

FROM alpine

COPY --from=builder /app/app /app/app

ENV FILEPATH /app/log.txt

ENTRYPOINT ["/app/app", "--logpath=${FILEPATH}"]