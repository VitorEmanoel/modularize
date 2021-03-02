package plugins

import (
	"os"
	"path/filepath"
)

type PluginLoader interface {
	LoadFile(file string) Plugin
	LoadFolder(folder string) []Plugin
}

type PluginLoaderContext struct {

}

func (p *PluginLoaderContext) LoadFile(file string) Plugin {
	return NewPlugin(file)
}

func (p *PluginLoaderContext) LoadFolder(folder string) []Plugin {
	var plugins []Plugin
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(info.Name()) != ".so" {
			return nil
		}
		plugins = append(plugins, NewPlugin(path))
		return nil
	})
	if err != nil {
		return nil
	}
	return plugins
}

func NewPluginLoader() PluginLoader {
	return &PluginLoaderContext{}
}
