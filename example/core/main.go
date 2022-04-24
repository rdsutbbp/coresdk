package main

import (
	"context"
	"fmt"

	"github.com/rdsutbbp/logx"

	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
	corev1 "github.com/rdsutbbp/coresdk/typed/core/v1"
)

func main() {
	logx.Logger()
	clientset, _ := coresdk.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithCoreAddr("127.0.0.1"),
		rest.WithCoreListenPort("8090"),
		rest.WithXForwardedAuthUser(&rest.XForwardedAuthUser{
			UserID:  "test_user_id",
			GroupID: "test_group_id",
			Viewer:  "test_viewer",
		}),
	)
	list, err := clientset.CoreV1().Delegation().List(context.Background(), &corev1.PageParam{
		Page: 2,
		Size: 10,
		All:  false,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(list)
}
