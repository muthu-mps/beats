// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package dissect

import (
	"fmt"
	"strings"
)

type trimMode byte

const (
	trimModeNone trimMode = iota
	trimModeRight
	trimModeLeft
	trimModeAll = trimModeRight | trimModeLeft
)

type config struct {
	Tokenizer     *tokenizer `config:"tokenizer" validate:"required"`
	Field         string     `config:"field"`
	TargetPrefix  string     `config:"target_prefix"`
	IgnoreFailure bool       `config:"ignore_failure"`
	OverwriteKeys bool       `config:"overwrite_keys"`
	TrimValues    trimMode   `config:"trim_values"`
	TrimChars     string     `config:"trim_chars"`
}

var defaultConfig = config{
	Field:        "message",
	TargetPrefix: "dissect",
	TrimChars:    " ",
}

// tokenizer add validation at the unpack level for this specific field.
type tokenizer = Dissector

// Unpack a tokenizer into a dissector this will trigger the normal validation of the dissector.
func (t *tokenizer) Unpack(v string) error {
	d, err := New(v)
	if err != nil {
		return err
	}
	*t = *d
	return nil
}

// Unpack the trim mode from a string.
func (tm *trimMode) Unpack(v string) error {
	switch strings.ToLower(v) {
	case "", "none":
		*tm = trimModeNone
	case "left":
		*tm = trimModeLeft
	case "right":
		*tm = trimModeRight
	case "all", "both":
		*tm = trimModeAll
	default:
		return fmt.Errorf("unsupported value %s. Must be one of [none, left, right, all]", v)
	}
	return nil
}
