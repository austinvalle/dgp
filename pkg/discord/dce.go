package discord

import "time"

// https://github.com/Tyrrrz/DiscordChatExporter
type DceExport struct {
	Guild        DceGuild     `json:"guild"`
	Channel      DceChannel   `json:"channel"`
	Messages     []DceMessage `json:"messages"`
	MessageCount int          `json:"messageCount"`
}
type DceGuild struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}
type DceChannel struct {
	ID         string `json:"id"`
	Type       string `json:"type"`
	CategoryID string `json:"categoryId"`
	Category   string `json:"category"`
	Name       string `json:"name"`
}
type DceAuthor struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Discriminator string `json:"discriminator"`
	Nickname      string `json:"nickname"`
	Color         string `json:"color"`
	IsBot         bool   `json:"isBot"`
	AvatarURL     string `json:"avatarUrl"`
}
type DceMention struct {
	ID            string `json:"id"`
	Name          string `json:"name"`
	Discriminator string `json:"discriminator"`
	Nickname      string `json:"nickname"`
	IsBot         bool   `json:"isBot"`
}
type DceReference struct {
	MessageID string `json:"messageId"`
	ChannelID string `json:"channelId"`
	GuildID   string `json:"guildId"`
}
type DceMessage struct {
	ID                 string          `json:"id"`
	Type               string          `json:"type"`
	Timestamp          *time.Time      `json:"timestamp,omitempty"`
	TimestampEdited    *time.Time      `json:"timestampEdited,omitempty"`
	CallEndedTimestamp *time.Time      `json:"callEndedTimestamp,omitempty"`
	IsPinned           bool            `json:"isPinned"`
	Content            string          `json:"content"`
	Author             DceAuthor       `json:"author"`
	Attachments        []DceAttachment `json:"attachments"`
	Embeds             []DceEmbed      `json:"embeds"`
	Reactions          []DceReaction   `json:"reactions"`
	Mentions           []DceMention    `json:"mentions"`
	Reference          DceReference    `json:"reference,omitempty"`
}

type DceAttachment struct {
	ID            string `json:"id"`
	URL           string `json:"url"`
	FileName      string `json:"fileName"`
	FileSizeBytes int    `json:"fileSizeBytes"`
}

type DceEmbed struct {
	Title       string       `json:"title"`
	URL         string       `json:"url"`
	Timestamp   *time.Time   `json:"timestamp,omitempty"`
	Description string       `json:"description"`
	Thumbnail   DceThumbnail `json:"thumbnail"`
	Fields      []DceField   `json:"fields"`
}

type DceThumbnail struct {
	URL    string `json:"url"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

type DceField struct {
	Name     string `json:"name"`
	Value    string `json:"value"`
	IsInline bool   `json:"isInline"`
}

type DceReaction struct {
	Emoji DceEmoji `json:"emoji"`
	Count int      `json:"count"`
}
type DceEmoji struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	IsAnimated bool   `json:"isAnimated"`
	ImageURL   string `json:"imageUrl"`
}
