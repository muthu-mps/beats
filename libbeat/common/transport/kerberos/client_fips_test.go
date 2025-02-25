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

//go:build requirefips

package kerberos

import (
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	cfg, err := os.CreateTemp(t.TempDir(), "config")
	require.NoError(t, err)
	c, err := NewClient(&Config{
		AuthType:   authPassword,
		ConfigPath: cfg.Name(),
	}, http.DefaultClient, "")
	require.Nil(t, c)
	require.EqualError(t, err, "kerberos is not supported in fips mode")
}
