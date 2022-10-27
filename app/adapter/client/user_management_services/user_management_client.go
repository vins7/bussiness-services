package user_management_services

import (
	"github.com/vins7/bussiness-services/app/util"
	cfgServer "github.com/vins7/bussiness-services/config"
	proto "github.com/vins7/module-protos/app/interface/grpc/proto/user_management"
	"google.golang.org/grpc"
)

type UserManagementClient struct {
	client proto.UsermanagementServiceClient
}

func NewUserManagementClient(c *grpc.ClientConn) UserManagementRepo {
	return &UserManagementClient{
		client: proto.NewUsermanagementServiceClient(c),
	}
}

func (u *UserManagementClient) Login(in interface{}) error {
	_, err := u.client.Login(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.LoginRequest{
		Username: "aaaa",
		Password: "aaaa",
	})
	if err != nil {
		return nil
	}
	return nil
}

func (u *UserManagementClient) Register(in interface{}) error {
	_, err := u.client.Login(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.LoginRequest{
		Username: "aaaa",
		Password: "aaaa",
	})
	if err != nil {
		return nil
	}
	return nil
}

func (u *UserManagementClient) UserInformation(in interface{}) error {
	_, err := u.client.Login(util.BuildContext(cfgServer.GetConfig().Server.App.Name), &proto.LoginRequest{
		Username: "aaaa",
		Password: "aaaa",
	})
	if err != nil {
		return nil
	}
	return nil
}