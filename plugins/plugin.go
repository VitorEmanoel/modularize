package plugins

import "plugin"

type Plugin interface {
	FindSymbol(symbol string) (interface{}, error)
}

type PluginContext struct {
	Path 	string
	Plugin	*plugin.Plugin
}

func (p *PluginContext) FindSymbol(symbol string) (interface{}, error) {
	if p.Plugin != nil {
		symbol, err := p.Plugin.Lookup(symbol)
		if err != nil {
			return nil, err
		}
		return symbol, nil
	}
	pl, err := plugin.Open(p.Path)
	if err != nil {
		return nil, err
	}
	p.Plugin = pl
	return p.FindSymbol(symbol)
}

func NewPlugin(path string) Plugin {
	return &PluginContext{Path: path}
}


