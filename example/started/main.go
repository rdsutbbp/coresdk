package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/rdsutbbp/logx"

	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
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

	env, err := config.DelegationV1().Hostagent().QueryPkg(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(env.AMD64)

	bytes, err := config.MiniofsV1().Object().GetBytes(context.Background(), "delegations", "virtualbox-v1_0_9-test.tar.gz")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = ioutil.WriteFile("virtualbox-v1_0_9-test.tar.gz", bytes, 777)
	if err != nil {
		fmt.Println(err)
		return
	}
}
