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

package template

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/apiserver"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/labels"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/template/templateapi"
	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"
)

type TemplateConfigStorage struct{}

// NewRegistryStorage returns a new RegistryStorage.
func NewRegistryStorage() apiserver.RESTStorage {
	return &TemplateConfigStorage{}
}

func (s *TemplateConfigStorage) List(selector labels.Selector) (interface{}, error) {
	return nil, fmt.Errorf("TemplateConfig can only be created.")
}

func (s *TemplateConfigStorage) Get(id string) (interface{}, error) {
	return nil, fmt.Errorf("TemplateConfig can only be created.")
}

func (s *TemplateConfigStorage) New() interface{} {
	return &templateapi.TemplateConfig{}
}

func (s *TemplateConfigStorage) Delete(id string) (<-chan interface{}, error) {
	return apiserver.MakeAsync(func() (interface{}, error) {
		return nil, fmt.Errorf("TemplateConfig can only be created.")
	}), nil
}

func (s *TemplateConfigStorage) Update(minion interface{}) (<-chan interface{}, error) {
	return nil, fmt.Errorf("TemplateConfig can only be created.")
}

func (s *TemplateConfigStorage) Create(obj interface{}) (<-chan interface{}, error) {
	t, ok := obj.(*templateapi.TemplateConfig)
	if !ok {
		return nil, fmt.Errorf("not a template: %#v", obj)
	}
	if t.ID == "" {
		return nil, fmt.Errorf("ID should not be empty: %#v", t)
	}

	t.CreationTimestamp = util.Now()

	GenerateParameterValues(t, rand.New(rand.NewSource(time.Now().UnixNano())))
	ProcessEnvParameters(t)

	return apiserver.MakeAsync(func() (interface{}, error) {
		return t, nil
	}), nil
}
