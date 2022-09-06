package http

import (
	"os"
	"strings"

	"go.uber.org/fx"

	"github.com/fluxninja/aperture/pkg/config"
	"github.com/fluxninja/aperture/pkg/log"
)

const (
	defaultProxyKey = "client.proxy"
)

// ProxyModule returns the fx module that applies the provided proxy configuration.
func ProxyModule() fx.Option {
	return fx.Options(
		fx.Invoke(ProxyConstructor{}.applyProxyConfig),
	)
}

// ProxyConstructor holds fields used to configure Proxy.
type ProxyConstructor struct {
	Key           string
	DefaultConfig ProxyConfig
}

// ProxyConfig holds proxy configuration.
// This configuration has preference over environment variables HTTP_PROXY, HTTPS_PROXY or NO_PROXY. See <https://pkg.go.dev/golang.org/x/net/http/httpproxy#Config>
// swagger:model
// +kubebuilder:object:generate=true
type ProxyConfig struct {
	//+kubebuilder:validation:Optional
	HTTPProxy string `json:"http,omitempty" validate:"omitempty,url|hostname_port"`

	//+kubebuilder:validation:Optional
	HTTPSProxy string `json:"https,omitempty" validate:"omitempty,url|hostname_port"`

	//+kubebuilder:validation:Optional
	NoProxy []string `json:"no_proxy,omitempty" validate:"dive,ip|cidr|fqdn|hostname_port"`
}

func (constructor ProxyConstructor) applyProxyConfig(unmarshaller config.Unmarshaller) error {
	if constructor.Key == "" {
		constructor.Key = defaultProxyKey
	}

	var err error

	config := constructor.DefaultConfig
	if err = unmarshaller.UnmarshalKey(constructor.Key, &config); err != nil {
		log.Error().Err(err).Msg("Unable to deserialize client proxy configuration!")
		return err
	}

	if config.HTTPProxy != "" ||
		config.HTTPSProxy != "" ||
		len(config.NoProxy) > 0 {
		err = os.Setenv("HTTP_PROXY", config.HTTPProxy)
		if err != nil {
			return err
		}
		err = os.Setenv("HTTPS_PROXY", config.HTTPSProxy)
		if err != nil {
			return err
		}
		err = os.Setenv("NO_PROXY", strings.Join(config.NoProxy, ","))
		if err != nil {
			return err
		}
	}
	return nil
}
