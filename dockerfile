FROM golang:latest as builder
RUN go get  -u -v  github.com/wudaoluo/downFile
WORKDIR /go/src/github.com/wudaoluo/downFile
RUN CGO_ENABLED=0 go build

FROM alpine:latest
MAINTAINER carlo "https://github.com/wudaoluo"
COPY --from=builder /go/src/github.com/wudaoluo/downFile/downFile .
RUN chmod +x downFile
EXPOSE 80
ENTRYPOINT ["./downFile"]