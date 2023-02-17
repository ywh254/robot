package dingtalk

import "testing"

var (
	token = ""
	key   = ""
)

func TestRobot_SendTextMessage(t *testing.T) {
	robot := New([]string{token}, key)

	err := robot.SendTextMessage("@all 大家好，我是测试机器人", WithAtAll(true))
	t.Fatal(err)
}

func TestRobot_SendLinkMessage(t *testing.T) {
	robot := New([]string{token}, key)

	err := robot.SendLinkMessage("这是一个测试", "@all 我是测试机器人!", "", "https://baidu.com", WithAtAll(true)))
	t.Fatal(err)
}

func TestRobot_SendMarkdownMessage(t *testing.T) {
	robot := New([]string{token}, key)

	err := robot.SendMarkdownMessage("这是一个测试", `# 标题
---
## 下面是
- 111
- 222
## 下面是
- 333
- 444
---
`, WithAtMembers([]string{"17888888888"}))
	t.Fatal(err)
}
