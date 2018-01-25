
start:
	go build -o server *.go && ./server

daemon_start:
	go build -o server *.go && ./server &

deps:
	go get -u "github.com/kataras/iris"
	go get -u "github.com/jinzhu/gorm"
	go get -u "github.com/jinzhu/gorm/dialects/mysql"

