package lang_handler

type LanguageConfig struct {
	Extension string `json:"extension"`
	Execute   string `json:"execute"`
	Compile   string `json:"compile"`
}
