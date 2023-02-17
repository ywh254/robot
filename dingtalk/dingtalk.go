package dingtalk

import (
	"github.com/blinkbean/dingtalk"
	"github.com/ywh254/robot"
)

var _ robot.Robot = (*Robot)(nil)

func WithAtAll(isAtAll bool) robot.Option {
	return func(r robot.Robot) {
		rbot := r.(*Robot)
		rbot.isAtAll = isAtAll
	}
}

func WithAtMembers(atMembers []string) func(r robot.Robot) {
	return func(r robot.Robot) {
		rbot := r.(*Robot)
		rbot.atMembers = atMembers
	}
}

type Robot struct {
	dingTalk *dingtalk.DingTalk

	isAtAll   bool
	atMembers []string
}

func New(tokens []string, key string) robot.Robot {
	return &Robot{
		dingTalk: dingtalk.InitDingTalk(tokens, key),
	}
}

func NewWithSecret(token string, secret string) robot.Robot {
	return &Robot{
		dingTalk: dingtalk.InitDingTalkWithSecret(token, secret),
	}
}

func (r *Robot) SendTextMessage(content string, opts ...robot.Option) error {
	r.applyOptions(opts...)

	if r.isAtAll {
		return r.dingTalk.SendTextMessage(content, dingtalk.WithAtAll())
	}

	return r.dingTalk.SendTextMessage(content, dingtalk.WithAtMobiles(r.atMembers))
}

func (r *Robot) SendLinkMessage(title, content, picURL, msgURL string, opts ...robot.Option) error {
	r.applyOptions(opts...)

	return r.dingTalk.SendLinkMessage(title, content, picURL, msgURL)
}

func (r *Robot) SendMarkdownMessage(title, content string, opts ...robot.Option) error {
	r.applyOptions(opts...)

	if r.isAtAll {
		return r.dingTalk.SendMarkDownMessage(title, content, dingtalk.WithAtAll())
	}

	return r.dingTalk.SendMarkDownMessage(title, content, dingtalk.WithAtMobiles(r.atMembers))
}

func (r *Robot) applyOptions(opts ...robot.Option) {
	for _, opt := range opts {
		opt(r)
	}
}
