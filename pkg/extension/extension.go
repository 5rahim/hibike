package extension

type Extension interface {
	GetInfo() ExtensionInfo
}

type ExtensionInfo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Version     string `json:"version"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Language    string `json:"language"`
	Type        string `json:"type"`
}
