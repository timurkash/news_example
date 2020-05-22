FROM golang:1.13 as builder

RUN mkdir -p /news/

WORKDIR /news

COPY . .

RUN GIT_COMMIT=$(git rev-list -1 HEAD) && \
    CGO_ENABLED=0 GOOS=linux go build -mod=vendor -ldflags "-s -w \
    -X gitlab.mcsolutions.ru/find-psy/back/admins/news/pkg/version.REVISION=${GIT_COMMIT}" \
    -a -o bin/news cmd/news/*

FROM alpine:latest

RUN addgroup -S app \
    && adduser -S -g app app \
    && apk --no-cache add \
    curl openssl netcat-openbsd mc

WORKDIR /home/app

COPY --from=builder /news/bin/news .
COPY --from=builder /news/docs/swagger.yaml ./docs/swagger.yaml

RUN chown -R app:app ./

USER app

EXPOSE 3008

CMD ["./news"]
