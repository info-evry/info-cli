package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type Contact struct {
	FullName string
	Emails   []Emails
	Urls     []Urls
	Office   Office
}

type Emails struct {
	FullEmail string
}

type Urls struct {
	Name string
	Href string
}

type Office struct {
	Number int
	Phone  string
}

func init() {
	rootCmd.AddCommand(cmdContact)
}

var cmdContact = &cobra.Command{
	Use:   "contact [name]",
	Short: "Search for a teacher's contact information",
	Long:  "Search for a teacher's contact information.",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		response, err := http.Get("https://api.info-evry.fr/v1/contacts/" + strings.Join(args, " "))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			data, _ := ioutil.ReadAll(response.Body)

			var contacts []Contact
			json.Unmarshal([]byte(data), &contacts)

			fmt.Printf("I found %d result(s):\n", len(contacts))

			for i := 0; i < len(contacts); i++ {
				fmt.Printf("\n\n")

				fmt.Printf("Name : %s\n", contacts[i].FullName)

				fmt.Println("\nMail(s):")
				for _, value := range contacts[i].Emails {
					fmt.Printf("• %s\n", value.FullEmail)
				}

				fmt.Println("\nLink(s):")
				for _, value := range contacts[i].Urls {
					fmt.Printf("• %s - %s\n", value.Name, value.Href)
				}

				fmt.Printf("\nOffice : %d\n", contacts[i].Office.Number)
				fmt.Printf("Phone : %s\n", contacts[i].Office.Phone)
			}
		}
	},
}
