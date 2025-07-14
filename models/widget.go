package models

type PollOption struct {
	Id       string   `json:"id"`
	Text     string   `json:"text"`
	VoterIds []string `json:"voterIds"`
}
