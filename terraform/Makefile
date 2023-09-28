PROJECT_NAME=$(shell basename $(shell pwd))
TARGET=devbox
IMAGE_TAG=${PROJECT_NAME}:dev

RUN=@docker run -it --rm --volume $(shell pwd):/src --volume ${HOME}/.aws:/root/.aws --workdir /src --env AWS_PROFILE ${IMAGE_TAG}

init:
	${RUN} init

plan:
	${RUN} plan -out plan.tfplan

apply:
	${RUN} apply plan.tfplan

destroy:
	${RUN} destroy
