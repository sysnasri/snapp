# TODO v1.0.0
FROM golang as builder 
# uncomment if you are located in Iran :(
#ENV HTTPS_PROXY=http://178.22.121.13:8080
#ENV HTTP_PROXY=http://178.22.121.13:8080
ENV APP_USER app
ENV APP_HOME /go/src/v1
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
COPY apps/v1.0.0/. .
RUN go mod download
RUN go mod verify
RUN go build -o v1
FROM ubuntu
RUN apt update -y
RUN apt install net-tools wget iputils-ping curl -y
ENV APP_USER app
ENV APP_HOME /go/src/v1
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME
COPY --chown=0:0 --from=builder $APP_HOME/v1 $APP_HOME
#USER $APP_USER
EXPOSE 3001
CMD ["./v1"]
