package coresdk

import (
	"github.com/rdsutbbp/coresdk/rest"
	corev1 "github.com/rdsutbbp/coresdk/typed/core/v1"
	delegationv1 "github.com/rdsutbbp/coresdk/typed/delegation/v1"
	miniofsv1 "github.com/rdsutbbp/coresdk/typed/miniofs/v1"
)

type Clientset struct {
	delegationV1 *delegationv1.DelegationV1Client
	miniofsV1    *miniofsv1.MiniofsV1Client
	coreV1       *corev1.CoreV1Client
}

func (c *Clientset) DelegationV1() delegationv1.DelegationV1Interface {
	return c.delegationV1
}

func (c *Clientset) MiniofsV1() miniofsv1.MiniofsV1Interface {
	return c.miniofsV1
}

func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

func NewClientWithOptions(ops ...rest.Opt) (*Clientset, error) {
	c := &rest.RESTClient{}
	for _, op := range ops {
		if err := op(c); err != nil {
			return nil, err
		}
	}
	configShallowCopy := *c
	var cs Clientset
	var err error
	cs.delegationV1, err = delegationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.miniofsV1, err = miniofsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1, err = corev1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}
