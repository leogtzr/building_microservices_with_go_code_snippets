FROM alpine:latest

WORKDIR /

MAINTAINER jackson.nic@gmail.com

EXPOSE 8080

COPY ./server ./
ENTRYPOINT /server 
