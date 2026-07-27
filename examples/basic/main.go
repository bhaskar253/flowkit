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

package main

import (
	"fmt"

	"github.com/bhaskar253/flowkit"
)

func main() {
	counter := 1
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
			return value * value
		},

		func(value int) {
			fmt.Println(value)
		},
	)
	pipeline.Run()
}
