FROM ubuntu:20.04

RUN apt-get update && \
    apt-get install -y s3fs fuse kmod && \
    rm -rf /var/lib/apt/lists/*

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

COPY .env /etc/environment

ENTRYPOINT ["/entrypoint.sh"]
