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

// Map transforms each value into another type.
func Map[T any, R any](
	values []T,
	fn func(T) R,
) []R {
	result := make([]R, len(values))
	for i, value := range values {
		result[i] = fn(value)
	}
	return result
}

// Filter keeps only values matching the predicate.
func Filter[T any](
	values []T,
	predicate func(T) bool,
) []T {
	result := make([]T, 0)
	for _, value := range values {
		if predicate(value) {
			result = append(
				result,
				value,
			)
		}
	}
	return result
}

// Reduce combines all values into a single result.
func Reduce[T any, R any](
	values []T,
	initial R,
	fn func(R, T) R,
) R {
	result := initial
	for _, value := range values {
		result = fn(
			result,
			value,
		)
	}
	return result
}

// Fold is an alias for Reduce.
func Fold[T any, R any](
	values []T,
	initial R,
	fn func(R, T) R,
) R {
	return Reduce(
		values,
		initial,
		fn,
	)
}

// Tap executes a side effect and returns the original value.
func Tap[T any](
	value T,
	fn func(T),
) T {
	fn(value)
	return value
}

// Collect consumes an iterator into a slice.
func Collect[T any](
	iterator Iterator[T],
) []T {
	result := make([]T, 0)
	iterator.ForEach(
		func(value T) {
			result = append(
				result,
				value,
			)
		},
	)
	return result
}

// Chunk splits a slice into smaller groups.
func Chunk[T any](
	values []T,
	size int,
) [][]T {
	if size <= 0 {
		return nil
	}
	result := make(
		[][]T,
		0,
	)
	for i := 0; i < len(values); i += size {
		end := i + size
		if end > len(values) {
			end = len(values)
		}
		result = append(
			result,
			values[i:end],
		)
	}
	return result
}
