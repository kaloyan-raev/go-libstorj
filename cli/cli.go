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

	"github.com/Storj/go-libstorj/storj"
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

	if len(os.Args) == 1 {
		// TODO print help
		fmt.Fprintln(os.Stderr, "No command specified")
		os.Exit(0)
	}

	cmd := os.Args[1]

	switch cmd {
	case "get-info":
		info, err := storj.GetInfo(storj.NewEnv())
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Title: %s\nDescription: %s\nVersion: %s\nHost: %s\n",
			info.Title, info.Description, info.Version, info.Host)

	default:
		// TODO print help
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", cmd)
		os.Exit(1)
	}
}
