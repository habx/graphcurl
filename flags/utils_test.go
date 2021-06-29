package flags

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func Test_BindFlags(t *testing.T) {
	Convey("testing bind variables", t, func() {
		commands := &cobra.Command{
			Use:     "graphcurl",
		}
		var userAgent string
		commands.Flags().StringVarP(&userAgent, "test-graphcurl", "", "", " ")
		commands.Flags().StringVarP(&userAgent, "user", "", "a", " ")
		commands.Flags().StringVarP(&userAgent, "-", "", "", " ")
		os.Setenv("TEST_GRAPHCURL", "xxxxx")
		BindFlags(commands, viper.New())
	})
}