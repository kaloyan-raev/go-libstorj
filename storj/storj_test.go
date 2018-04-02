/*
 * Copyright (C) 2018 Storj Labs Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */
package storj

import (
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
		} else {
			fmt.Fprintf(w, `{"info":{"title":"%s","description":"%s","version":"%s"},"host":"%s"}`,
				MockTitle, MockDescription, MockVersion, MockHost)
		}
	})
	go http.ListenAndServe(MockBridgeAddr, nil)
	// TODO better way to wait for the mock server to start listening
	time.Sleep(1000)
}

func TestGetInfo(t *testing.T) {
	env := Env{URL: MockBridgeURL}
	info, err := GetInfo(env)

	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	} else {
		if info.Title != MockTitle {
			t.Errorf("Title is incorrected, got: %s, want: %s", info.Title, MockTitle)
		}
		if info.Description != MockDescription {
			t.Errorf("Description is incorrected, got: %s, want: %s", info.Description, MockDescription)
		}
		if info.Version != MockVersion {
			t.Errorf("Version is incorrected, got: %s, want: %s", info.Version, MockVersion)
		}
		if info.Host != MockHost {
			t.Errorf("Host is incorrected, got: %s, want: %s", info.Host, MockHost)
		}
	}
}

func TestGetInfoBadPath(t *testing.T) {
	env := Env{URL: MockBridgeURL + "/info"}
	_, err := GetInfo(env)

	if err == nil {
		t.Error("Expected error, but call was successful")
	}
}
