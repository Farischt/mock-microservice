FROM golang:1.19-bullseye

WORKDIR /app

COPY . go.mod ./
COPY . go.sum ./
COPY . Makefile ./

COPY . ./

RUN make deps
RUN make build-linux

EXPOSE 9090

CMD ["make", "dev", "p=9090"]

