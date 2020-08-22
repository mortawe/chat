FROM golang:1.13-buster AS build
WORKDIR /app
ADD . .

RUN make build

ENV DB_USER=user
ENV DB_PASS=pass
ENV DB_NAME=chat-db
ENV DB_ADDR=db:5432

CMD ["/app/bin/chat"]
