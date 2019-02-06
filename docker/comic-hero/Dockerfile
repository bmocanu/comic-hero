FROM alpine:3.8
MAINTAINER you@mail.com

ENV PROCESS_USER cmchero
ENV PROCESS_UID 7000
ENV PROCESS_GROUP cmchero
ENV PROCESS_GID 8000

RUN set -x \
    && addgroup $PROCESS_USER \
    && adduser -u $PROCESS_UID -G $PROCESS_USER -H -D -s /bin/bash $PROCESS_USER \
    && addgroup -g $PROCESS_GID $PROCESS_GROUP

ADD comic-hero /
RUN chown $PROCESS_USER:$PROCESS_GROUP /comic-hero
RUN chmod +x /comic-hero

ENTRYPOINT ["/comic-hero", "-config-dir=/config"]
