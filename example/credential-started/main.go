package main

import (
	"context"

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
		rest.WithCoreListenPort("8001"),
	)

	credentialClient := config.DelegationV1().Credential()
	credentialClient.Init(context.Background(), &delegationv1.CoreCredential{
		Name:     "name",
		Type:     "type",
		Version:  "v1",
		FullData: "{}",
		Args:     "{}",
	})
}
