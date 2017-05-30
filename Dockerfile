FROM alpine:latest

ENV PROXY_CONTAINER_NAME nginx
ENV INFLUX_DB_NAME monitoring

COPY GeoLite2-City.mmdb /
COPY build/main /main

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]