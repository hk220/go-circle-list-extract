package main

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
		Use:   "go-circle-list-extract",
		Short: "go-circle-list-extract extracts a circle list of Comitia in JSON or CSV format.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			// Validate flags
			if viper.GetString("format") != "json" && viper.GetString("format") != "csv" {
				log.Fatal("invalid format")
			}

			client := &http.Client{}

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
			prsr := parser.GetParser(args[0])
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

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "config.yaml", "config file (default is config.yaml)")
	rootCmd.PersistentFlags().StringP("format", "f", "csv", "output format")
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

func main() {
	cobra.CheckErr(rootCmd.Execute())
}
