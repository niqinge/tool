package apollocenter

import (
	"github.com/philchia/agollo/v4"
)

type ApolloConfig struct {
	AppID          string   `json:"app_id"`
	Cluster        string   `json:"cluster,omitempty"`
	NameSpaceNames []string `json:"namespace_names,omitempty"`
	MetaAddr       string   `json:"meta_addr,omitempty"`
}

type ApolloCenter struct {
	conf *ApolloConfig
}

func NewApolloCenter(conf *ApolloConfig) *ApolloCenter {
	return &ApolloCenter{conf: conf}
}

func (ac *ApolloCenter) Start() error {
	return agollo.Start(&agollo.Conf{
		AppID:          ac.conf.AppID,
		Cluster:        ac.conf.Cluster,
		NameSpaceNames: ac.conf.NameSpaceNames,
		MetaAddr:       ac.conf.MetaAddr,
	})
}

func (ac *ApolloCenter) GetValue(key string, nameSpace string) string {
	return agollo.GetString(key, agollo.WithNamespace(nameSpace))
}

func (ac *ApolloCenter) Close() error {
	return agollo.Stop()
}
