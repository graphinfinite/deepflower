FROM golang:1.15-alpine3.12 AS builder

RUN go version

COPY ./ ./
WORKDIR /deepflower/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/deepflower ./cmd/deepflower/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /deepflower/.bin/deepflower .
COPY --from=builder /deepflower/config config/

EXPOSE 80 8787

CMD ["./deepflower"]