proto:
#protoc pkg/pb/*.proto --go_out=plugins=grpc:.
	protoc --go_out=. --go-grpc_out=. pkg/**/pb/*.proto 
server:
	go run cmd/main.go