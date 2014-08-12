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

package generator

import (
	"math/rand"
	"testing"

	generator "."
)

func TestCreateGenerator(t *testing.T) {
	_, err := generator.New(rand.New(rand.NewSource(1337)))
	if err != nil {
		t.Errorf("Failed to create generator")
	}
}

func TestExpressionValueGenerator(t *testing.T) {
	sampleGenerator, _ := generator.New(rand.New(rand.NewSource(1337)))

	var tests = []struct {
		Expression    string
		ExpectedValue string
	}{
		{"test[A-Z0-9]{4}template", "testQ3HVtemplate"},
		{"[\\d]{4}", "6841"},
		{"[\\w]{4}", "DVgK"},
		{"[\\a]{10}", "nFWmvmjuaZ"},
	}

	for _, test := range tests {
		value, _ := sampleGenerator.GenerateValue(test.Expression)
		if value != test.ExpectedValue {
			t.Errorf("Failed to generate expected value from %s\n. Generated: %s\n. Expected: %s\n", test.Expression, value, test.ExpectedValue)
		}
	}
}

func TestPasswordGenerator(t *testing.T) {
	sampleGenerator, _ := generator.New(rand.New(rand.NewSource(1337)))

	value, _ := sampleGenerator.GenerateValue("password")
	if value != "4U390O49" {
		t.Errorf("Failed to generate expected password. Generated: %s\n. Expected: %s\n", value, "4U390O49")
	}
}

func TestRemoteValueGenerator(t *testing.T) {
	sampleGenerator, _ := generator.New(rand.New(rand.NewSource(1337)))

	_, err := sampleGenerator.GenerateValue("[GET:http://api.example.com/new]")
	if err == nil {
		t.Errorf("Expected error while fetching non-existent remote.")
	}
}
