release:
	bash ./scripts/make.sh release

build:
	go build .

build-bin:
	export DOCKER_BUILDKIT=1
	@docker build . --target bin --output bin/ --platform local
	bin/go-dev install

