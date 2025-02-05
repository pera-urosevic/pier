package types

type DatabaseIndex struct {
	Database   string `json:"database"`
	Collection string `json:"collection"`
	Facets     string `json:"facets"`
}
