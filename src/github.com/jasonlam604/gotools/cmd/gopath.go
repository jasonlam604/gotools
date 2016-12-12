// Copyright © 2016 Jason Lam
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"path"
	"strings"
	"syscall"
)

// gopathCmd represents the gopath command
var gopathCmd = &cobra.Command{
	Use:   "gopath",
	Short: "Updates $GOPATH and $GOPATH/bin to current working directory",
	Long: `Updates $GOPATH and $GOPATH/bin to current working directory, 
	       it will overwrite the existing $GOPATH and remove the existing
	       $GOPATH/bin set in $PATH`,
	Run: func(cmd *cobra.Command, args []string) {

		loadPathInfo()

		clearGoPathBin()

		setGoPathAndBin()

		displayUpdates()

		updateShell()
	},
}

var goPathOrig string

func init() {
	RootCmd.AddCommand(gopathCmd)
}

func loadPathInfo() {
	goPathOrig = os.Getenv("GOPATH")
}

func clearGoPathBin() {

	if len(goPathOrig) > 0 {

		var goBin = path.Join(goPathOrig, "bin")

		if strings.Index(os.Getenv("PATH"), goBin) > -1 {

			var path string = os.Getenv("PATH")

			path = strings.Replace(path, goBin, "", -1)

			path = strings.Replace(path, "::", ":", -1)

			if strings.HasSuffix(path, ":") {
				path = strings.TrimRight(path, ":")
			}

			if strings.HasPrefix(path, ":") {
				path = strings.TrimLeft(path, ":")
			}

			os.Setenv("PATH", path)
		}
	}
}

func setGoPathAndBin() {

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	os.Setenv("GOPATH", pwd)
	os.Setenv("PATH", os.Getenv("PATH")+":"+path.Join(pwd, "bin"))
}

func displayUpdates() {
	fmt.Println("GOPATH: ", os.Getenv("GOPATH"))
	fmt.Println("PATH: ", os.Getenv("PATH"))
}

func updateShell() {
	syscall.Exec(os.Getenv("SHELL"), []string{os.Getenv("SHELL")}, syscall.Environ())
}
