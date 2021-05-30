/*
MIT License

Copyright (c) 2021 Kazuki Hara

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/
package cmd

import (
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/hk220/go-circle-list-extract/event"
	"github.com/hk220/go-circle-list-extract/parser"
	"github.com/hk220/go-circle-list-extract/printer"
	"github.com/spf13/afero"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use:   "go-circle-list-extract [event name]",
		Short: "go-circle-list-extract extracts a circle list of Comitia in JSON or CSV format.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Validate flags
			if !printer.HasPrinter(viper.GetString("format")) {
				log.Fatal("invalid format")
			}

			// Parse events
			var events map[string]event.Event
			if err := viper.UnmarshalKey("event", &events); err != nil {
				log.Fatal(err)
			}

			// event key check
			event, ok := events[args[0]]
			if !ok {
				log.Fatalf("no event key: %s", args[0])
			}

			// Request the HTML Page
			client := &http.Client{}
			resp, err := client.Get(event.CircleListURL)
			if err != nil {
				log.Fatal(err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				log.Fatalf("status code error: %d %s", resp.StatusCode, resp.Status)
			}

			// Load the HTML document
			doc, err := goquery.NewDocumentFromReader(resp.Body)
			if err != nil {
				log.Fatal(err)
			}

			// Parse HTML document
			prsr := parser.GetParser(event.Parser)
			cl, err := prsr.Parse(doc)
			if err != nil {
				log.Fatal(err)
			}

			// Write File
			var appFs = afero.NewOsFs()

			file, err := appFs.OpenFile(viper.GetString("output"), os.O_RDWR|os.O_CREATE, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			prnt := printer.GetPrinter(viper.GetString("format"))
			err = prnt.Print(file, cl)
			if err != nil {
				log.Fatal(err)
			}
		},
	}
)

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringP("format", "f", "csv", "output format (support: json, csv)")
	rootCmd.PersistentFlags().StringP("output", "o", "circles.csv", "output file name")
	viper.BindPFlag("format", rootCmd.PersistentFlags().Lookup("format"))
	viper.BindPFlag("output", rootCmd.PersistentFlags().Lookup("output"))
}

func initConfig() {
	viper.SetConfigFile(cfgFile)
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
