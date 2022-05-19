package v1

import (
	"context"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"

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

	// parse response data
	// code message data
	j, err := simplejson.NewJson(result.Body())
	if err != nil {
		return nil, err
	}
	code, err := j.Get("code").Int()
	if err != nil {
		return nil, err
	}
	if code != http.StatusOK {
		message, _ := j.Get("message").String()
		return nil, fmt.Errorf(message)
	}

	return result.Body(), nil
}
