.PHONY: update
update:
	sh ./scripts/update_modules.sh
	make build

.PHONY: build
build:
	go mod vendor
	go build -mod=vendor ./pkg