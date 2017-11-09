run:
	go run main.go --settings=settings.local.toml hoge
install:
	go-bindata settings.local.toml
	go install
