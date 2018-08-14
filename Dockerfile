FROM golang as builder
ENV WORKDIR /go/src/github.com/threecommaio/helloworld

WORKDIR ${WORKDIR}

COPY . ${WORKDIR}

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app

FROM alpine:latest
ENV WORKDIR /go/src/github.com/threecommaio/helloworld
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder ${WORKDIR}/app /go/bin/app
ENTRYPOINT [ "/go/bin/app" ]
