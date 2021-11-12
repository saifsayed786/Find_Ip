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
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/spf13/cobra"
)

// findIpCmd represents the findIp command
var findIpCmd = &cobra.Command{
	Use:   "findIp",
	Short: "Finds Ip",
	Long:  `Looks up the IP addresses for a particular host`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("findIp called")
		c := make(chan int)
		go getIp(c, args)

	},
}

func getIp(c chan int, args []string) {
	file, err := os.Open(args[0])
	if err != nil {
		panic("error opening file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		// fmt.Println(eachline)
		ip, err := net.LookupIP(eachline)
		if err != nil {
			panic(err)
		}
		for i := 0; i < len(ip); i++ {
			c <- i
			// fmt.Println(eachline, ":", ip[i])
			fmt.Println(eachline, ":", c)
		}
		close(c)

	}

}
func init() {
	rootCmd.AddCommand(findIpCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findIpCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findIpCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
