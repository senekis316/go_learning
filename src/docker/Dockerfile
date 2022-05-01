FROM golang:1.16-buster as build
WORKDIR /app
COPY ./ /app/ 
RUN go build -o /go-http-server
# RUN CGO_ENABLED=0 GODS=linux GOARCH=amd64 go build - ldflags "-s -w" -o /go-http-server

# FROM gcr.io/distroless/base-debian10
FROM alpine:3.10
WORKDIR /
COPY --from=build /go-http-server /go-http-server
ENV GOPROXY https://goproxy.cn
ENV GOHTTPSERVERPORT :8080
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/go-http-server"]
