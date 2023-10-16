# li17server

## make image

```
make image
```

## init db

```
psql   -U postgres -h localhost -p 5432 -c "create database mpc"
```

```
psql   -U postgres -h localhost -p 5432 -d mpc -f ./manifest/migration/mpc_dump.sql
```

## run

```
docker run -it -v /path/config.yaml:/server/config.yaml -p 8000:8000 apiserver
```
