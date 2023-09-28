PROJECT_NAME=$(shell basename $(shell pwd))
TARGET=devbox
IMAGE_TAG=${PROJECT_NAME}:dev
RUN=@docker run -it --rm --volume $(shell pwd):/src --workdir /src --publish 3000:3000 ${IMAGE_TAG}

build:
	@docker build --target ${TARGET} . --tag ${IMAGE_TAG}

shell: build
	${RUN} sh

dev: build
	${RUN} npm start
