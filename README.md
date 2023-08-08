# watcher
Calico watcher will monitor changing on etcd cluster and update to mongodb database. Then calicobe (calico backend) project will provide API for frontend of bamboo firewall can read them.

## Howto build
- build daemon
```shell
 docker run --rm \
            -v `pwd`:/go/src/github.com/bamboo-firewall/watcher \
            -w /go/src/github.com/bamboo-firewall/watcher \
            golang:1.18 sh -c \
            'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -a -installsuffix cgo -o watcher cmd/main.go'
```
- build migration tool
```shell
 docker run --rm \
            -v `pwd`:/go/src/github.com/bamboo-firewall/watcher \
            -w /go/src/github.com/bamboo-firewall/watcher \
            golang:1.18 sh -c \
            'GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=mod -a -installsuffix cgo -o watcher_migration tools/migration.go'
```

## Howto run
- Create an environment like above:
```shell
  "BAMBOOFW_APP_ENV": "production",
  "BAMBOOFW_ENABLE_DEBUG_MODE": "false",
  "BAMBOOFW_ETCD_CA_CERT": "base64_encode_content_ca_crt",
  "BAMBOOFW_ETCD_PASSWORD": "change_me",
  "BAMBOOFW_ETCD_PEER_URLS": "https://etcd01:2379 https://etcd02:2379 https://etcd03:2379",
  "BAMBOOFW_ETCD_SERVER_CERT": "base64_encode_content_server_crt",
  "BAMBOOFW_ETCD_SERVER_KEY": "base64_encode_content_server_key",
  "BAMBOOFW_ETCD_USERNAME": "watcher",
  "BAMBOOFW_ETCD_WATCHER_PATH": "/calico/resources/v3/projectcalico.org",
  "BAMBOOFW_MONGODB_NAME": "bamboofw",
  "BAMBOOFW_MONGODB_URI": "mongodb://bamboofw:changeme@mongodb01:27017,mongodb02:27017,mongodb03:27017/?authSource=bamboofw"
```
- Run
```shell
./watcher --> to run daemon
./watcher_migration --> to migrate exits data from etcd to mongodb
```
