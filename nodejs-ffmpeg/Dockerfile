FROM node:8-alpine

RUN \
    apk update && \
    apk add git ffmpeg && \
    rm -rf /var/cache/apk/*

CMD ["node"]
