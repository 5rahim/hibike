package extension

const (
	ConfigFieldTypeText   ConfigFieldType = "text"
	ConfigFieldTypeSwitch ConfigFieldType = "switch"
	ConfigFieldTypeSelect ConfigFieldType = "select"
	ConfigFieldTypeNumber ConfigFieldType = "number"
)

type (
	// ConfigField represents a field in an extension's configuration.
	// The fields are defined in the manifest file.
	ConfigField struct {
		Type    ConfigFieldType           `json:"type"`
		Name    string                    `json:"name"`
		Options []ConfigFieldSelectOption `json:"options"`
		Default string                    `json:"default"`
	}

	ConfigFieldType string

	ConfigFieldSelectOption struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	ConfigFieldValueValidator func(value string) error
)
