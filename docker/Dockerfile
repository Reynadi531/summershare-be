FROM golang:1.18 as builder

WORKDIR /app
COPY . .

RUN go mod download

ENV cgo_enabled=0

RUN make build

FROM gcr.io/distroless/static

COPY --from=builder /app/bin/main /main
CMD ["/main"]