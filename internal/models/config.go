package models

type Config struct {
	Bot struct {
		Token           string   `json:"token"`
		Channels        []string `json:"channels"`
		Trigger         string   `json:"trigger"`
		Interval        string   `json:"interval"`
		Mentions        []string `json:"mentions"`
		PrivilegedUsers []string `json:"privileged_users"`
	}
	Scrape struct {
		Instances []string `json:"instances"`
		Interval  string   `json:"interval"`
		Usernames []string `json:"usernames"`
		Redirect  string   `json:"redirect"`
	}
	Translate struct {
		DeepL struct {
			Url string `json:"url"`
			Key string `json:"key"`
		} `json:"deepl"`
	}
	Chat struct {
		Public bool `json:"public"`
		OpenAI struct {
			Model  string `json:"model"`
			Url    string `json:"url"`
			Key    string `json:"key"`
			System string `json:"system"`
		} `json:"openai"`
	}
}
