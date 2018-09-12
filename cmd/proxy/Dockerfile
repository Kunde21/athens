FROM golang:1.11-alpine AS build-env

WORKDIR $GOPATH/src/github.com/gomods/athens
COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 go build -mod=vendor -o /bin/app ./cmd/proxy
RUN ./scripts/create_default_config.sh
RUN cp $(which go) /bin/go && cp config.toml /config/proxy.toml

FROM alpine:3.8

RUN apk add --no-cache ca-certificates bzr git mercurial openssh-client subversion procps
ENV GO_ENV=production
EXPOSE 3000

COPY --from=build-env /bin/app /bin/athens
COPY --from=build-env /config/proxy.toml /config/proxy.toml
COPY --from=build-env /bin/go /bin/go
COPY --from=build-env /go/src/github.com/gomods/athens/cmd/proxy /go/src/github.com/gomods/athens/cmd/proxy

ENTRYPOINT /bin/athens -config_file=/config/proxy.toml
