FROM golang:1.19-bullseye

WORKDIR /app

COPY . go.mod ./
COPY . go.sum ./
COPY . Makefile ./

COPY . ./

RUN make deps
RUN make build-linux

EXPOSE 3000

CMD ["make", "start", "p=3000"]

