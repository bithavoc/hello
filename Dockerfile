FROM golang:1.6-alpine

MAINTAINER bithavoc

ENV BREAKDECK_ORG $GOPATH/src/github.com/bithavoc/
ENV BREAKDECK_PATH $BREAKDECK_ORG/hello

COPY . /hello
RUN mkdir -p $BREAKDECK_ORG
RUN cp -R /hello $BREAKDECK_PATH

RUN cd $BREAKDECK_PATH && go install

EXPOSE 8080
CMD hello
