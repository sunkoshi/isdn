package lang_handler

import (
	"encoding/json"

	"github.com/orted-org/isdn/internal/file_manager"
)

func getLanguageConfig() ([]LanguageConfig, error) {

	var configs []LanguageConfig
	fm := file_manager.New("./config")
	configJson, err := fm.Get("language_config.json")

	if err != nil {
		return configs, err
	}

	err = json.Unmarshal(configJson, &configs)
	if err != nil {
		return configs, err
	}

	return configs, nil
}
