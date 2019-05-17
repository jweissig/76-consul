# Consul Demo

[#76 - Consul (screencast)](https://sysadmincasts.com/episodes/76-consul)

##### Build go app for linux (from mac)

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o web ./main.go
```

##### Build/push container

```sh
docker build -t jweissig/consul-demo .
docker push jweissig/consul-demo
```

##### Run container

https://hub.docker.com/r/jweissig/consul-demo

```sh
docker run -p 5005:5005 --rm jweissig/consul-demo
```

##### Connect to localhost

http://localhost:5005/
