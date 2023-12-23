package translator

import (
	"errors"
	"mamelon/internal/models"
)

type Translator interface {
	Translate(string) (string, error)
	WrapText(string) string
}

type DeepLTranslator struct {
	apiUrl   string
	apiKey   string
	glossary string
}

type GptTranslator struct {
	apiUrl string
	apiKey string
	model  string
}

func New(tcfg models.TranslatorConfig) (Translator, error) {

	switch tcfg.Engine {

	case "deepl":
		return &DeepLTranslator{
			apiUrl:   tcfg.DeepL.Url,
			apiKey:   tcfg.DeepL.Key,
			glossary: tcfg.DeepL.Glossary,
		}, nil

	case "gpt":
		return &GptTranslator{
			apiUrl: tcfg.Gpt.Url,
			apiKey: tcfg.Gpt.Key,
			model:  tcfg.Gpt.Model,
		}, nil

	default:
		return nil, errors.New("unsupported translator engine")

	}

}
