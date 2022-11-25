package lang_handler

import (
	"path"
	"strings"
)

type LanguageHandler struct {
	Configs map[string]LanguageConfig `json:"configs"`
}

func New() (*LanguageHandler, error) {

	var lh LanguageHandler = LanguageHandler{
		Configs: map[string]LanguageConfig{},
	}

	configArray, err := getLanguageConfig()
	if err != nil {
		return nil, err
	}

	for _, v := range configArray {
		lh.Configs[v.Name] = v
	}

	return &lh, nil
}

func (lh *LanguageHandler) IfConfigExists(language string) bool {
	_, ok := lh.Configs[language]
	return ok
}

func (lh *LanguageHandler) GetExtension(language string) string {
	return lh.Configs[language].Extension
}

func (lh *LanguageHandler) GetCompileCmd(wDir, filename, language string) string {
	filePath := path.Join(wDir, filename)
	config := lh.Configs[language]
	tmp := config.Compile
	if strings.Contains(tmp, `%%filename%%`) {
		tmp = strings.ReplaceAll(tmp, `%%filename%%`, filePath)
	}

	if strings.Contains(tmp, `%%extension%%`) {
		tmp = strings.ReplaceAll(tmp, `%%extension%%`, config.Extension)
	}

	if strings.Contains(tmp, `%%path%%`) {
		tmp = strings.ReplaceAll(tmp, `%%path%%`, wDir)
	}

	return tmp
}

func (lh *LanguageHandler) GetExecutionCmd(wDir, filename, language string) string {
	filePath := path.Join(wDir, filename)
	config := lh.Configs[language]
	tmp := config.Execute
	if strings.Contains(tmp, `%%filename%%`) {
		tmp = strings.ReplaceAll(tmp, `%%filename%%`, filePath)
	}

	if strings.Contains(tmp, `%%extension%%`) {
		tmp = strings.ReplaceAll(tmp, `%%extension%%`, config.Extension)
	}

	if strings.Contains(tmp, `%%path%%`) {
		tmp = strings.ReplaceAll(tmp, `%%path%%`, wDir)
	}

	return tmp
}
