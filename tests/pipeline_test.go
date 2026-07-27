// Copyright 2026 Bhaskar
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package flowkit_test

import (
	"testing"

	"github.com/bhaskar253/flowkit"
)

func TestPipeline(t *testing.T) {
	counter := 1
	total := 0

	pipeline := flowkit.New(
		func() (int, bool) {
			if counter > 5 {
				return 0, false
			}

			value := counter
			counter++
			return value, true
		},

		func(value int) int {
			return value * 2
		},

		func(value int) {
			total += value
		},
	)

	pipeline.Run()
	expected := 30

	if total != expected {
		t.Fatalf(
			"expected %d got %d",
			expected,
			total,
		)
	}
}

func TestMap(t *testing.T) {
	values := []int{1, 2, 3}
	result := flowkit.Map(
		values,
		func(value int) int {
			return value * 2
		},
	)

	expected := []int{2, 4, 6}

	for i := range expected {
		if result[i] != expected[i] {
			t.Fatalf(
				"expected %v got %v",
				expected,
				result,
			)
		}
	}
}

func TestFilter(t *testing.T) {

	values := []int{1, 2, 3, 4}

	result := flowkit.Filter(
		values,
		func(value int) bool {
			return value%2 == 0
		},
	)

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 results got %d",
			len(result),
		)
	}
}

func TestReduce(t *testing.T) {
	values := []int{1, 2, 3, 4}

	result := flowkit.Reduce(
		values,
		0,
		func(total int, value int) int {
			return total + value
		},
	)

	if result != 10 {
		t.Fatalf(
			"expected 10 got %d",
			result,
		)
	}
}
