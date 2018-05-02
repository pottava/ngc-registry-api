# REST APIs for NVIDIA GPU Cloud (NGC)

You can retrieve repository / image information from NGC via simple REST API.

## Usage

### Run API server

```
$ docker-compose -f app/tools.yml run --rm deps ensure
$ docker-compose -f app/runtime.yml up
```

### Consume APIs

signin

```
$ curl -s -X POST -H 'Content-Type:application/json' -d '{"email":"ngc-user@example.com","password":"Passw0rd"}' http://localhost:9000/api/v1/signin
```

get repositories

```
$ curl -s -X GET -H 'Content-Type:application/json' -H 'Authorization:base64encodedsession' http://localhost:9000/api/v1/repositries
```

get images

```
$ curl -s -X GET -H 'Content-Type:application/json' -H 'Authorization:base64encodedsession' http://localhost:9000/api/v1/repositories/nvidia/tensorflow/images
```
