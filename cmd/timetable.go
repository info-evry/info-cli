package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type Timetable struct {
	Url string
}

func init() {
	rootCmd.AddCommand(cmdTimetable)
}

var cmdTimetable = &cobra.Command{
	Use:   "timetable [level] [group]",
	Short: "Get a link to timetable image",
	Long:  "Get a link to timetable image.",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		date := time.Now()

		response, err := http.Get("https://api.info-evry.fr/v1/edt/" + args[0] + "/" + args[1] + "/" + fmt.Sprint(date.Day()) + "/" + fmt.Sprintf("%d", date.Month()) + "/" + fmt.Sprint(date.Year()))
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			data, _ := ioutil.ReadAll(response.Body)

			var timetable Timetable
			json.Unmarshal([]byte(data), &timetable)

			fmt.Println(timetable.Url)
		}
	},
}
