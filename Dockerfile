FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

WORKDIR /app/
ADD ./app /app/
ADD ./config_stg.yaml /app/config.yaml
ENTRYPOINT ["./app"] 