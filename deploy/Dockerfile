FROM alpine:latest
MAINTAINER sochas-ras@u.washington.edu

COPY wfe /usr/local/bin/wfe

ENTRYPOINT ["wfe"]
CMD ["serve", "--http-addr=:80", "--https-addr=:443"]

EXPOSE 80 443