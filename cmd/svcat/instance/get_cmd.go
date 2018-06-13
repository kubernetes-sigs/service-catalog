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

package instance

import (
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/command"
	"github.com/kubernetes-incubator/service-catalog/cmd/svcat/output"
	"github.com/spf13/cobra"
)

type getCmd struct {
	*command.NamespacedCommand
	name         string
	outputFormat string
}

func (c *getCmd) SetFormat(format string) {
	c.outputFormat = format
}

// NewGetCmd builds a "svcat get instances" command
func NewGetCmd(cxt *command.Context) *cobra.Command {
	getCmd := &getCmd{NamespacedCommand: command.NewNamespacedCommand(cxt)}
	cmd := &cobra.Command{
		Use:     "instances [NAME]",
		Aliases: []string{"instance", "inst"},
		Short:   "List instances, optionally filtered by name",
		Example: command.NormalizeExamples(`
  svcat get instances
  svcat get instances --all-namespaces
  svcat get instance wordpress-mysql-instance
  svcat get instance -n ci concourse-postgres-instance
`),
		PreRunE: command.PreRunE(getCmd),
		RunE:    command.RunE(getCmd),
	}
	getCmd.AddNamespaceFlags(cmd.Flags(), true)
	command.AddOutputFlags(cmd.Flags())
	return cmd
}

func (c *getCmd) Validate(args []string) error {
	if len(args) > 0 {
		c.name = args[0]
	}

	return nil
}

func (c *getCmd) Run() error {
	if c.name == "" {
		return c.getAll()
	}

	return c.get()
}

func (c *getCmd) getAll() error {
	instances, err := c.App.RetrieveInstances(c.Namespace)
	if err != nil {
		return err
	}

	output.WriteInstanceList(c.Output, c.outputFormat, instances)
	return nil
}

func (c *getCmd) get() error {
	instance, err := c.App.RetrieveInstance(c.Namespace, c.name)
	if err != nil {
		return err
	}

	output.WriteInstance(c.Output, c.outputFormat, *instance)

	return nil
}
