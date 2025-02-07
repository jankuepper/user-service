FROM golang:1.22.5

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN make build

CMD ["make", "run"]