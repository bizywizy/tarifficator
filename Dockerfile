FROM golang:1.16-alpine as builder

WORKDIR /src
COPY ./ ./
RUN mkdir /out && CGO_ENABLED=0 go test -v ./... && go build -o /out/ ./cmd/server

FROM alpine:3.14
COPY --from=builder /out/server /app/server

EXPOSE 8000
CMD ["/app/server"]
