package cmd

import (
	"errors"
	"fmt"
	"net/url"

	eventtests "github.com/yale-mgt-656/eventbrite-clone-selenium-tests/tests"

	"github.com/spf13/cobra"
)

func isValidURL(u string) bool {
	_, err := url.ParseRequestURI(u)
	if err == nil {
		return true
	}
	return false
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test SELENIUM_URL TEST_URL",
	Short: "Runs the automated test suite against your application",
	Long: `Runs the automated test suite against your application.
You should be running your app and selenium/chromedriver.
You'll need to know the URL for both of these and provide
them as arguments to this program.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) != 2 {
			return errors.New("Requires exactly two arguments: SELENIUM_URL and TEST_URL")
		}
		for _, arg := range args {
			if isValidURL(arg) == false {
				return fmt.Errorf("invalid URL: %s", arg)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		eventtests.RunForURL(args[0], args[1], false, 0)
	},
}

func init() {
	RootCmd.AddCommand(testCmd)
}
