package models


type Plugin struct {
	ID		int64		`json:"id" gorm:"primary"`
	Name	string		`json:"name"`
}

type PluginVersion struct {
	ID			int64		`json:"id" gorm:"id"`
	Version		string		`json:"version"`
	Plugin		*Plugin		`json:"plugin,omitempty"`
	PluginID	int64		`json:"pluginId"`
}