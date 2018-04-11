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
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBuckets(t *testing.T) {
	for i, example := range []struct {
		env       Env
		response  string
		buckets   []Bucket
		errString string
	}{
		{NewMockNoAuthEnv(), "", nil, "Unexpected response code: 401"},
		{NewMockBadPassEnv(), "", nil, "Unexpected response code: 401"},
		{NewMockEnv(), "[]", []Bucket{}, ""},
		{NewMockEnv(),
			`[
			  {
			    "id": "e3eca45f4d294132c07b49f4",
			    "name": "cqiNhd3Y16uXRBpRKbcGdrhVvouLRFlBM5O1jMUOr6OKJUVGpvv0LLaBv+6kqzyVvp5jFw==",
			    "created": "2016-10-12T14:40:21.259Z"
			  }
			]`, []Bucket{
				Bucket{
					ID:        "e3eca45f4d294132c07b49f4",
					Name:      "test",
					Created:   "2016-10-12T14:40:21.259Z",
					Decrypted: true,
				},
			}, ""},
		{NewMockNoMnemonicEnv(),
			`[
			  {
			    "id": "e3eca45f4d294132c07b49f4",
			    "name": "cqiNhd3Y16uXRBpRKbcGdrhVvouLRFlBM5O1jMUOr6OKJUVGpvv0LLaBv+6kqzyVvp5jFw==",
			    "created": "2016-10-12T14:40:21.259Z"
			  }
			]`, []Bucket{
				Bucket{
					ID:        "e3eca45f4d294132c07b49f4",
					Name:      "cqiNhd3Y16uXRBpRKbcGdrhVvouLRFlBM5O1jMUOr6OKJUVGpvv0LLaBv+6kqzyVvp5jFw==",
					Created:   "2016-10-12T14:40:21.259Z",
					Decrypted: false,
				},
			}, ""},
		{NewMockEnv(),
			`[
			  {
			    "id": "e3eca45f4d294132c07b49f4",
			    "name": "Yq3Ky6jJ7dwWiC9MEcb5nhAl5P0xfYe6jCwwwlzd1a1kZxKYLcft/WkOC8dhcwLb3Ka9xA==",
			    "created": "2016-10-12T14:40:21.259Z"
			  }
			]`, []Bucket{
				Bucket{
					ID:        "e3eca45f4d294132c07b49f4",
					Name:      "Yq3Ky6jJ7dwWiC9MEcb5nhAl5P0xfYe6jCwwwlzd1a1kZxKYLcft/WkOC8dhcwLb3Ka9xA==", // encrypted with a different mnemonic
					Created:   "2016-10-12T14:40:21.259Z",
					Decrypted: false,
				},
			}, ""},
		{NewMockEnv(),
			`[
			  {
			    "id": "e3eca45f4d294132c07b49f4",
			    "name": "test",
			    "created": "2016-10-12T14:40:21.259Z"
			  }
			]`, []Bucket{
				Bucket{
					ID:        "e3eca45f4d294132c07b49f4",
					Name:      "test", // unencrypted name
					Created:   "2016-10-12T14:40:21.259Z",
					Decrypted: false,
				},
			}, ""},
		{NewMockNoMnemonicEnv(),
			`[
			  {
			    "id": "e3eca45f4d294132c07b49f4",
			    "name": "test",
			    "created": "2016-10-12T14:40:21.259Z"
			  }
			]`, []Bucket{
				Bucket{
					ID:        "e3eca45f4d294132c07b49f4",
					Name:      "test", // unencrypted name
					Created:   "2016-10-12T14:40:21.259Z",
					Decrypted: false,
				},
			}, ""},
	} {
		mockGetBucketsResponse = example.response
		buckets, err := GetBuckets(example.env)
		errTag := fmt.Sprintf("Test case #%d", i)
		if example.errString != "" {
			assert.EqualError(t, err, example.errString, errTag)
			continue
		}
		if assert.NoError(t, err, errTag) {
			assert.Equal(t, example.buckets, buckets, errTag)
		}
	}
}
