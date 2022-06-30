package engine

import (
	"github.com/limitcool/push/engine/mirai"
	"github.com/limitcool/push/engine/pushplus"
	. "github.com/limitcool/push/global"
)

type Pusher interface {
	Send(content string) error
	Name() string
	EditTitle(title string)
	GetTitle() string
}

type Nop struct{}

func (Nop) Send(string) error {
	return nil
}

func (Nop) Name() string {
	return "nop"
}

func (Nop) EditTitle(string) {}
func (Nop) GetTitle() string {
	return "nop"
}
func Select(platform string) Pusher {
	switch platform {
	case "Mirai":
		return mirai.New(Config.Mirai.Host, Config.Mirai.Target)
	case "PushPlus":
		return pushplus.New(Config.PushPlus.Token, Config.PushPlus.Title)
	default:
		return Nop{}
	}
}
