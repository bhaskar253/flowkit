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

// Stream provides chainable collection operations.
type Stream[T any] struct {
	values []T
}

// From creates a stream from a slice.
func From[T any](
	values []T,
) Stream[T] {
	return Stream[T]{
		values: values,
	}
}

// Map transforms stream values.
func (s Stream[T]) Map(
	fn func(T) T,
) Stream[T] {
	return Stream[T]{
		values: Map(
			s.values,
			fn,
		),
	}
}

// Filter removes values that do not match.
func (s Stream[T]) Filter(
	fn func(T) bool,
) Stream[T] {
	return Stream[T]{
		values: Filter(
			s.values,
			fn,
		),
	}
}

// Values returns stream contents.
func (s Stream[T]) Values() []T {
	return s.values
}

// Reduce aggregates stream values.
func (s Stream[T]) Reduce(
	initial T,
	fn func(T, T) T,
) T {
	return Reduce(
		s.values,
		initial,
		fn,
	)
}
