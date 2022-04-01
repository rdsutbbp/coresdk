package main

import (
	"github.com/rdsutbbp/coresdk"
	"github.com/rdsutbbp/coresdk/rest"
)

func main() {
	config, _ := coresdk.NewClientWithOptions(
		rest.WithProtocol("http"),
		rest.WithCoreAddr("127.0.0.1"),
		rest.WithCoreListenPort("8001"),
	)

	credentialClient := config.DelegationV1().Credential()
	credentialClient.Init()
}
