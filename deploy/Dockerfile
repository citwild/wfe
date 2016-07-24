FROM alpine:latest
MAINTAINER sochas-ras@u.washington.edu

ADD wfe /usr/local/bin/wfe

ENTRYPOINT ["wfe"]
CMD ["serve", "--https-addr=:443"]

EXPOSE 443