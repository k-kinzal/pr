FROM golang:1.13 as builder

ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src/github.com/k-kinzal/pr
COPY . /go/src/github.com/k-kinzal/pr

RUN make build

FROM buildpack-deps:scm

COPY --from=builder /go/src/github.com/k-kinzal/pr/dist/linux-amd64/pr /usr/local/bin/pr

ENTRYPOINT ["/usr/local/bin/pr"]
CMD ["--help"]