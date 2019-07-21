#cd到创建服务的目录下，执行该语句
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
