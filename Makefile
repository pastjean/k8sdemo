server-linux: main.go
	GOOS=linux go build -o server-linux .

.PHONY: v1
v1: server-linux
	docker build -t pastjean/thequoter:v1 .
	docker push pastjean/thequoter:v1

v2: server-linux
	docker build -t pastjean/thequoter:v2 .
	docker push pastjean/thequoter:v2