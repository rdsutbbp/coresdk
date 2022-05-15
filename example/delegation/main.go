package main

import (
	"context"
	"fmt"

	"github.com/rdsutbbp/logx"

	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
	delegationv1 "github.com/rdsutbbp/coresdk/typed/delegation/v1"
)

var clientset *coresdk.Clientset

func init() {
	clientset, _ = coresdk.NewClientWithOptions(
		rest.WithDefaultCoreRESTMode,
		rest.WithProtocol("http"),
		rest.WithCoreAddr("127.0.0.1"),
		rest.WithCoreListenPort("8090"),
		rest.WithXForwardedAuthUser(&rest.XForwardedAuthUser{
			UserID:  "test_user_id",
			GroupID: "test_group_id",
			Viewer:  "test_viewer",
		}),
	)
}

func main() {
	logx.Logger()
	ExecMachineCommandAPI()
}

func ExecMachineCommandAPI() {
	command, err := clientset.DelegationV1().Machine().ExecCommand(context.Background(), &delegationv1.ExecCommandRequest{
		MachineID: 49,
		Command:   "ls -a",
		Timeout:   10,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(command)
}

func MachineAPI() {
	init, err := clientset.DelegationV1().Machine().Init(context.Background(), &delegationv1.CoreMachine{
		NickName:     "test_nickname",
		HostIP:       "test_host_ip",
		VirtualIP:    "test_virtual_ip",
		CPU:          "test_cpu",
		Memory:       "test_memory",
		Disk:         "test_disk",
		Bandwidth:    "test_bandwidth",
		Args:         `{"name":"value"}`,
		SystemInfo:   "test_system_info",
		FullData:     `{"name":"value"}`,
		CredentialID: 0,
		Tag:          "test_tag",
		InstallPath:  "test_install_path",
	})
	if err != nil {
		logx.Errorf(err.Error())
		return
	}

	fmt.Println(init)

	err = clientset.DelegationV1().Machine().Update(context.Background(), &delegationv1.CoreMachine{
		ID:        init.ID,
		HostIP:    "test_host_ip_2",
		VirtualIP: "test_virtual_ip_2",
	})

	if err != nil {
		logx.Errorf(err.Error())
		return
	}

	query, err := clientset.DelegationV1().Machine().Query(context.Background(), init.ID)

	if err != nil {
		logx.Errorf(err.Error())
		return
	}
	fmt.Println(query)
}

func HostagentAPI() {
	logx.Logger()
	clientset, err := coresdk.NewClientWithOptions(
		rest.WithDefaultCoreRESTMode,
		rest.WithProtocol("http"),
		rest.WithCoreAddr("127.0.0.1"),
		rest.WithCoreListenPort("8090"),
		rest.WithXForwardedAuthUser(&rest.XForwardedAuthUser{
			UserID:  "test_user_id",
			GroupID: "test_group_id",
			Viewer:  "test_viewer",
		}),
	)

	if err != nil {
		fmt.Println(err)
		return
	}

	env, err := clientset.DelegationV1().Hostagent().QueryEnv(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(env)
}
