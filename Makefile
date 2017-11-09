install:
	go-bindata settings.local.toml
	go install
