FROM alpine:latest

COPY build/main /main

COPY entrypoint.sh /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]