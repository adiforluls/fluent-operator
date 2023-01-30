package custom

import (
	"github.com/fluent/fluent-operator/apis/fluentd/v1alpha1/plugins"
	"github.com/fluent/fluent-operator/apis/fluentd/v1alpha1/plugins/params"
	"github.com/fluent/fluent-operator/pkg/utils"
)

type CustomPlugin struct {
	Config string `json:"config,omitempty"`
}

func (c *CustomPlugin) Name() string {
	return ""
}

func (c *CustomPlugin) Params(_ plugins.SecretLoader) (*params.PluginStore, error) {
	ps := params.NewPluginStore("")
	ps.Content = utils.Indentation(c.Config)
	return ps, nil
}
