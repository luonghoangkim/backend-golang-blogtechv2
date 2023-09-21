FROM alpine:3.18

RUN apk update && apk add git go

ENV GO111MODULE=on

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH/src/backend-blogtechv2

COPY . .

RUN go mod init backend-blogtechv2

# Add go get commands for missing dependencies
RUN go get github.com/go-playground/locales/en
RUN go get github.com/go-playground/universal-translator
RUN go get github.com/go-playground/validator/v10
RUN go get github.com/go-playground/validator/v10/translations/en
RUN go get github.com/golang-jwt/jwt
RUN go get github.com/google/uuid
RUN go get github.com/jmoiron/sqlx
RUN go get github.com/labstack/echo/v4
RUN go get github.com/labstack/echo/v4/middleware
RUN go get github.com/labstack/gommon/log
RUN go get github.com/lestrrat/go-file-rotatelogs
RUN go get github.com/lib/pq
RUN go get github.com/pkg/errors
RUN go get github.com/rifflock/lfshook
RUN go get github.com/sirupsen/logrus
RUN go get golang.org/x/crypto/bcrypt

WORKDIR cmd/pro
RUN go build -o app

ENTRYPOINT ["./app"]

EXPOSE 3000
