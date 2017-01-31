/**
 * Copyright (c) 2017 CA. All rights reserved.
 * This software may be modified and distributed under the terms
 * of the MIT license. See the LICENSE file for details.
 */

package cmd

import (
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/bevanhunt/gowrex"
	"github.com/dixonwille/wmenu"
	"github.com/gosuri/uitable"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// flag value for limit
var limit int

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Keyword search apis.io",
	Long:  `Keyword search apis.io. Example: search trade`,
	Run: func(cmd *cobra.Command, args []string) {
		type JSONReceive struct {
			Status string `json:"status"`
			Data   []struct {
				CreatedAt   *time.Time `json:"createdAt"`
				Name        string     `json:"name"`
				Description string     `json:"description"`
				Image       string     `json:"image"`
				BaseURL     string     `json:"baseURL"`
				HumanURL    string     `json:"humanURL"`
				Properties  []struct {
					Type string `json:"type"`
					URL  string `json:"url"`
				} `json:"properties"`
				Contact []struct {
					FN               string `json:"FN"`
					Email            string `json:"email"`
					OrganizationName string `json:"organizationName"`
					XTwitter         string `json:"X-twitter"`
				} `json:"contact"`
				Authoritative bool       `json:"authoritative"`
				APIFileURL    string     `json:"apiFileUrl"`
				Slug          string     `json:"slug"`
				UpdatedAt     *time.Time `json:"updatedAt"`
				Tags          []string   `json:"tags,omitempty"`
			} `json:"data"`
			Limit  int `json:" limit"`
			Skip   int `json:" skip"`
			Paging struct {
				Next     string `json:"next"`
				Previous string `json:"previous"`
			} `json:" paging"`
		}

		// join all args into one keyword seperated by spaces and url encoded
		keywords := strings.Join(args, " ")
		keywordsEscaped := url.QueryEscape(keywords)

		// make API request to apis.io
		req, err := gowrex.Request{
			URI: "http://apis.io/api/search?q=" + keywordsEscaped + "&limit=" + strconv.Itoa(limit),
		}.GetJSON()
		if err != nil {
			log.Println(err)
		}
		res, err := req.Do()
		if err != nil {
			log.Println(err)
		}
		resp := &JSONReceive{}
		_, err = res.JSON(resp)
		if err != nil {
			log.Println(err)
		}

		// create results top bar
		fmt.Println(chalk.Red.Color("Search: " + keywords))
		fmt.Println(chalk.Yellow.Color("Results: "))
		fmt.Println("")

		// handle no results
		if len(resp.Data) == 1 && resp.Data[0].Name == "" {
			fmt.Println(chalk.Red.Color("No Results"))
			return
		}

		// create results table
		table := uitable.New()
		table.MaxColWidth = 50
		table.Wrap = true

		table.AddRow(
			"#",
			"Name",
			"Description",
			"Swagger",
		)
		for i, el := range resp.Data {
			swagger := false
			for _, elem := range el.Properties {
				if elem.Type == "Swagger" {
					swagger = true
				}
			}
			table.AddRow("")
			table.AddRow(
				i,
				el.Name,
				el.Description,
				swagger,
			)
		}
		fmt.Println(table)

		// site viewing menu
		actFunc := func(opt wmenu.Opt) error {
			selection, ok := opt.Value.(int)
			if ok {
				open.Run(resp.Data[selection].HumanURL)
			}
			return nil
		}
		menu := wmenu.NewMenu("Select to show site:")
		menu.Action(actFunc)
		for x, elem := range resp.Data {
			menu.Option(elem.Name, x, false, nil)
		}
		menuErr := menu.Run()
		if menuErr != nil {
			log.Fatal(menuErr)
		}

		// count swagger items
		swaggerCount := 0
		for _, el := range resp.Data {
			swagger := false
			for _, elem := range el.Properties {
				if elem.Type == "Swagger" {
					swagger = true
				}
			}
			if swagger == true {
				swaggerCount++
			}
		}

		// if more than one swagger item then show menu
		if swaggerCount > 0 {
			// swagger viewing menu
			swFunc := func(opt wmenu.Opt) error {
				url, ok := opt.Value.(string)
				if ok {
					open.Run("http://petstore.swagger.io?url=" + url)
				}
				return nil
			}
			swMenu := wmenu.NewMenu("Select to display Swagger:")
			swMenu.Action(swFunc)
			for _, elem := range resp.Data {
				var url string
				swaggered := false
				for _, el := range elem.Properties {
					if el.Type == "Swagger" {
						url = el.URL
						swaggered = true
					}
				}
				if swaggered == true {
					swMenu.Option(elem.Name, url, false, nil)
				}
			}
			swMenuErr := swMenu.Run()
			if swMenuErr != nil {
				log.Fatal(swMenuErr)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("limit", "l", "# of max results (limit)")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	searchCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit number of results - default 10")

}
