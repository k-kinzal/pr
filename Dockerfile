FROM buildpack-deps:scm

COPY pr /usr/local/bin/pr

ENTRYPOINT ["/usr/local/bin/pr"]
CMD ["--help"]