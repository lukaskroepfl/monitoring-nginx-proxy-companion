FROM alpine:latest

ENV PROXY_CONTAINER_NAME nginx
ENV INFLUX_URL http://172.18.0.2:8086

COPY GeoLite2-City.mmdb /
COPY build/main /main

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]