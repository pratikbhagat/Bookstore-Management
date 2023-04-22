all:	clean build system-service

build:
	go mod download
	GOOS=linux go build cmd/main/main.go

clean:
	rm -vf main

start:
	./main

srcstart:
	go run cmd/main/main.go

system-service:
	cp bookstore-management.service /etc/systemd/system/
	mkdir -p /var/log/apps/bookstore-management
	systemctl daemon-reload
	systemctl start bookstore-management.service