.PHONY: app-up
app-up:
	docker compose up -d

.PHONY: app-down
app-down:
	docker compose down

IMAGES := gateway task user reminder course frontend

.PHONY: build-images $(IMAGES)
build-images: $(IMAGES)

$(IMAGES):
	docker build -t image-$@ -f Dockerfile.$@ .