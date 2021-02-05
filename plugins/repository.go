package plugins

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PluginPack struct {
	Name		string			`json:"name"`
	Plugins		[]PluginInfo	`json:"plugins,omitempty"`
}

func (p *PluginPack) Download(folder string) error {
	for _, plugin := range p.Plugins {
		err := plugin.Download(folder)
		if err != nil {
			return err
		}
	}
	return nil
}

type Repository interface {
	ListPacks() (*PluginPack, error)
}

type RepositoryContext struct {
	Url		string
}

func (r *RepositoryContext) ListPacks() (*PluginPack, error) {
	response, err :=  http.Get(fmt.Sprintf("%s/packs", r.Url))
	if err != nil {
		return nil, err
	}
	bodyContent, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var pack *PluginPack
	err = json.Unmarshal(bodyContent, pack)
	if err != nil {
		return nil, err
	}
	return pack, nil
}

func NewRepository(url string) Repository {
	return &RepositoryContext{Url: url}
}
