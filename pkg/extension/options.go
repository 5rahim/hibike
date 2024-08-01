package extension

const (
	SettingInputTypeText   SettingInputType = "text"
	SettingInputTypeSwitch SettingInputType = "switch"
	SettingInputTypeSelect SettingInputType = "select"
	SettingInputTypeNumber SettingInputType = "number"
)

type (
	SettingInputType string

	SettingInputOption struct {
		Value string `json:"value"`
		Label string `json:"label"`
	}

	SettingInputValueValidator func(value string) error

	SettingInput struct {
		Type    SettingInputType     `json:"type"`
		Name    string               `json:"name"`
		Options []SettingInputOption `json:"options"`
		Default string               `json:"default"`
	}
)
