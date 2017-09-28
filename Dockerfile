FROM ysitd/glide

ADD . /go/src/github.com/ysitd-cloud/status

WORKDIR /go/src/github.com/ysitd-cloud/status

RUN glide install -v && \
    go install

ENV PORT 80
EXPOSE 80

CMD ["status"]
