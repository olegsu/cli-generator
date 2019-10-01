FROM alpine:3.8

RUN apk add --update ca-certificates

COPY dist/cli-generator_linux_386/scout /usr/local/bin/

ENTRYPOINT ["cli-generator"]
CMD [ "--help" ]