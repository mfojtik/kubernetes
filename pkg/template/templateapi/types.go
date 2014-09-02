/*
Copyright 2014 Google Inc. All rights reserved.

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

package templateapi

import "github.com/GoogleCloudPlatform/kubernetes/pkg/api"

type TemplateConfig struct {
	Config     `json:",inline" yaml:",inline"`
	Parameters []Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

type Config struct {
	api.JSONBase           `json:",inline" yaml:",inline"`
	Name                   string                      `yaml:"name" json:"name"`
	Description            string                      `yaml:"description" json:"description"`
	Services               []api.Service               `json:"services,omitempty" yaml:"services,omitempty"`
	Pods                   []api.Pod                   `json:"pods,omitempty" yaml:"pods,omitempty"`
	ReplicationControllers []api.ReplicationController `json:"replicationControllers,omitempty" yaml:"replicationControllers,omitempty"`

	// TODO: Add these as soon as the buildconfigapi.BuildConfig is available
	// BuildConfigs      []buildconfigapi.BuildConfig      `json:"buildConfigs" yaml:"buildConfigs"`
	// ImageRepositories []ImageRepository  `json:"imageRepositories" yaml:"imageRepositories"`
	// TODO: Add this as soon as the deployapi.DeploymentConfig is available
	// DeploymentConfigs []deployapi.DeploymentConfig `json:"deploymentConfigs" yaml:"deploymentConfigs"`
}

type Parameter struct {
	Name        string `json:"name" yaml:"name"`
	Description string `json:"description" yaml:"description"`
	Type        string `json:"type" yaml:"type"`
	Generate    string `json:"generate,omitempty" yaml:"generate,omitempty"`
	Value       string `json:"value,omitempty" yaml:"value,omitempty"`
}

func init() {
	api.AddKnownTypes("",
		TemplateConfig{},
		Config{},
	)

	api.AddKnownTypes("v1beta1",
		TemplateConfig{},
		Config{},
	)
}
