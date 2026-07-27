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

package flowkit

// Compose creates a transformation chain.
//
// Functions execute from left to right.
func Compose[T any](
	functions ...func(T) T,
) func(T) T {
	return func(value T) T {
		for _, fn := range functions {
			value = fn(value)
		}
		return value
	}
}

// Pipe applies a value through multiple functions.
func Pipe[T any](
	value T,
	functions ...func(T) T,
) T {
	for _, fn := range functions {
		value = fn(value)
	}
	return value
}
