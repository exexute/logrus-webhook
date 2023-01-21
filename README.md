# logrus.webhook

dingtalk、feiao.rebot notify for logrus.

## install

```shell
go get -u github.com/exexute/logrus-webhook@v0.1.1
```

## Send Log To [DingTalk Robot](https://open.dingtalk.com/document/group/call-robot-api-operations)

### DingTalk Robot Msg Supports

- [x] text msg (default)
- [x] link msg
- [x] markdown msg

### DingTalk Robot At Supports

- [x] at all
- [x] at with mobile
- [x] at with empId

### Examples

**send msg and at all member**

```go
package main

import (
	logrusWebhook "github.com/exexute/logrus-webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	dingTalkHook, err := logrusWebhook.NewDingTalkHook(logrusWebhook.DefaultDingTalkApi, "<ding_robot_token>", "<ding_robot_secret>", logrus.WarnLevel)
	if err == nil {
		logrus.AddHook(dingTalkHook)
	}

	logrus.WithFields(
		logrus.Fields{
			"at":      logrusWebhook.NewDingTalkAt(true, nil),
			"msgType": "text",
		},
	).Warn("this is a test msg.")
}

```

**send text msg to robot**

```go
package main

import (
	logrusWebhook "github.com/exexute/logrus-webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	dingTalkHook, err := logrusWebhook.NewDingTalkHook(logrusWebhook.DefaultDingTalkApi, "<ding_robot_token>", "<ding_robot_secret>", logrus.WarnLevel)
	if err == nil {
		logrus.AddHook(dingTalkHook)
	}

	logrus.WithFields(
		logrus.Fields{
			"at":      logrusWebhook.NewDingTalkAt(false, []string{"<member-mobile>"}),
			"msgType": "text",
		},
	).Warn("this is a test msg.")
}

```

**send link msg to robot**

```go
package main

import (
	logrusWebhook "github.com/exexute/logrus-webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	dingTalkHook, err := logrusWebhook.NewDingTalkHook(logrusWebhook.DefaultDingTalkApi, "<ding_robot_token>", "<ding_robot_secret>", logrus.WarnLevel)
	if err == nil {
		logrus.AddHook(dingTalkHook)
	}

	logrus.WithFields(
		logrus.Fields{
			"at":      logrusWebhook.NewDingTalkAt(false, []string{"<member-mobile>"}),
			"msgType": "link",
		},
	).Warn(logrusWebhook.NewDingTalkLink("link msg", "Cheng Xiang is a singer I like very much", "https://img.mp.itc.cn/q_70,c_zoom,w_640/upload/20170615/c37f702fb76e4e64aaa12a85e6b0ae43_th.jpg", "https://baike.baidu.com/item/%E7%A8%8B%E5%93%8D/6058905").String())
}

```

**send markdown msg to robot**

```go
package main

import (
	logrusWebhook "github.com/exexute/logrus-webhook"
	"github.com/sirupsen/logrus"
)

func main() {
	dingTalkHook, err := logrusWebhook.NewDingTalkHook(logrusWebhook.DefaultDingTalkApi, "<ding_robot_token>", "<ding_robot_secret>", logrus.WarnLevel)
	if err == nil {
		logrus.AddHook(dingTalkHook)
	}

	logrus.WithFields(
		logrus.Fields{
			"at":      logrusWebhook.NewDingTalkAt(false, []string{"<member-mobile>"}),
			"msgType": "markdown",
		},
	).Warn(logrusWebhook.NewDingTalkMarkdown("Cheng Xiang", "[Cheng Xiang](https://baike.baidu.com/item/%E7%A8%8B%E5%93%8D/6058905) is a singer I like very much. ![](https://img.mp.itc.cn/q_70,c_zoom,w_640/upload/20170615/c37f702fb76e4e64aaa12a85e6b0ae43_th.jpg)").String())
}

```

## Send Log To [Aliyun Sls](https://help.aliyun.com/document_detail/48869.html)

### Examples

**send msg to sls logStore**

```go

```