
FROM centos:centos7

ADD wsserver /home/wsserver
ADD cert /home/cert

WORKDIR /home/

EXPOSE 3333
EXPOSE 3334


ENTRYPOINT /home/wsserver

#docker run -tdi -p3333:3333 -p3334:3334  --name="wsserver1.0"  wsserver:v1.0
