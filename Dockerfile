# syntax=docker/dockerfile:1

FROM golang:1.19-alpine
WORKDIR /app
COPY main.go ./
COPY *.tmpl ./
COPY types ./types
RUN go mod init poketype
RUN go mod tidy
RUN go mod download
RUN go build -o ./docker-poketype
EXPOSE 8088
CMD [ "/app//docker-poketype" ]
