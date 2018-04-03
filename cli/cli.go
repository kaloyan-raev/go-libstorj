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
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func getInfo() {
	info, err := storj.GetInfo(storj.NewEnv())
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Title: %s\nDescription: %s\nVersion: %s\nHost: %s\n",
		info.Title, info.Description, info.Version, info.Host)
}
