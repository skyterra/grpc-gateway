.PHONY: proto gateway clean serv-1 serv-2

gateway:
	cd gateway && go build .

serv-1:
	cd serv-1 && go build .

serv-2:
	cd serv-2 && go build .

proto:
	#protoc -I=./protobuf --go_out=plugins=grpc:pb ./protobuf/*.proto
	# generate the messages
	protoc -I=./protobuf --go_out=pb ./protobuf/*.proto

	# generate the services
	protoc -I=./protobuf --go-grpc_out=pb ./protobuf/*.proto

clean:
	rm -f ./gateway/gateway
	rm -f ./serv-1/serv-1
	rm -f ./serv-2/serv-2



