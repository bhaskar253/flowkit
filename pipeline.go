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

// Pipeline represents:
//
// Producer -> Transformer -> Consumer
//
// It reads values from a producer,
// transforms them,
// and sends them to a consumer.
type Pipeline[T any, R any] struct {
	producer    Producer[T]
	transformer Transformer[T, R]
	consumer    Consumer[R]
}

// New creates a new pipeline.
func New[T any, R any](
	producer Producer[T],
	transformer Transformer[T, R],
	consumer Consumer[R],
) *Pipeline[T, R] {
	return &Pipeline[T, R]{
		producer:    producer,
		transformer: transformer,
		consumer:    consumer,
	}
}

// Run executes the pipeline.
//
// Execution stops when the producer returns false.
func (p *Pipeline[T, R]) Run() {
	for {
		value, ok := p.producer()
		if !ok {
			return
		}
		result := p.transformer(value)
		p.consumer(result)
	}
}
