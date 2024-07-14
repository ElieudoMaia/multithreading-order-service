FROM golang:1.22.0
WORKDIR /app
ENTRYPOINT ["tail", "-f", "/dev/null"]