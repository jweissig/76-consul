FROM scratch
ADD web web
WORKDIR /
ENTRYPOINT ["/web"]
