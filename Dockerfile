FROM golang:1.8.3-onbuild
ENV ONMYOJI_EVENT_BOT_TOKEN token


RUN go get -u github.com/syhlion/onmyoji_event_bot
CMD mkdir /onmyoji_bot && mkdir /onmyoji_bot/db && cp /go/bin/onmyoji_event_bot /onmyoji_bot
WORKDIR /onmyoji_bot
VOLUME ["db"]

ENTRYPOINT ["onmyoji_event_bot"]

