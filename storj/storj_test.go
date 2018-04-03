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

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"
)

const (
	MockBridgeAddr  = "localhost:8091"
	MockBridgeURL   = "http://" + MockBridgeAddr
	MockTitle       = "Storj Bridge"
	MockDescription = "Some description"
	MockVersion     = "1.2.3"
	MockHost        = "1.2.3.4"
)

func TestMain(m *testing.M) {
	mockBridge()
	os.Exit(m.Run())
}

func mockBridge() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, fmt.Sprintf("Cannot GET %s", r.URL.Path), 404)
			return
		}

		fmt.Fprintf(w, `{"info":{"title":"%s","description":"%s","version":"%s"},"host":"%s"}`,
			MockTitle, MockDescription, MockVersion, MockHost)
	})
	go http.ListenAndServe(MockBridgeAddr, nil)
	// TODO better way to wait for the mock server to start listening
	time.Sleep(1000)
}

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

var unmarshalTests = []struct {
	raw           string
	expectedError bool
}{
	{"", true},
	{"{", true}, // syntax error
	{"{}", true},
	{`{"info":{}}`, true},
	{fmt.Sprintf(`{"info":{description":"%s","version":"%s"},"host":"%s"}`,
		MockDescription, MockVersion, MockHost), true},
	{fmt.Sprintf(`{"info":{"title":"%s","version":"%s"},"host":"%s"}`,
		MockTitle, MockVersion, MockHost), true},
	{fmt.Sprintf(`{"info":{"title":"%s","description":"%s"},"host":"%s"}`,
		MockTitle, MockDescription, MockHost), true},
	{fmt.Sprintf(`{"info":{"title":"%s","description":"%s","version":"%s"}}`,
		MockTitle, MockDescription, MockVersion), true},
	{fmt.Sprintf(`{"info":{"title":"%s","description":"%s","version":"%s"},"host":"%s"}`,
		MockTitle, MockDescription, MockVersion, MockHost), false},
}

func TestUnmarshalJSON(t *testing.T) {

	for _, tt := range unmarshalTests {
		var info Info
		err := json.Unmarshal([]byte(tt.raw), &info)

		if err != nil {
			if !tt.expectedError {
				t.Errorf("Unexpected error: %s", err)
			}
			continue
		}

		checkInfo(info, t)
	}
}

var getInfoTests = []struct {
	env           Env
	expectedError bool
}{
	{Env{URL: MockBridgeURL}, false},
	{Env{URL: MockBridgeURL + "/info"}, true},
}

func TestGetInfo(t *testing.T) {
	for _, tt := range getInfoTests {
		info, err := GetInfo(tt.env)

		if err != nil {
			if !tt.expectedError {
				t.Errorf("Unexpected error: %s", err)
			}
			continue
		}

		checkInfo(info, t)
	}
}

func checkInfo(info Info, t *testing.T) {
	if info.Title != MockTitle {
		t.Errorf("Title is incorrect, got: %s, want: %s", info.Title, MockTitle)
	}
	if info.Description != MockDescription {
		t.Errorf("Description is incorrect, got: %s, want: %s", info.Description, MockDescription)
	}
	if info.Version != MockVersion {
		t.Errorf("Version is incorrect, got: %s, want: %s", info.Version, MockVersion)
	}
	if info.Host != MockHost {
		t.Errorf("Host is incorrect, got: %s, want: %s", info.Host, MockHost)
	}
}
