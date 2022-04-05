package main

import (
	"context"
	"fmt"
	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
	delegationv1 "github.com/rdsutbbp/coresdk/typed/delegation/v1"
	"github.com/rdsutbbp/logx"
)

func main() {
	logx.Logger()
	config, _ := coresdk.NewClientWithOptions(
		rest.WithDefaultCoreRESTMode,
		rest.WithProtocol("http"),
		rest.WithCoreAddr("127.0.0.1"),
		rest.WithCoreListenPort("8090"),
		rest.WithXForwardedAuthUser(&rest.XForwardedAuthUser{
			UserId:  "test_user_id",
			GroupId: "test_group_id",
			Viewer:  "test_viewer",
		}),
	)

	init, err := config.DelegationV1().Machine().Init(context.Background(), &delegationv1.InitMachineBody{
		NickName:     "test_nickname11",
		HostIP:       "test_host_ip",
		VirtualIP:    "test_virtual_ip",
		CPU:          "test_cpu",
		Memory:       "test_mem",
		Disk:         "test_disk",
		Bandwidth:    "test_bandwidth",
		Args:         "{}",
		SystemInfo:   "test_system",
		FullData:     "{}",
		CredentialID: 1,
		Tag:          "test_tag",
		InstallPath:  "test_install_path",
	})
	if err != nil {
		return
	}
	fmt.Println(init)
}
