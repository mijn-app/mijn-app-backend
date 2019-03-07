FROM frolvlad/alpine-glibc

RUN apk --no-cache add ca-certificates tzdata && update-ca-certificates

COPY . /

EXPOSE 80

STOPSIGNAL SIGTERM

CMD ["./main"]
