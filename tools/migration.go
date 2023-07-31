package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"github.com/bamboo-firewall/watcher/config"
	"github.com/bamboo-firewall/watcher/repository"
	"github.com/bamboo-firewall/watcher/utils"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	log.Println("start etcd watcher")
	cfg := config.New()
	options := options.Client().ApplyURI(cfg.MongoURI)
	options.SetMaxPoolSize(128)
	options.SetMinPoolSize(5)
	options.SetMaxConnIdleTime(time.Second * 30)
	options.SetHeartbeatInterval(time.Second * 10)

	db, err := mongo.Connect(context.TODO(), options)
	if err != nil {
		log.Fatalln(err)
	}

	keyChain, err := utils.KeyDecode(cfg.EtcdCACert, cfg.EtcdServerCert, cfg.EtcdServerKey)
	if err != nil {
		log.Fatalln(err)
	}
	c, _ := tls.X509KeyPair(keyChain.Cert, keyChain.Key)

	rootCA := x509.NewCertPool()
	rootCA.AppendCertsFromPEM(keyChain.CA)

	tlsConf := &tls.Config{Certificates: []tls.Certificate{c}, RootCAs: rootCA}
	etcd, err := clientv3.New(clientv3.Config{
		Endpoints:   cfg.EtcdPeerURLS,
		DialTimeout: 5 * time.Second,
		Username:    cfg.EtcdUsername,
		Password:    cfg.EtcdPassword,
		TLS:         tlsConf,
	})
	if err != nil {
		log.Fatalln(err)
	}

	ds := repository.New(&repository.WatcherRepository{
		MongoConnect: db,
		EtcdConnect:  etcd,
	})

	ds.Migration(context.Background(), cfg.EtcdWatcherPath, cfg.MongoDBName)
}
