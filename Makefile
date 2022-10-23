release:
	bash ./scripts/make.sh release

build:
	bash -x ./.github/scripts/build.sh

build-dev:
	go build -ldflags "-X 'github.com/guionardo/go-dev/pkg.consts.CompileTimeString=2022-10-23T13:35:39' -X 'pkg.consts.Version=1.2.0' -X 'pkg.consts.BuildRunner=guionardo@guio-notehp'" .


build-bin:
	export DOCKER_BUILDKIT=1
	@docker build . --target bin --output bin/ --platform local
	bin/go-dev install

