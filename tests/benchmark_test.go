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

func BenchmarkMap(b *testing.B) {
	values := make(
		[]int,
		10000,
	)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowkit.Map(
			values,
			func(value int) int {
				return value * 2
			},
		)
	}
}

func BenchmarkReduce(b *testing.B) {
	values := make([]int, 10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		flowkit.Reduce(
			values,
			0,
			func(total int, value int) int {
				return total + value
			},
		)
	}
}
