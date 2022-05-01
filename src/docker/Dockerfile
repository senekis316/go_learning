FROM golang:1.16-buster as build
WORKDIR /app
COPY ./ /app/ 
RUN go build -o /go-http-server

FROM alpine:3.10
WORKDIR /
COPY --from=build /go-http-server /go-http-server
ENV GOPROXY https://goproxy.cn
ENV GOHTTPSERVERPORT :8080
EXPOSE 8080
ENTRYPOINT ["/go-http-server"]
