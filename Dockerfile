FROM golang:1.8.3-onbuild
ENV ONMYOJI_EVENT_BOT_TOKEN token


RUN mkdir /onmyoji_bot && mkdir /onmyoji_bot/db
WORKDIR /onmyoji_bot
VOLUME ["/onmyoji_bot/db"]

ENTRYPOINT ["app"]

