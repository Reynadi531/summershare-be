FROM golang:1.18-buster as builder

WORKDIR /app

COPY . .
RUN go mod download

RUN make build

FROM gcr.io/google.com/cloudsdktool/cloud-sdk:slim

WORKDIR /app

COPY --from=builder /app/bin/main /app/server
RUN chmod +x /app/server

CMD ["/app/server"]