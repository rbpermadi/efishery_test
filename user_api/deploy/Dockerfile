FROM alpine:3

RUN apk update && apk add ca-certificates tzdata && rm -rf /var/cache/apk/*

COPY deploy/_output/user_api_docker /app/user_api

WORKDIR /app

EXPOSE 7171

CMD ["./user_api"]