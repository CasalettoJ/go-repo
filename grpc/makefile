BUILD_OUT := bin/

SERVER_PKG := server/main.go
CLIENT_PKG := client/main.go

SERVER_OUT := $(BUILD_OUT)server.bin
CLIENT_OUT := $(BUILD_OUT)client.bin
NETWORKING_OUT := networking/chat.pb.go

DEP_FILE := Gopkg.toml
SSL_CRT := consensusprotocol.crt
SSL_KEY := consensusprotocol.key

all: networking dep build_server build_client

networking/chat.pb.go: networking/chat.proto
	protoc -I=networking/ --go_out=plugins=grpc:networking/ networking/chat.proto

networking: networking/chat.pb.go

build_server: networking dep
	go build -o $(SERVER_OUT) $(SERVER_PKG)

build_client: networking dep
	go build -o $(CLIENT_OUT) $(CLIENT_PKG)

clean:
	rm $(SERVER_OUT) $(CLIENT_OUT) $(NETWORKING_OUT)

$(DEP_FILE):
	dep init

dep: $(DEP_FILE)
	dep ensure
	cp $(SSL_CRT) $(BUILD_OUT)$(SSL_CRT)
	cp $(SSL_KEY) $(BUILD_OUT)$(SSL_KEY)