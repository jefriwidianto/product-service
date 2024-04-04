### Terminal

How to run with Docker after clone
```bash
# Script for running build docker image
docker build --no-cache --build-arg APP_ENV=Docker -t diksha_product_service -f Infrastructures/Docker/Dockerfile .

# Running script Docker-Compose
docker compose up
```