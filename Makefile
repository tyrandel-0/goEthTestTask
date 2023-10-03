all: install run
run:
	go run cmd/main.go

install:
	go mod tidy

deleteWallets:
	rm storage/keystore/*