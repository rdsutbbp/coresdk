package main

import (
	"context"
	"fmt"

	"github.com/rdsutbbp/logx"

	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
	delegationv1 "github.com/rdsutbbp/coresdk/typed/delegation/v1"
)

func main() {
	logx.Logger()
	MachineAPI()
}

func MachineAPI() {
	clientset, _ := coresdk.NewClientWithOptions(
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
	clientset, _ := coresdk.NewClientWithOptions(
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

	init, err := clientset.DelegationV1().Hostagent().Init(context.Background(), &delegationv1.CoreHostagent{})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(init)

	err = clientset.DelegationV1().Hostagent().Update(context.Background(), &delegationv1.CoreHostagent{
		ID:   init.ID,
		Addr: "10.1.40.108:9009",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	query, err := clientset.DelegationV1().Hostagent().Query(context.Background(), init.ID)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(query)
}
