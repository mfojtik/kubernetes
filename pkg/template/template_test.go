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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"testing"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/template/templateapi"
)

func TestNewTemplate(t *testing.T) {
	var template templateapi.TemplateConfig

	jsonData, _ := ioutil.ReadFile("example/project.json")
	if err := json.Unmarshal(jsonData, &template); err != nil {
		t.Errorf("Unable to process the JSON template file: %v", err)
	}
}

func TestCustomParameter(t *testing.T) {
	var template templateapi.TemplateConfig

	jsonData, _ := ioutil.ReadFile("example/project.json")
	json.Unmarshal(jsonData, &template)

	AddCustomTemplateParameter(templateapi.Parameter{Name: "CUSTOM_PARAM", Value: "1"}, &template)
	AddCustomTemplateParameter(templateapi.Parameter{Name: "CUSTOM_PARAM", Value: "2"}, &template)

	if p := GetTemplateParameterByName("CUSTOM_PARAM", &template); p == nil {
		t.Errorf("Unable to add a custom parameter to the template")
	} else {
		if p.Value != "2" {
			t.Errorf("Unable to replace the custom parameter value in template")
		}
	}

}

func ExampleProcessTemplateParameters() {
	var template templateapi.TemplateConfig

	jsonData, _ := ioutil.ReadFile("example/project.json")
	json.Unmarshal(jsonData, &template)

	// Define custom parameter for transformation:
	customParam := templateapi.Parameter{Name: "CUSTOM_PARAM1", Value: "1"}
	AddCustomTemplateParameter(customParam, &template)

	// Generate parameter values
	GenerateParameterValues(&template, rand.New(rand.NewSource(1337)))

	// Substitute parameters with values in container env vars
	ProcessEnvParameters(&template)

	result, _ := json.Marshal(template)
	fmt.Println(string(result))
	// Output:
	// {"kind":"templateConfig","id":"guestbook","creationTimestamp":null,"parameters":[{"creationTimestamp":null,"name":"ADMIN_USERNAME","description":"Guestbook administrator username","type":"string","generate":"admin[A-Z0-9]{3}","value":"adminQ3H"},{"creationTimestamp":null,"name":"ADMIN_PASSWORD","description":"Guestboot administrator password","type":"string","generate":"[a-zA-Z0-9]{8}","value":"dwNJiJwW"},{"creationTimestamp":null,"name":"REDIS_PASSWORD","description":"The password Redis use for communication","type":"string","generate":"[a-zA-Z0-9]{8}","value":"P8vxbV4C"},{"creationTimestamp":null,"name":"CUSTOM_PARAM1","description":"","type":"","value":"1"}],"services":[{"kind":"Service","id":"frontend","creationTimestamp":null,"apiVersion":"v1beta1","port":5432,"selector":{"name":"frontend"},"containerPort":0},{"kind":"Service","id":"redismaster","creationTimestamp":null,"apiVersion":"v1beta1","port":10000,"selector":{"name":"redis-master"},"containerPort":0},{"kind":"Service","id":"redisslave","creationTimestamp":null,"apiVersion":"v1beta1","port":10001,"labels":{"name":"redisslave"},"selector":{"name":"redisslave"},"containerPort":0}],"pods":[{"kind":"Pod","id":"redis-master-2","creationTimestamp":null,"apiVersion":"v1beta1","labels":{"name":"redis-master"},"desiredState":{"manifest":{"version":"v1beta1","id":"redis-master-2","volumes":null,"containers":[{"name":"master","image":"dockerfile/redis","ports":[{"containerPort":6379}],"env":[{"name":"REDIS_PASSWORD","value":"P8vxbV4C"}]}]},"restartpolicy":{}},"currentState":{"manifest":{"version":"","id":"","volumes":null,"containers":null},"restartpolicy":{}}}],"replicationControllers":[{"kind":"ReplicationController","id":"frontendController","creationTimestamp":null,"apiVersion":"v1beta1","desiredState":{"replicas":3,"replicaSelector":{"name":"frontend"},"podTemplate":{"desiredState":{"manifest":{"version":"v1beta1","id":"frontendController","volumes":null,"containers":[{"name":"php-redis","image":"brendanburns/php-redis","ports":[{"hostPort":8000,"containerPort":80}],"env":[{"name":"ADMIN_USERNAME","value":"adminQ3H"},{"name":"ADMIN_PASSWORD","value":"dwNJiJwW"},{"name":"REDIS_PASSWORD","value":"P8vxbV4C"}]}]},"restartpolicy":{}},"labels":{"name":"frontend"}}},"labels":{"name":"frontend"}},{"kind":"ReplicationController","id":"redisSlaveController","creationTimestamp":null,"apiVersion":"v1beta1","desiredState":{"replicas":2,"replicaSelector":{"name":"redisslave"},"podTemplate":{"desiredState":{"manifest":{"version":"v1beta1","id":"redisSlaveController","volumes":null,"containers":[{"name":"slave","image":"brendanburns/redis-slave","ports":[{"hostPort":6380,"containerPort":6379}],"env":[{"name":"REDIS_PASSWORD","value":"P8vxbV4C"}]}]},"restartpolicy":{}},"labels":{"name":"redisslave"}}},"labels":{"name":"redisslave"}}]}
	//
}
