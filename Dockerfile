FROM alpine

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

COPY ./bin/hashicups-traffic .

ENTRYPOINT ./hashicups-traffic