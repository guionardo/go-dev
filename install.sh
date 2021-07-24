#!/bin/bash
export DOCKER_BUILDKIT=1
@docker build . --target bin --output bin/ --platform local
bin/go-dev install