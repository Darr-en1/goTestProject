# -I=. 可加可不加，暂时不知道干嘛的
# protoc -I=. --go_out=plugins=grpc,paths=source_relative:gen/go trip.proto

# 需要下载一下内容,保证 $GOPATH/bin 下存在可执行文件
# go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
# go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
# go install  github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
protoc --go_out=gen/go --go_opt=paths=source_relative --go-grpc_out=gen/go --go-grpc_opt=paths=source_relative trip.proto
protoc --grpc-gateway_out=paths=source_relative,grpc_api_configuration=trip.yaml:gen/go  trip.proto