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

// DefaultAntsPoolSize The default capacity of the default goroutine pool
const DefaultAntsPoolSize = 10000

// Options Options
type Options struct {
	poolSize    int
	disablePool bool // Pool is enabled by default
}

type Option func(options *Options)

// loadOptions return a Options
func loadOptions(options ...Option) *Options {
	op := new(Options)
	for _, option := range options {
		option(op)
	}
	// set the default pool size
	if op.poolSize <= 0 {
		op.poolSize = DefaultAntsPoolSize
	}
	return op
}

// WithOption return a Option interface
func WithOption(options *Options) Option {
	return func(ops *Options) {
		ops = options
	}
}

// WithPoolSize return a Option that set pool size
func WithPoolSize(size int) Option {
	return func(options *Options) {
		options.poolSize = size
	}
}

// WithDisablePool return a Option with pool closed
func WithDisablePool() Option {
	return func(options *Options) {
		options.disablePool = true
	}
}

// WithEnablePool return a Option
func WithEnablePool(enable bool) Option {
	return func(options *Options) {
		options.disablePool = !enable
	}
}
