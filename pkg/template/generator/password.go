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
	"fmt"
	"math/rand"
	"regexp"
)

// PasswordGenerator generates string of 8 random alphanumeric characters
// from an input expression matching "password" string.
//
// Example:
//   - "password" => "hW4yQU5i"
type PasswordGenerator struct {
	expressionValueGenerator ExpressionValueGenerator
}

var passwordExp = regexp.MustCompile(`password`)

func init() {
	RegisterGenerator(passwordExp, func(seed *rand.Rand) (GeneratorInterface, error) { return newPasswordGenerator(seed) })
}

func newPasswordGenerator(seed *rand.Rand) (PasswordGenerator, error) {
	return PasswordGenerator{ExpressionValueGenerator{seed: seed}}, nil
}

func (g PasswordGenerator) GenerateValue(string) (interface{}, error) {
	return g.expressionValueGenerator.GenerateValue(fmt.Sprintf("[\\a]{%d}", 8))
}
