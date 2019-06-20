FROM fedora:30

RUN dnf -y update && dnf -y install golang curl python3-requests python3-yaml && dnf clean all

RUN mkdir -p /go && chmod -R 777 /go

ENV GOPATH=/go

WORKDIR /go

RUN mkdir /go/app 
ADD . /go/app/
WORKDIR /go/app 
RUN go get gopkg.in/src-d/go-git.v4
RUN go build -o main .
CMD ["./main"]
EXPOSE 9666/tcp