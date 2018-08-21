CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main.out ./cmd/ladrawex/
docker build -t statnlp/service-ladrawex .
docker tag  statnlp/service-ladrawex statnlp/service-ladrawex:${APP_VERSION}