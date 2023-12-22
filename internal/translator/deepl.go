package translator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (dt *DeepLTranslator) Translate(text string) (string, error) {

	url := dt.apiUrl
	key := dt.apiKey
	glossary := dt.glossary

	requestData := map[string]interface{}{
		"text":        []string{text},
		"source_lang": "JA",
		"target_lang": "EN",
		"glossary_id": glossary,
	}

	jsonData, err := json.Marshal(requestData)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "DeepL-Auth-Key "+key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var tr TranslationResponse

	err = json.NewDecoder(resp.Body).Decode(&tr)
	if err != nil {
		return "", err
	}

	return dt.WrapText(tr.Translations[0].Text), nil
}

func (dt *DeepLTranslator) WrapText(text string) string {
	return fmt.Sprintf("```\nTranslation (DeepL):\n------------------------\n%s\n```", text)
}
