FROM golang:alpine AS builder
LABEL maintainer="hjs7747@khu.ac.kr"
LABEL version="0.0.1"
LABEL description="CPFD data collect server"

ENV GOPATH /go
ENV PATH $PATH:/go/bin:$GOPATH/bin
ENV GOBIN=$GOPATH/bin
ENV PATH=$PATH:$GOROOT:$GOPATH:$GOBIN

RUN apk --no-cache add tzdata

EXPOSE 8080

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN TZ=Asia/Seoul CGO_ENABLED=0 GOOS=linux go build -o app cmd/cpfd/*.go

FROM scratch AS production

COPY --from=builder /app .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Seoul

ENTRYPOINT ["/app"]