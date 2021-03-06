/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/martijnxd/myhue/functions"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all lights",
	Long: `myhue list lights
	  -w width
	  -r only reachable
	  -f "name" filter name 

	  examples:
	  myhue list -w
	  myhue list -w -r -f "light"

MYHue is a cli app to interact with a philips hue bridge.`,
	Run: func(cmd *cobra.Command, args []string) {
		token, bridge := functions.ConnectHUE()
		w, _ := cmd.Flags().GetBool("width")
		r, _ := cmd.Flags().GetBool("reachable")
		f, _ := cmd.Flags().GetString("filter")
		functions.ListAll(r, w, f, token, bridge)
	},
}

func init() {
	var w, r bool
	var f string
	rootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&w, "width", "w", false, "width list")
	listCmd.Flags().BoolVarP(&r, "reachable", "r", false, "display reachable lights only")
	listCmd.Flags().StringVarP(&f, "filter", "f", "", "specify keyword to filter name")
}
