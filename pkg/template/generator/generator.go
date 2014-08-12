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
	"sync"
)

// GeneratorInterface is an abstract interface for generating
// random values from an input expression
type GeneratorInterface interface {
	GenerateValue(expression string) (interface{}, error)
}

// Generator implements GeneratorInterface
type Generator struct {
	seed *rand.Rand
}

func New(seed *rand.Rand) (Generator, error) {
	return Generator{seed: seed}, nil
}

// GenerateValue loops over registered generators and tries to find the
// one matching the input expression. If match is found, it then generates
// value using that matching generator
func (g *Generator) GenerateValue(expression string) (interface{}, error) {
	if len(generators) <= 0 {
		return expression, fmt.Errorf("No registered generators.")
	}

	for regexp, generatorFactory := range generators {
		if regexp.FindStringIndex(expression) != nil {
			generator, err := generatorFactory(g.seed)
			if err != nil {
				return expression, err
			}
			return generator.GenerateValue(expression)
		}
	}

	return expression, fmt.Errorf("No matching generators found.")
}

// GeneratorFactory is an abstract factory for creating generators
// (objects that implement GeneratorInterface interface)
type GeneratorFactory func(*rand.Rand) (GeneratorInterface, error)

// generators stores registered generators
var generators = make(map[*regexp.Regexp]GeneratorFactory)
var generatorsMutex sync.Mutex

// RegisterGenerator registers new generator for a certain input expression
func RegisterGenerator(r *regexp.Regexp, f GeneratorFactory) {
	generatorsMutex.Lock()
	defer generatorsMutex.Unlock()
	generators[r] = f
}
