package v1

import (
	"context"
	"fmt"

	"github.com/rdsutbbp/coresdk/rest"
)

type ObjectGetter interface {
	Object() ObjectInterface
}

type ObjectInterface interface {
	GetBytes(ctx context.Context, bucket, object string) ([]byte, error)
}

type object struct {
	client rest.Interface
}

func newObject(c *MiniofsV1Client) *object {
	return &object{
		client: c.RESTClient(),
	}
}

func (o *object) GetBytes(ctx context.Context, bucket, object string) ([]byte, error) {
	result := o.client.Get().
		SubPath(fmt.Sprintf("/minio/api/v1/get/bucket/%s/object/%s", bucket, object)).
		Do(ctx)

	if result.Error() != nil {
		return nil, result.Error()
	}
	return result.Body(), nil
}
