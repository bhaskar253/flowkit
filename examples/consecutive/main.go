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
	nums := []int{1, 1, 0, 1, 1, 1, 0, 1}
	index := 0
	max := 0
	pipeline := flowkit.New(
		func() (int, bool) {
			for index < len(nums) {
				if nums[index] == 0 {
					index++
					continue
				}
				start := index
				for index < len(nums) && nums[index] == 1 {
					index++
				}
				return index - start, true
			}
			return 0, false
		},
		func(value int) int {
			return value
		},
		func(value int) {
			if value > max {
				max = value
			}
		},
	)
	pipeline.Run()
	fmt.Println(
		"max consecutive ones:",
		max,
	)
}
