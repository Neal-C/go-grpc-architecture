proto-codegen:
	protoc --proto_path=proto-stuff proto-stuff/*.proto --go_out=. --go-grpc_out=.