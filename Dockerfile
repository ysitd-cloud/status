FROM ysitd/glide as builder

ADD . /go/src/github.com/ysitd-cloud/status

WORKDIR /go/src/github.com/ysitd-cloud/status

RUN glide install -v && \
    go build -v

FROM alpine:3.6
COPY --from=builder /go/src/github.com/ysitd-cloud/status/status /

ENV PORT 80
EXPOSE 80

CMD ["/status"]
