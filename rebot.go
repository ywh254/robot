package robot

type Option func(r Robot)

type Robot interface {
	SendTextMessage(content string, opts ...Option) error

	SendLinkMessage(title, content, picURL, msgURL string, opts ...Option) error

	SendMarkdownMessage(title, content string, opts ...Option) error
}
