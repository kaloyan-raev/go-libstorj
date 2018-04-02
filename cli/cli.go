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
package main

import (
	"flag"
	"fmt"
	"os"

	"storj.io/storj"
)

const version = "0.0.1"

func main() {
	// TODO support long flag names, i.e. --help, --version, etc.
	h := flag.Bool("h", false, "output usage information")
	v := flag.Bool("v", false, "output the version number")

	flag.Parse()

	if *h {
		// TODO print help
		os.Exit(0)
	}

	if *v {
		fmt.Printf("libstorj cli %s\n", version)
		os.Exit(0)
	}

	cmd := os.Args[1]

	env := storj.Env{
		URL: storj.DefaultURL,
	}

	switch cmd {
	case "get-info":
		info, err := storj.GetInfo(env)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		} else {
			fmt.Printf("title: %s\ndescription: %s\nversion: %s\nhost:%s\n",
				info.Title, info.Description, info.Version, info.Host)
		}
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
	}
}
