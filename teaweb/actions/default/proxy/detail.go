package proxy

import (
	"github.com/TeaWeb/code/teaconfigs"
	"github.com/iwind/TeaGo/actions"
)

type DetailAction actions.Action

// 代理详情
func (this *DetailAction) Run(params struct {
	ServerId string
}) {
	server := teaconfigs.NewServerConfigFromId(params.ServerId)
	if server == nil {
		this.Fail("找不到Server")
	}

	if server.Index == nil {
		server.Index = []string{}
	}

	this.Data["selectedTab"] = "basic"
	this.Data["server"] = server

	this.Show()
}
