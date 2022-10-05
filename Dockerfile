FROM golang:1.16

COPY ./src/ /opt/app
WORKDIR /opt/app

RUN chmod +x run.sh

CMD ["./run.sh"]