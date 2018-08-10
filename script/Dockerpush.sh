echo "$DOCKER_PASSWORD" | docker login -u "$DOCKER_USERNAME" --password-stdin
docker push statnlp/service-ladrawex
docker push statnlp/service-ladrawex:${APP_VERSION}