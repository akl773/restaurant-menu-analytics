FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go build -o app

FROM alpine

WORKDIR /app

COPY --from=builder /app/app /app/app

ENTRYPOINT ["/app/app"]


## docker build -t restaurantapp .
## docker run -v /Users/infra/go/src/test-files/restaurant-menu-analytics/log.txt:/app/log.txt restaurantapp