// this is handmade crd & controller

// Code generated by client-gen. DO NOT EDIT.

package v1

import (
	"net/http"

	v1 "github.com/trainyao/playground/k8s/crd-controller/handmade/pkg/apis/handmade/v1"
	"github.com/trainyao/playground/k8s/crd-controller/handmade/pkg/client/clientset/versioned/scheme"
	rest "k8s.io/client-go/rest"
)

type HandmadeV1Interface interface {
	RESTClient() rest.Interface
	GoddessMomentsGetter
}

// HandmadeV1Client is used to interact with features provided by the handmade.crd.playground.trainyao.io group.
type HandmadeV1Client struct {
	restClient rest.Interface
}

func (c *HandmadeV1Client) GoddessMoments(namespace string) GoddessMomentInterface {
	return newGoddessMoments(c, namespace)
}

// NewForConfig creates a new HandmadeV1Client for the given config.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*HandmadeV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	httpClient, err := rest.HTTPClientFor(&config)
	if err != nil {
		return nil, err
	}
	return NewForConfigAndClient(&config, httpClient)
}

// NewForConfigAndClient creates a new HandmadeV1Client for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
func NewForConfigAndClient(c *rest.Config, h *http.Client) (*HandmadeV1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientForConfigAndClient(&config, h)
	if err != nil {
		return nil, err
	}
	return &HandmadeV1Client{client}, nil
}

// NewForConfigOrDie creates a new HandmadeV1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *HandmadeV1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new HandmadeV1Client for the given RESTClient.
func New(c rest.Interface) *HandmadeV1Client {
	return &HandmadeV1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = scheme.Codecs.WithoutConversion()

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *HandmadeV1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
