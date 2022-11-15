package flags

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/habx/graphcurl/logger"
)

func BindFlags(cmd *cobra.Command, v *viper.Viper) {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		if strings.Contains(f.Name, "-") {
			err := v.BindEnv(f.Name, strings.ToUpper(strings.ReplaceAll(f.Name, "-", "_")))
			if err != nil {
				logger.GetLogger("info").Sugar().Errorw("cannot bind env variable", "err", err)
			}
		}
		if !f.Changed && v.IsSet(f.Name) {
			val := v.Get(f.Name)
			err := cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
			if err != nil {
				logger.GetLogger("info").Sugar().Errorw("cannot set value from env variable", "err", err)
			}
		}
	})
}
