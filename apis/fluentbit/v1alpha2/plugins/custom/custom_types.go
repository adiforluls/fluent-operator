package custom

import (
	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins"
	"github.com/fluent/fluent-operator/apis/fluentbit/v1alpha2/plugins/params"
	"github.com/fluent/fluent-operator/pkg/utils"
)

// +kubebuilder:object:generate:=true

// CustomPlugin is used to support filter plugins that are not implemented yet
type CustomPlugin struct {
	Config string `json:"config,omitempty"`
}

func (c *CustomPlugin) Name() string {
	return ""
}

func (a *CustomPlugin) Params(_ plugins.SecretLoader) (*params.KVs, error) {
	kvs := params.NewKVs()
	kvs.Content = utils.Indentation(a.Config)
	return kvs, nil
}
