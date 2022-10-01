FROM golang:1.19.1-alpine3.16 AS GO_BUILD
COPY . /api
WORKDIR /api
RUN touch /.env
RUN printenv > /.env
RUN go build -o /go/bin/api

FROM alpine:3.10
WORKDIR app
COPY --from=GO_BUILD /go/bin/api/ ./
EXPOSE 8080
CMD ./api
