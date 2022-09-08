/*
Copyright 2022 The KubeVela Authors.

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

package definition

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"
)

// InstallDefinitionFromYAML will install definition from yaml file
func InstallDefinitionFromYAML(ctx context.Context, cli client.Client, defPath string, replace func(string) string) error {
	b, err := ioutil.ReadFile(defPath)
	if err != nil {
		return err
	}
	s := string(b)
	if replace != nil {
		s = replace(s)
	}
	defJson, err := yaml.YAMLToJSON([]byte(s))
	if err != nil {
		return err
	}
	u := &unstructured.Unstructured{}
	if err := json.Unmarshal(defJson, u); err != nil {
		return err
	}
	return cli.Create(ctx, u.DeepCopy())
}
