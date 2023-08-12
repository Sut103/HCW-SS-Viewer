package aws

import "time"

type Screenshot struct {
	URL            string `json:"URL"`
	ChannelMessage struct {
		Attachments []struct {
			URL         string `json:"URL"`
			ContentType string `json:"ContentType"`
			Filename    string `json:"Filename"`
			Height      int    `json:"Height"`
			ID          string `json:"ID"`
			ProxyURL    string `json:"ProxyURL"`
			Size        int    `json:"Size"`
			Width       int    `json:"Width"`
		} `json:"Attachments"`
		Author struct {
			Avatar           string `json:"Avatar"`
			AvatarDecoration any    `json:"AvatarDecoration"`
			Discriminator    string `json:"Discriminator"`
			ID               string `json:"ID"`
			PublicFlags      int    `json:"PublicFlags"`
			Username         string `json:"Username"`
		} `json:"Author"`
		ChannelID       string    `json:"ChannelID"`
		Components      []any     `json:"Components"`
		Content         string    `json:"Content"`
		EditedTimestamp any       `json:"EditedTimestamp"`
		Embeds          []any     `json:"Embeds"`
		Flags           int       `json:"Flags"`
		ID              string    `json:"ID"`
		MentionEveryone bool      `json:"MentionEveryone"`
		MentionRoles    []any     `json:"MentionRoles"`
		Mentions        []any     `json:"Mentions"`
		Pinned          bool      `json:"Pinned"`
		Timestamp       time.Time `json:"Timestamp"`
		Tts             bool      `json:"Tts"`
		Type            int       `json:"Type"`
	} `json:"ChannelMessage"`
}
