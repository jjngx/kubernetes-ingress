// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/nginxinc/kubernetes-ingress/v3/pkg/client/clientset/versioned/typed/configuration/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeK8sV1alpha1 struct {
	*testing.Fake
}

func (c *FakeK8sV1alpha1) GlobalConfigurations(namespace string) v1alpha1.GlobalConfigurationInterface {
	return &FakeGlobalConfigurations{c, namespace}
}

func (c *FakeK8sV1alpha1) Policies(namespace string) v1alpha1.PolicyInterface {
	return &FakePolicies{c, namespace}
}

func (c *FakeK8sV1alpha1) TransportServers(namespace string) v1alpha1.TransportServerInterface {
	return &FakeTransportServers{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeK8sV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
