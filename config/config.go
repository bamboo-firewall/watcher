package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	AppEnv          string   `mapstructure:"BAMBOOFW_APP_ENV"`
	EnableDebugMode bool     `mapstructure:"BAMBOOFW_ENABLE_DEBUG_MODE"`
	MongoURI        string   `mapstructure:"BAMBOOFW_MONGODB_URI"`
	MongoDBName     string   `mapstructure:"BAMBOOFW_MONGODB_NAME"`
	EtcdPeerURLS    []string `mapstructure:"BAMBOOFW_ETCD_PEER_URLS"`
	EtcdUsername    string   `mapstructure:"BAMBOOFW_ETCD_USERNAME"`
	EtcdPassword    string   `mapstructure:"BAMBOOFW_ETCD_PASSWORD"`
	EtcdCACert      string   `mapstructure:"BAMBOOFW_ETCD_CA_CERT"`
	EtcdServerCert  string   `mapstructure:"BAMBOOFW_ETCD_SERVER_CERT"`
	EtcdServerKey   string   `mapstructure:"BAMBOOFW_ETCD_SERVER_KEY"`
	EtcdWatcherPath string   `mapstructure:"BAMBOOFW_ETCD_WATCHER_PATH"`
}

func New() *Configuration {
	vip := viper.New()
	vip.AutomaticEnv()
	vip.SetConfigFile(".env")
	err := vip.ReadInConfig()
	if err != nil {
		log.Println(err)
		log.Println("file .env not exits will be read from environment")
		return &Configuration{
			AppEnv:          vip.GetString("BAMBOOFW_APP_ENV"),
			EnableDebugMode: vip.GetBool("BAMBOOFW_ENABLE_DEBUG_MODE"),
			MongoURI:        vip.GetString("BAMBOOFW_MONGODB_URI"),
			MongoDBName:     vip.GetString("BAMBOOFW_MONGODB_NAME"),
			EtcdPeerURLS:    vip.GetStringSlice("BAMBOOFW_ETCD_PEER_URLS"),
			EtcdUsername:    vip.GetString("BAMBOOFW_ETCD_USERNAME"),
			EtcdPassword:    vip.GetString("BAMBOOFW_ETCD_PASSWORD"),
			EtcdCACert:      vip.GetString("BAMBOOFW_ETCD_CA_CERT"),
			EtcdServerCert:  vip.GetString("BAMBOOFW_ETCD_SERVER_CERT"),
			EtcdServerKey:   vip.GetString("BAMBOOFW_ETCD_SERVER_KEY"),
			EtcdWatcherPath: vip.GetString("BAMBOOFW_ETCD_WATCHER_PATH"),
		}
	}
	config := Configuration{}
	err = vip.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}
	return &config

}
