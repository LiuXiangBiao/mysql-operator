package resources

import (
	"fmt"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/blinkbean/dingtalk"
	"time"
)

// 配置钉钉通知：有相关资源创建或更新机器人会发消息

func Notice(resourceKind, resourceName *string, msg string) {
	tokens := "f4de0ca30038d12d46d054dfc4ce3fea0624664a2611398a0f225adaa6b51d59"
	secret := "SEC687aefffd3bbda62b182fb214909325390be7fd3f966ce536843d587253236f9"
	d := dingtalk.InitDingTalkWithSecret(tokens, secret)
	mdmsg := []string{
		"### Kubernetes 资源警告信息",
		fmt.Sprintf("- Time：%v", time.Now().Format("2006-01-02 15:04")),
		fmt.Sprintf("- Kind：%s", tea.StringValue(resourceKind)),
		fmt.Sprintf("- Name：%s", tea.StringValue(resourceName)),
		fmt.Sprintf("- Message: %s", msg),
	}
	mobiles := []string{"刘向标"}
	err := d.SendMarkDownMessageBySlice("test1", mdmsg, dingtalk.WithAtAll(), dingtalk.WithAtMobiles(mobiles))
	if err != nil {
		return
	}
}
