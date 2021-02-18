FROM golang:1.12-alpine
ENV HELLO_MESSAGE=hello
ENV GO111MODULE=on
WORKDIR /go/src/app
COPY . .
RUN apk add --update git
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...
CMD go run main.go
EXPOSE 8080
