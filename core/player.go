package core

import (
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zinx_app_demo/mmo_game/pb"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"sync"
)

type Player struct {
	Conn ziface.IConnection
	Pid  int32   // 玩家id
	X    float32 // 平面的x坐标
	Y    float32 // 高度
	Z    float32 // 平面y坐标
	V    float32 // 旋转的高度0-360
}

// 全局的玩家id
var globalPid int32 = 1
var pidLock sync.Mutex

// 新建一个用户
func NewPlayer(conn ziface.IConnection) *Player {
	pidLock.Lock()
	id := globalPid
	globalPid++
	pidLock.Unlock()

	return &Player{
		Conn: conn,
		Pid:  id,
		X:    float32(rand.Intn(10) + 100),
		Y:    0,
		Z:    float32(rand.Intn(20) + 100),
		V:    0,
	}
}

// 给玩家发送消息
func (p *Player) SendMsg(msgId uint32, msg proto.Message) {
	data, err := proto.Marshal(msg)
	if err != nil {
		fmt.Println("marshal msg err: ", err)
		return
	}
	if p.Conn == nil {
		fmt.Println("player conn is nil")
		return
	}
	p.Conn.SendMsg(msgId, data)
}

// 同步玩家数据给客户端
func (p *Player) SyncPid() {
	msg := &pb.SyncPid{Pid: p.Pid}
	p.SendMsg(1, msg)
}

// 广播玩家的位置变化
func (p *Player) BroadcastPosition() {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  0, // 广播玩家的位置
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}
	p.SendMsg(2, msg)
}
