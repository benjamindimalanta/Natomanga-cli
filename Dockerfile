FROM golang

WORKDIR /go/src/natomanga-cli
COPY . .

RUN go get

CMD ["go","run","."]
