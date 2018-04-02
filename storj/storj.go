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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DefaultURL of the Storj Bridge API endpoint
const DefaultURL = "https://api.storj.io"

// Env contains parameters for accessing the Storj network
type Env struct {
	URL string // TODO set DefaultURL as default value
}

// Info structure of the GetInfo() response
type Info struct {
	Title       string
	Description string
	Version     string
	Host        string
}

// UnmarshalJSON overrides the unmarshalling for Info to correctly extract the data from the Swagger JSON response
func (info *Info) UnmarshalJSON(b []byte) error {
	var f interface{}
	json.Unmarshal(b, &f)

	m := f.(map[string]interface{})

	infomap := m["info"]
	v := infomap.(map[string]interface{})

	info.Title = v["title"].(string)
	info.Description = v["description"].(string)
	info.Version = v["version"].(string)

	info.Host = m["host"].(string)

	return nil
}

// GetInfo returns info about the Storj Bridge server
func GetInfo(env Env) (Info, error) {
	info := Info{}
	resp, err := http.Get(env.URL)
	if err != nil {
		return info, err
	}
	if resp.StatusCode != 200 {
		return info, fmt.Errorf("Unexpected response code: %d", resp.StatusCode)
	}
	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return info, err
	}
	err = json.Unmarshal(b, &info)
	return info, err
}
