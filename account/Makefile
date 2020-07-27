protoc:
	protoc -I/usr/local/include -I. \
  		-I$$GOPATH/src \
  		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		--proto_path=proto account.proto \
  		--go_out=plugins=grpc:gen \
		--grpc-gateway_out=:gen \
		--swagger_out=:swagger