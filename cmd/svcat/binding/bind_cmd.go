/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package binding

import (
	"fmt"
	"time"

	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/parameters"
	"github.com/spf13/cobra"
)

type bindCmd struct {
	*command.Namespaced
	instanceName string
	bindingName  string
	externalID   string
	secretName   string
	rawParams    []string
	jsonParams   string
	params       interface{}
	rawSecrets   []string
	secrets      map[string]string
	wait         bool
	rawTimeout   string
	timeout      *time.Duration
}

// NewBindCmd builds a "svcat bind" command
func NewBindCmd(cxt *command.Context) *cobra.Command {
	bindCmd := &bindCmd{Namespaced: command.NewNamespacedCommand(cxt)}
	cmd := &cobra.Command{
		Use:   "bind INSTANCE_NAME",
		Short: "Binds an instance's metadata to a secret, which can then be used by an application to connect to the instance",
		Example: `
  svcat bind wordpress
  svcat bind wordpress-mysql-instance --name wordpress-mysql-binding --secret-name wordpress-mysql-secret
  svcat bind wordpress-mysql-instance --name wordpress-mysql-binding --external-id c8ca2fcc-4398-11e8-842f-0ed5f89f718b
  svcat bind wordpress-instance --params type=admin
  svcat bind wordpress-instance --params-json '{
	"type": "admin",
	"teams": [
		"news",
		"weather",
		"sports"
	]
}
'
`,
		PreRunE: command.PreRunE(bindCmd),
		RunE:    command.RunE(bindCmd),
	}
	command.AddNamespaceFlags(cmd.Flags(), false)
	cmd.Flags().StringVarP(
		&bindCmd.bindingName,
		"name",
		"",
		"",
		"The name of the binding. Defaults to the name of the instance.",
	)
	cmd.Flags().StringVar(&bindCmd.externalID, "external-id", "",
		"The ID of the binding for use with OSB API (Optional)",
	)
	cmd.Flags().StringVarP(
		&bindCmd.secretName,
		"secret-name",
		"",
		"",
		"The name of the secret. Defaults to the name of the instance.",
	)
	cmd.Flags().StringSliceVarP(&bindCmd.rawParams, "param", "p", nil,
		"Additional parameter to use when binding the instance, format: NAME=VALUE. Cannot be combined with --params-json")
	cmd.Flags().StringSliceVarP(&bindCmd.rawSecrets, "secret", "s", nil,
		"Additional parameter, whose value is stored in a secret, to use when binding the instance, format: SECRET[KEY]")
	cmd.Flags().StringVar(&bindCmd.jsonParams, "params-json", "",
		"Additional parameters to use when binding the instance, provided as a JSON object. Cannot be combined with --param")
	cmd.Flags().BoolVar(&bindCmd.wait, "wait", false,
		"Wait until the operation completes.")
	cmd.Flags().StringVar(&bindCmd.rawTimeout, "timeout", "5m",
		"Timeout for --wait, specified in human readable format: 30s, 1m, 1h. Specify -1 to wait indefinitely.")
	return cmd
}

func (c *bindCmd) Validate(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("instance is required")
	}
	c.instanceName = args[0]

	var err error

	if c.jsonParams != "" && len(c.rawParams) > 0 {
		return fmt.Errorf("--params-json cannot be used with --param")
	}

	if c.jsonParams != "" {
		c.params, err = parameters.ParseVariableJSON(c.jsonParams)
		if err != nil {
			return fmt.Errorf("invalid --params-json value (%s)", err)
		}
	} else {
		c.params, err = parameters.ParseVariableAssignments(c.rawParams)
		if err != nil {
			return fmt.Errorf("invalid --param value (%s)", err)
		}
	}

	c.secrets, err = parameters.ParseKeyMaps(c.rawSecrets)
	if err != nil {
		return fmt.Errorf("invalid --secret value (%s)", err)
	}

	if c.wait && c.rawTimeout != "-1" {
		timeout, err := time.ParseDuration(c.rawTimeout)
		if err != nil {
			return fmt.Errorf("invalid --timeout value (%s)", err)
		}
		c.timeout = &timeout
	}

	return nil
}

func (c *bindCmd) Run() error {
	return c.bind()
}

func (c *bindCmd) bind() error {
	binding, err := c.App.Bind(c.Namespace, c.bindingName, c.externalID, c.instanceName, c.secretName, c.params, c.secrets)
	if err != nil {
		return err
	}

	if c.wait {
		pollInterval := 1 * time.Second
		binding, err = c.App.WaitForBinding(binding.Namespace, binding.Name, pollInterval, c.timeout)
	}

	output.WriteBindingDetails(c.Output, binding)
	return err
}
