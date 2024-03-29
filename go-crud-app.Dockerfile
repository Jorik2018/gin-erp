FROM golang:latest

#setup working directory
WORKDIR /go/src/github.com/Jorik2018/gin-erp

#copy files
COPY dao ./dao
COPY db ./db
COPY handler ./handler
COPY repository ./repository
COPY server.go ./server.go

#install dependancies and build executable
RUN go get -d -v ./... && go install -v ./...

#copy config files
COPY config.prod.json ./config.json

#run executable
CMD ["gin-erp"]

#expose port
EXPOSE 8080