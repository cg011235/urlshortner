FROM golang:1.21 as builder
WORKDIR /app
COPY .  .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o url_shortner_main ./src/cmd/url_shortner_main/

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/url_shortner_main .
EXPOSE 8080
CMD ["./url_shortner_main"]
