FROM golang

WORKDIR /app

COPY . .

RUN go build

EXPOSE 8080

ENV DROP_PATH=/tmp

ENTRYPOINT [ "sh", "-c", "/app/file-upload" ]