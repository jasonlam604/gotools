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
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"net"
)

// myipCmd represents the myip command
var myipCmd = &cobra.Command{
	Use:   "myip",
	Short: "Finds your local IP, ignores loopback",
	Long:  `Finds your local IP, ignores loopback, not much more to say about this`,
	Run: func(cmd *cobra.Command, args []string) {

		ip, err := externalIP()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(ip)

	},
}

func init() {
	RootCmd.AddCommand(myipCmd)
}

// Source from http://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}
