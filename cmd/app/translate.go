package main

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TranslationResponse struct {
	Translations []struct {
		DetectedSourceLanguage string `json:"detected_source_language"`
		Text                   string `json:"text"`
	} `json:"translations"`
}

func (app *application) translate(text string) (string, error) {

	url := app.config.Translate.DeepL.Url
	key := app.config.Translate.DeepL.Key

	requestData := map[string]interface{}{
		"text":        []string{text},
		"target_lang": "EN",
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

	return tr.Translations[0].Text, nil
}
