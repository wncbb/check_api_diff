package conf

type EnvValue struct {
	Key     string `json:"key"`
	Value   string `json:"value"`
	Typ     string `json:"type"`
	Enabled bool   `json:"enabled"`
}

type Env struct {
	ID             string      `json:"id"`
	Name           string      `json:"name"`
	Values         []*EnvValue `json:"values"`
	Timestamp      int64       `json:"timestamp"`
	Synced         bool        `json:"synced"`
	SyncedFilename string      `json:"syncedFilename"`
	Team           string      `json:"team"`
	IsDeleted      bool        `json:"isDeleted"`
}
