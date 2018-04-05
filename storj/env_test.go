// Copyright (C) 2018 Storj Labs Inc.
//
// This file is part of go-libstorj.
//
// go-libstorj is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// go-libstorj is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with go-libstorj.  If not, see <http://www.gnu.org/licenses/>.

package storj

import "testing"

var envTests = []struct {
	env         Env
	expectedURL string
}{
	{Env{}, ""},
	{NewEnv(), DefaultURL},
	{Env{URL: MockBridgeURL}, MockBridgeURL},
}

func TestNewEnv(t *testing.T) {
	for _, tt := range envTests {
		if tt.env.URL != tt.expectedURL {
			t.Errorf("URL is incorrect, got: %s, want: %s", tt.env.URL, tt.expectedURL)
		}
	}
}
