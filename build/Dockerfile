FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/scout_linux_386/scout /usr/local/bin/

ENTRYPOINT ["cli-generator"]
CMD [ "--help" ]