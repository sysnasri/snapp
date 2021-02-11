FROM haproxy:2.3
RUN apt update -y
RUN apt install net-tools wget iputils-ping curl -y
ENV HA_HOME /usr/local/etc/haproxy
COPY haproxy/haproxy.cfg /usr/local/etc/haproxy/haproxy.cfg
EXPOSE 3000