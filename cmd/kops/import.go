/*
Copyright 2016 The Kubernetes Authors.

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

package main

import (
	"github.com/spf13/cobra"
	"k8s.io/kubernetes/pkg/kubectl/cmd/templates"
	"k8s.io/kubernetes/pkg/util/i18n"
)

var (
	import_long = templates.LongDesc(i18n.T(`
	Imports a kubernetes cluster created by kube-up.sh into a state store.  This command
	only support AWS clusters at this time.`))

	import_example = templates.Examples(i18n.T(`
	# Import a cluser
	kops import cluster --name k8s-cluster.example.com --region us-east-1 \
	  --state=s3://k8s-cluster.example.com`))

	import_short = i18n.T(`Import a cluster.`)
)

// importCmd represents the import command
var importCmd = &cobra.Command{
	Use:     "import",
	Short:   import_short,
	Long:    import_long,
	Example: import_example,
}

func init() {
	rootCommand.AddCommand(importCmd)
}
