package translator

import "fmt"

func (gt *GptTranslator) Translate(text string, args ...string) (string, error) {
	return "", nil
}

func (gt *GptTranslator) WrapText(text string) string {
	return fmt.Sprintf("```\nTranslation (GPT):\n------------------------\n%s\n```", text)
}
