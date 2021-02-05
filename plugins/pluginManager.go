package plugins

import (
	"os"
	"path/filepath"
)

type PluginManager interface {
	LoadFile(file string) Plugin
	LoadFolder(folder string) []Plugin
}

type PluginManagerContext struct {

}

func (p *PluginManagerContext) LoadFile(file string) Plugin {
	return NewPlugin(file)
}

func (p *PluginManagerContext) LoadFolder(folder string) []Plugin {
	var plugins []Plugin
	err := filepath.Walk(folder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
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

func NewPluginManager() PluginManager {
	return &PluginManagerContext{}
}
