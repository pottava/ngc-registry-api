# REST APIs for NVIDIA GPU Cloud (NGC)

You can retrieve repository / image information from NGC via simple REST API.

## Usage

### Run API server

```console
pushd app
go mod vendor
popd
docker-compose -f app/runtime.yml up
```

### Consume APIs

signin

```console
curl -s -X POST -H 'Content-Type:application/json' -d '{"email":"ngc-user@example.com","password":"Passw0rd"}' http://localhost:9000/api/v1/signin
```

get repositories

```console
curl -s -X GET -H 'Content-Type:application/json' -H 'Authorization:base64encodedsession' http://localhost:9000/api/v1/repositories
```

get images

```console
curl -s -X GET -H 'Content-Type:application/json' -H 'Authorization:base64encodedsession' http://localhost:9000/api/v1/repositories/nvidia/tensorflow/images
```
