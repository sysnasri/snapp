FROM haproxy:2.3
RUN apt update -y
RUN apt install net-tools wget iputils-ping curl -y
ENV HA_USER haproxy
ENV HA_HOME /usr/local/etc/haproxy
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
USER $HA_USER
COPY haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
EXPOSE 3000