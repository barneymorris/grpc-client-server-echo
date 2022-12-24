# grpc-client-server-echo
Simple Client-Server gRPC pet-project app for studing

Just trying out gRPC with golang.

For use: 
1) install all required libs (protoc, etc)
2) generate all protobufs files by running the following commands from root of this project: 

chmod +x ./proto.sh
./proto.sh

3) run server: go run server.go
4) run client: go run client.go

and see at you terminal for some messages
