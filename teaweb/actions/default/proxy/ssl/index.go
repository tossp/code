package ssl

import (
	"github.com/TeaWeb/code/teaconfigs"
	"github.com/iwind/TeaGo/actions"
)

type IndexAction actions.Action

// SSL设置
func (this *IndexAction) Run(params struct {
	ServerId string
}) {
	server := teaconfigs.NewServerConfigFromId(params.ServerId)
	if server == nil {
		this.Fail("找不到Server")
	}

	this.Data["selectedTab"] = "https"
	this.Data["server"] = server

	this.Show()
}
