FROM golang:alpine

RUN apk add vim

RUN mkdir -p /go/src/memestring

WORKDIR /go/src/memestring

COPY . .

RUN go install

ENTRYPOINT ["memestring"]

CMD ["big", "fat", "titties"]
