FROM golang:latest


# Declare required environment variables
ENV GOPATH=/go

# Get the required Go packages
RUN go get -u github.com/gorilla/mux



# Transpile and install the client-side application code
#RUN go get -v ./..


# Build and install the server-side application code
WORKDIR /go/src/golangapp
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...


# Specify the entrypoint
#ENTRYPOINT /go/src/svcrm/main

# Expose port 8080 of the container
#EXPOSE 8080

CMD [ "main","run" ]