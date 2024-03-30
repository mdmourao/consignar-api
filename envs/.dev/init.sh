docker build --target builder -f ../../Dockerfile ../../ -t api
docker compose -p api up