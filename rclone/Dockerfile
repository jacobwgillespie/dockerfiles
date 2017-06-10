FROM alpine:3.5

ENV RCLONE_VERSION=current
ENV ARCH=amd64

RUN apk -U add ca-certificates fuse \
    && update-ca-certificates \
    && rm -rf /var/cache/apk/* \
    && cd /tmp \
    && wget -q http://downloads.rclone.org/rclone-${RCLONE_VERSION}-linux-${ARCH}.zip \
    && unzip /tmp/rclone-${RCLONE_VERSION}-linux-${ARCH}.zip \
    && mv /tmp/rclone-*-linux-${ARCH}/rclone /usr/bin \
    && rm -r /tmp/rclone* \
    && addgroup rclone \
    && adduser -h /config -s /bin/ash -G rclone -D rclone

USER rclone

VOLUME ["/config"]

ENTRYPOINT ["/usr/bin/rclone"]

CMD ["--version"]
