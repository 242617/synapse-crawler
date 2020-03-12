FROM 242617/go-builder:1.0.0 AS builder

ARG PROJECT
ARG APPLICATION
ARG ENVIRONMENT
ARG VERSION

ENV PROJECT=${PROJECT}
ENV APPLICATION=${APPLICATION}
ENV ENVIRONMENT=${ENVIRONMENT}
ENV VERSION=${VERSION}

WORKDIR /root
COPY . .
RUN go build \
        -o build/crawler \
        -ldflags "\
        -X '${PROJECT}/version.Application=${APPLICATION}'\
        -X '${PROJECT}/version.Environment=${ENVIRONMENT}'\
        -X '${PROJECT}/version.Version=${VERSION}'\
        "\
        cmd/crawler/main.go

FROM alpine:3.10.2

WORKDIR /usr/local
COPY --from=builder /root/build/crawler .
WORKDIR /etc/crawler
COPY build/config.yaml .

CMD /usr/local/crawler --config /etc/crawler/config.yaml