package test

import (
	"fmt"
	"github.com/aceld/zinx/examples/zinx_version_ex/protoDemo/pb"
	"github.com/golang/protobuf/proto"
	"testing"
)

func TestPb(t *testing.T) {
	person := &pb.Person{
		Name:   "oo",
		Age:    18,
		Emails: []string{"a@qq.com", "b@163.com"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "123",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "456",
				Type:   pb.PhoneType_WORK,
			},
		},
	}

	b, _ := proto.Marshal(person)
	fmt.Println(b)
	newPerson := &pb.Person{}
	proto.Unmarshal(b, newPerson)
	fmt.Println(newPerson)
}
