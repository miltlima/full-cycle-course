FROM golang:alpine as build

COPY . .

RUN go build message.go

FROM scratch

COPY --from=build /go/message ./message
ENTRYPOINT ["./message"]
