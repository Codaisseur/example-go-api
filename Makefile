build:
	docker build -t quay.io/codaisseur/example-go-api .

run:
	docker run --rm -P quay.io/codaisseur/example-go-api

push:
	docker push quay.io/codaisseur/example-go-api
