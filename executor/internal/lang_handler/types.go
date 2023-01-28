package lang_handler

type LanguageConfig struct {
	Name      string `json:"name"`
	Extension string `json:"extension"`
	Execute   string `json:"execute"`
	Compile   string `json:"compile"`
}
