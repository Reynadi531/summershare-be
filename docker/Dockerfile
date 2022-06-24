FROM golang:1.18 as builder

WORKDIR /go/src/app
COPY . .

RUN go mod download

ENV cgo_enabled=0

RUN make build

FROM gcr.io/distroless/static-debian11

COPY --from=build /go/bin/main /
CMD ["/main"]