# gin-erp ![alt text][build_status]

an example rest api in golang and mysql

## prerequisites

- go
- mysql

## next steps

- clone this repository
- restore database `mysql -u <mysql-username> -p < db.sql`
- run program `go run server/server.go`

[build_status]: https://travis-ci.org/Jorik2018/gin-erp.svg?branch=master "Travis Build Status"

go mod init github.com/Jorik2018/gin-erp
go get -u github.com/gin-contrib/gzip

go mod tidy

mkdir gin-erp
cd gin-erp
go mod init isobit/erp
go get github.com/gin-gonic/gin


https://github.com/lakshanwd/go-crud.git