# TODO v2.0.0
FROM golang as builder 
ENV HTTPS_PROXY=http://178.22.121.13:8080
ENV HTTP_PROXY=http://178.22.121.13:8080
ENV APP_USER app
ENV APP_HOME /go/src/v2
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY . .
RUN go mod download
RUN go mod verify
RUN go build -o v2
FROM ubuntu 
ENV APP_USER app
ENV APP_HOME /go/src/v2
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/v2 $APP_HOME
USER $APP_USER
EXPOSE 3000
CMD ["./v2"]