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

// Iterator represents a lazy data source.
type Iterator[T any] struct {
	next Producer[T]
}

// NewIterator creates a new iterator.
func NewIterator[T any](
	producer Producer[T],
) Iterator[T] {
	return Iterator[T]{
		next: producer,
	}
}

// Next returns the next value.
func (i Iterator[T]) Next() (T, bool) {
	return i.next()
}

// ForEach consumes all iterator values.
func (i Iterator[T]) ForEach(
	consumer Consumer[T],
) {
	for {
		value, ok := i.Next()
		if !ok {
			return
		}
		consumer(value)
	}
}
