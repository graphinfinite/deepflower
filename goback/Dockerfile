FROM golang:alpine AS builder

RUN go version

WORKDIR /temp

COPY . .
RUN go mod download

RUN GOOS=linux go build -o ./.bin/deepflower ./cmd/deepflower/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder ./temp/.bin/deepflower .
COPY --from=builder ./temp/config config/

EXPOSE 8787 8787

CMD ["./deepflower"]