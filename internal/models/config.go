package models

type BotConfig struct {
	Token           string   `json:"token"`
	Channels        []string `json:"channels"`
	Trigger         string   `json:"trigger"`
	Interval        string   `json:"interval"`
	Mentions        []string `json:"mentions"`
	PrivilegedUsers []string `json:"privileged_users"`
}

type ScraperConfig struct {
	Instances []string `json:"instances"`
	Interval  string   `json:"interval"`
	Usernames []string `json:"usernames"`
	Redirect  string   `json:"redirect"`
}

type TranslatorConfig struct {
	Engine string `json:"engine"`
	DeepL  struct {
		Url      string `json:"url"`
		Key      string `json:"key"`
		Glossary string `json:"glossary"`
	} `json:"deepl"`
	Gpt struct {
		Url   string `json:"url"`
		Key   string `json:"key"`
		Model string `json:"model"`
	} `json:"gpt"`
}

type ChatConfig struct {
	Public bool `json:"public"`
	OpenAI struct {
		Model  string `json:"model"`
		Url    string `json:"url"`
		Key    string `json:"key"`
		System string `json:"system"`
	} `json:"openai"`
}

type Config struct {
	Bot       BotConfig
	Scrape    ScraperConfig
	Translate TranslatorConfig
	Chat      ChatConfig
}
