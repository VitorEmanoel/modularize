package plugins

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type PluginInfo struct {
	UUID		string		`json:"uuid"`
	Name		string		`json:"name"`
	Version		string		`json:"version"`
	Url			string		`json:"url"`
}

func (p *PluginInfo) Download(folder string) error {
	response, err := http.Get(p.Url)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	file, err := os.Create(fmt.Sprintf("%s/%s@%s.so", folder, p.Name, p.Version))
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}
	return nil
}
