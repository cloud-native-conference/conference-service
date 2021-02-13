FROM golang:1.15.8-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine:3.13.1
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=build-env /src/goapp /app/

EXPOSE 5000

ENV MONGO_DB_HOST=
ENV MONGO_DB_PORT=
ENV MONGO_DB_USER=
ENV MONGO_DB_PASSWORD=
ENV MONGO_DB_PARAMETERS=

ENTRYPOINT ./goapp
