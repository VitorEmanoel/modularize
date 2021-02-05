package models

type PluginPack	struct {
	ID		int64		`json:"id" gorm:"primary"`
	Name	string		`json:"name"`
	Plugins	[]*Plugin	`json:"plugins,omitempty"`
}

