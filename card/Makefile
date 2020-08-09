protoc:
	protoc --proto_path=proto card.proto \
  		--go_out=plugins=grpc:gen \
		--grpc-gateway_out=:gen \
		--swagger_out=:swagger