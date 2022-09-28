FROM registry.semaphoreci.com/golang:1.18 as builder

ENV APP_HOME /go/src/cadastropalestrante

WORKDIR "$APP_HOME"
COPY . .

RUN go mod download
RUN go mod verify
RUN go build -o cadastropalestrante
 

FROM registry.semaphoreci.com/golang:1.18

ENV APP_HOME /go/src/cadastropalestrante
RUN mkdir -p "$APP_HOME"
WORKDIR "$APP_HOME"


COPY --from=builder "$APP_HOME"/cadastropalestrante $APP_HOME

EXPOSE 8080
CMD ["./cadastropalestrante"]