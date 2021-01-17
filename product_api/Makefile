GOOS	 ?= linux
GOARCH  = amd64
ODIR    = _output

IMAGE_NAME = rufuskopral/efishery_product_api

docker-build:
	docker build -t $(IMAGE_NAME) -f ./deploy/Dockerfile .

docker-push:
		docker push $(IMAGE_NAME);