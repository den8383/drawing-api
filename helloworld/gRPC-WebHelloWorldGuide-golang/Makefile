protoc-js:
	protoc -I=. proto/helloworld.proto  --js_out=import_style=commonjs:.  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:.

protoc-go:
	protoc -I=. --go_out=plugins=grpc:. helloworld.proto
