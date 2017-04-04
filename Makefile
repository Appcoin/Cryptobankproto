
PROTO_PATH?=${GOPATH}/src/github.com/crypto-bank/proto

deps:
	go get github.com/gogo/protobuf/gogoproto
	go get -u github.com/gogo/protobuf/protoc-gen-gogofaster

protoc:
	protoc --gogofaster_out=${GOPATH}/src  -I. -I${GOPATH}/src  $(PROTO_PATH)/currency/currency.proto
	protoc --gogofaster_out=${GOPATH}/src  -I. -I${GOPATH}/src  $(PROTO_PATH)/exchange/exchange.proto
	protoc --gogofaster_out=${GOPATH}/src  -I. -I${GOPATH}/src  $(PROTO_PATH)/order/order.proto
	protoc --gogofaster_out=${GOPATH}/src  -I. -I${GOPATH}/src  $(PROTO_PATH)/orderbook/orderbook.proto
	protoc --gogofaster_out=Mgithub.com/gogo/protobuf/protobuf/google/protobuf/timestamp.proto=github.com/gogo/protobuf/types:${GOPATH}/src  \
		-I. -I${GOPATH}/src  $(PROTO_PATH)/order/order.proto
