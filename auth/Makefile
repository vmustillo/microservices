protoc:
	protoc --proto_path=proto auth.proto \
  		--go_out=plugins=grpc:gen \
		--grpc-gateway_out=:gen \
		--swagger_out=:swagger