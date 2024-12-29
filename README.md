```
make start

make stop

docker run --name pgadmin \
  -e PGADMIN_DEFAULT_EMAIL=chanyeintun@gmail.com \
  -e PGADMIN_DEFAULT_PASSWORD=admin \
  -p 5050:80 \
  dpage/pgadmin4

docker rm -f $(docker ps -aq)

docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' teacher_booking_db

air -c .air.toml
```