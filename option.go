/*
 *
 *     Copyright 2020 yunqi
 *
 *     Licensed under the Apache License, Version 2.0 (the "License");
 *     you may not use this file except in compliance with the License.
 *     You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 *     Unless required by applicable law or agreed to in writing, software
 *     distributed under the License is distributed on an "AS IS" BASIS,
 *     WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *     See the License for the specific language governing permissions and
 *     limitations under the License.
 *
 */

package flow

import "math"

// DefaultAntsPoolSize 默认goroutine池的默认容量
const DefaultAntsPoolSize = math.MaxInt32

type Options struct {
	poolSize    int
	disablePool bool // 默认启用池
}

type Option func(options *Options)

func loadOptions(options ...Option) *Options {
	op := new(Options)
	for _, option := range options {
		option(op)
	}
	if op.poolSize <= 0 {
		op.poolSize = 10000
	}
	return op
}
func WithOption(options *Options) Option {
	return func(ops *Options) {
		ops = options
	}
}
func WithPoolSize(size int) Option {
	return func(options *Options) {
		options.poolSize = size
	}
}
func WithDisablePool() Option {
	return func(options *Options) {
		options.disablePool = true
	}
}
func WithEnablePool(enable bool) Option {
	return func(options *Options) {
		options.disablePool = !enable
	}
}
