build:
	@protoc src/interfaces/rpcnode/rpcnode.proto --go_out=plugins=grpc:.

# 	go test || echo "error"
	go interfaces/restful || echo "error"

	./swag init \
	--exclude src/github.com \
	--exclude src/golang.org \
	--exclude src/google.golang.org \
	-g init.go \
	-o api -d src/interfaces/restful || exit 0
	redoc-cli bundle -o api/api.html api/swagger.json
# 	sshpass -p "iwdokserver" scp api.html iwdok@iwdok.team:/var/www/html/iwdok/mostik_dosc

install:
	go get github.com/hoisie/web
	go get google.golang.org/grpc
	go get github.com/hoisie/web
	go get google.golang.org/grpc/encoding/proto
	go get github.com/sirupsen/logrus
	go get github.com/patrickmn/go-cache
	go get github.com/DATA-DOG/go-txdb
	go get github.com/jinzhu/gorm
	go get github.com/go-sql-driver/mysql
	go get github.com/google/uuid