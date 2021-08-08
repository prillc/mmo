package main

import (
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/znet"
	"github.com/prillc/mmo/core"
)

func main() {
	s := znet.NewServer()
	// 连接创建和销毁的钩子函数
	s.SetOnConnStart(func(connection ziface.IConnection) {
		player := core.NewPlayer(connection)
		// 给客户端同步玩家id
		player.SyncPid()
		// 给客户端同步玩家坐标
		player.BroadcastPosition()

	})
	// 注册路由
	s.Serve()
}
