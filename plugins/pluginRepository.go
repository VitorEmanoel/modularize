package plugins

type PluginSearch struct {
	Pack	string		`json:"pack"`
}

type PluginRepository interface {
	List(search PluginSearch) ([]PluginInfo, error)
}
