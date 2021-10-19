package protos

import (
	"context"
	"log"
	"os"

	protobuf "github.com/ali2210/didactic-octo-eureka/protos"
	"google.golang.org/grpc"
)

const _ = grpc.SupportPackageIsVersion7

type Grpc_client struct {
	Grpc_client_conn grpc.ClientConnInterface
}

type GRPC_Client_Services interface {
	CreateFile(ctx context.Context, action *protobuf.Action, opts ...grpc.CallOption) *protobuf.Action
	CopyContent(ctx context.Context, event *protobuf.Program_Event, opts ...grpc.CallOption) *protobuf.Program_Event
}

func Addons(self grpc.ClientConnInterface) GRPC_Client_Services {
	return &Grpc_client{Grpc_client_conn: self}
}

func (self *Grpc_client) CreateFile(ctx context.Context, action *protobuf.Action, opts ...grpc.CallOption) *protobuf.Action {

	event := new(protobuf.Action)
	fs, err := os.OpenFile(action.Event+".Dockerfile", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Println("File Creation Failed", err.Error(), event.Event)
		event.Message_Action = protobuf.Action_ERROR
		return event
	}

	fsinfo, err := fs.Stat()
	if err != nil {
		log.Println("File Creation Failed", err.Error(), event.Event)
		event.Message_Action = protobuf.Action_ERROR
		return event
	}
	event.Event = fsinfo.Name()
	event.Message_Action = protobuf.Action_OK
	return event
}

func (self *Grpc_client) CopyContent(ctx context.Context, event *protobuf.Program_Event, opts ...grpc.CallOption) *protobuf.Program_Event {
	// out_evnt = new(protobuf.Program_Event)

}
