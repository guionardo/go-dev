FROM golang:latest as build
WORKDIR /go/src/github.com/guionardo/go-dev
ENV CGO_ENABLED=0
COPY . .
ARG TARGETOS
ARG TARGETARCH
RUN GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /out/go-dev .

FROM scratch as bin
COPY --from=build /out/go-dev /


#FROM debian:buster as goenv
#RUN apt update && \
#    apt install -y curl
## Work inside the /tmp directory
#WORKDIR /tmp
#RUN curl https://storage.googleapis.com/golang/go1.16.2.linux-amd64.tar.gz -o go.tar.gz && \
#    tar -zxf go.tar.gz && \
#    rm -rf go.tar.gz && \
#    mv go /go
#ENV GOPATH /go
#ENV PATH $PATH:/go/bin:$GOPATH/bin
## If you enable this, then gcc is needed to debug your app
#ENV CGO_ENABLED 0
## TODO: Add other dependencies and stuff here

#FROM golang:latest
#
#WORKDIR /go/src/github.com/guionardo/go-dev
#
#COPY . .
#
#RUN go build -o /output .
#COPY go-dev /output
#RUN ls -la /output
