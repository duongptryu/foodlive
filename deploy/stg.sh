APP_NAME=foodlive

docker load -i ${APP_NAME}.tar
docker rm -f ${APP_NAME}

docker run -d --name foodlive \
  -e VIRTUAL_HOST="foodlive.tech" \
  -e LETSENCRYPT_HOST="foodlive.tech" \
  -e LETSENCRYPT_EMAIL="duongptryu@gmail.com" \
  --net my-net \
  -p 8080:8080 \
  foodlive