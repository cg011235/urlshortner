FROM alpine:latest
WORKDIR /app
COPY  bin/url_shortner_main_for_docker /app/url_shortner_main
EXPOSE 8080
CMD ["./url_shortner_main"]

