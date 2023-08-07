package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/bamboo-firewall/watcher/handler"
	models "github.com/bamboo-firewall/watcher/model"
	"go.etcd.io/etcd/client/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

var (
	hostEndpoint          = "hostendpoints"
	globalNetworkSets     = "globalnetworksets"
	globalNetworkPolicies = "globalnetworkpolicies"
)

type WatcherRepository struct {
	MongoConnect *mongo.Client
	EtcdConnect  *clientv3.Client
}

func New(ds *WatcherRepository) *WatcherRepository {
	return &WatcherRepository{
		MongoConnect: ds.MongoConnect,
		EtcdConnect:  ds.EtcdConnect,
	}
}

func (w *WatcherRepository) Watch(c context.Context, key, dbname string) {
	watchChan := w.EtcdConnect.Watch(c, key, clientv3.WithPrefix())
	log.Printf("watch on %s path\n", key)
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			coll, eventKey, err := handler.Event(string(event.Kv.Key))
			if err != nil {
				log.Printf("Got an error: %v", err)
			}
			if string(event.Type.String()) == "DELETE" {
				err := w.DeleteDoc(dbname, coll, eventKey)
				if err != nil {
					log.Printf("Got an error: %v", err)
				}
			} else {
				var hep models.HostEndPoint
				var gns models.GlobalNetworkSet
				var gnp models.GlobalNetworkPolicies
				if coll == hostEndpoint {
					json.Unmarshal(event.Kv.Value, &hep)
					hep.ID = hep.Metadata.UID
					hep.Metadata.Name = eventKey
					count, err := w.FindDoc(dbname, coll, hep.ID)
					if err != nil {
						fmt.Println("Error find doc", err)
					}
					if count >= 1 {
						err := w.DeleteByID(dbname, coll, hep.ID)
						if err != nil {
							fmt.Printf("Error to delete id: %s\n", hep.ID)
						}
						fmt.Printf("Deleted id: %s \n", hep.ID)
					}
					inErr := w.InsertDoc(dbname, coll, hep)
					if inErr != nil {
						fmt.Println("Error write doc", inErr)
					}
				}
				if coll == globalNetworkSets {
					json.Unmarshal(event.Kv.Value, &gns)
					gns.ID = gns.Metadata.UID
					gns.Metadata.Name = eventKey
					count, err := w.FindDoc(dbname, coll, gns.ID)
					if err != nil {
						fmt.Println("Error find doc", err)
					}
					if count >= 1 {
						err := w.DeleteByID(dbname, coll, gns.ID)
						if err != nil {
							fmt.Printf("Error to delete id: %s\n", gns.ID)
						}
						fmt.Printf("Deleted id: %s \n", gns.ID)
					}
					inErr := w.InsertDoc(dbname, coll, gns)
					if inErr != nil {
						fmt.Println("Error write doc", inErr)
					}
				}
				if coll == globalNetworkPolicies {
					json.Unmarshal(event.Kv.Value, &gnp)
					gnp.ID = gnp.Metadata.UID
					gnp.Metadata.Name = eventKey
					count, err := w.FindDoc(dbname, coll, gnp.ID)
					if err != nil {
						fmt.Println("Error find doc", err)
					}
					if count >= 1 {
						err := w.DeleteByID(dbname, coll, gnp.ID)
						if err != nil {
							fmt.Printf("Error to delete id: %s\n", gnp.ID)
						}
						fmt.Printf("Deleted id: %s \n", gnp.ID)
					}
					inErr := w.InsertDoc(dbname, coll, gnp)
					if inErr != nil {
						fmt.Println("Error write doc", inErr)
					}
				}
			}
		}
	}
}

func (w *WatcherRepository) Migration(c context.Context, key, dbname string) {
	r, _ := w.EtcdConnect.Get(c, key, clientv3.WithPrefix())
	for i := 0; i < len(r.Kvs); i++ {
		coll, eventKey, err := handler.Event(string(r.Kvs[i].Key))
		if err != nil {
			log.Printf("Got an error: %v", err)
		}
		var hep models.HostEndPoint
		var gns models.GlobalNetworkSet
		var gnp models.GlobalNetworkPolicies
		if coll == hostEndpoint {
			json.Unmarshal(r.Kvs[i].Value, &hep)
			hep.ID = hep.Metadata.UID
			hep.Metadata.Name = eventKey
			err := w.InsertDoc(dbname, coll, hep)
			if err != nil {
				fmt.Println("Error write doc", err)
			}
		}
		if coll == globalNetworkSets {
			json.Unmarshal(r.Kvs[i].Value, &gns)
			gns.ID = gns.Metadata.UID
			gns.Metadata.Name = eventKey
			err := w.InsertDoc(dbname, coll, gns)
			if err != nil {
				fmt.Println("Error write doc", err)
			}
		}
		if coll == globalNetworkPolicies {
			json.Unmarshal(r.Kvs[i].Value, &gnp)
			gnp.ID = gnp.Metadata.UID
			gnp.Metadata.Name = eventKey
			err := w.InsertDoc(dbname, coll, &gnp)
			if err != nil {
				fmt.Println("Error write doc", err)
			}
		}
	}

}

func (w *WatcherRepository) InsertDoc(dbname, coll string, data interface{}) error {
	collCon := w.MongoConnect.Database(dbname).Collection(coll)
	res, err := collCon.InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	log.Printf("document id: %s inserted\n", res.InsertedID)
	return nil
}

func (w *WatcherRepository) DeleteByID(dbname, coll string, id string) error {
	collCon := w.MongoConnect.Database(dbname).Collection(coll)
	filter := bson.D{{"_id", id}}

	_, err := collCon.DeleteOne(context.TODO(), filter)
	return err
}

func (w *WatcherRepository) FindDoc(dbname, coll string, id string) (int64, error) {
	collCon := w.MongoConnect.Database(dbname).Collection(coll)
	filter := bson.D{{"_id", id}}
	return collCon.CountDocuments(context.TODO(), filter)
}

func (w *WatcherRepository) DeleteDoc(dbname, coll string, name string) error {
	collCon := w.MongoConnect.Database(dbname).Collection(coll)
	filter := bson.M{
		"metadata.name": name,
	}
	res, err := collCon.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	log.Printf("Number of documents deleted: %d with filter is {\"metadata.name\": \"%s\"} \n", res.DeletedCount, name)
	return nil
}

func (w *WatcherRepository) FetchAll(c context.Context) {
	r, _ := w.EtcdConnect.Get(c, "/", clientv3.WithPrefix())
	fmt.Println(r.Kvs)
}
