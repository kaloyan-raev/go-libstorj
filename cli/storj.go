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

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Storj/go-libstorj/storj"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "storj"
	app.Version = "0.0.1"
	app.Usage = "command line interface to the Storj network"

	app.Commands = []cli.Command{
		{
			Name:      "get-info",
			Usage:     "prints bridge api information",
			ArgsUsage: " ", // no args
			Category:  "bridge api information",
			Action: func(c *cli.Context) error {
				getInfo()
				return nil
			},
		},
		{
			Name:      "list-buckets",
			Usage:     "lists the available buckets",
			ArgsUsage: " ", // no args
			Category:  "working with buckets and files",
			Action: func(c *cli.Context) error {
				listBuckets()
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getInfo() {
	env := storj.NewEnv()
	info, err := storj.GetInfo(env)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Storj bridge: %s\n\n", env.URL)
	fmt.Printf("Title:       %s\n", info.Title)
	fmt.Printf("Description: %s\n", info.Description)
	fmt.Printf("Version:     %s\n", info.Version)
	fmt.Printf("Host:        %s\n", info.Host)
}

func listBuckets() {
	buckets, err := storj.GetBuckets(storj.NewEnv())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	for _, b := range buckets {
		fmt.Printf("ID: %s\tDecrypted: %t\t\tCreated: %s\tName: %s\n",
			b.ID, b.Decrypted, b.Created, b.Name)
	}
}
