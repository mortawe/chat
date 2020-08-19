FROM golang:1.13-buster AS build
WORKDIR /app
ADD . .

RUN make build

#FROM debian:buster-slim
#
#WORKDIR /app
#
#COPY --from=build /app/bin/chat .

#RUN chmod +x /app/chat
CMD ["/app/bin/chat"]
