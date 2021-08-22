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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("random called")
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// randomCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// randomCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Joke struct {
	ID     string `json:"id"`
	Joke   string `json:"joke"`
	Status int    `json:"status"`
}

func getRandomJoke() {
	//fmt.Println("Get Random Joke: P")
	url := "https://icanhazdadjoke.com/"
	responseByte := getJokeData(url)
	joke := Joke{}
	if err := json.Unmarshal(responseByte, &joke); err != nil {
		log.Printf("Could not unmarshall response - %v", err)
	}
	fmt.Println(string(joke.Joke))
}

func getJokeData(baseApi string) []byte {
	request, err := http.NewRequest(http.MethodGet, baseApi, nil)

	if err != nil {
		log.Printf("could not request a dad joke -%v ", err)
	}
	request.Header.Add("Accept", "application/json")
	request.Header.Add("User-Agent", "Dad Joke CLI")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request -v %v", err)
	}
	responseByte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Could not read response body", err)
	}
	return responseByte
}
