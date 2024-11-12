package runtime

import (
	"github.com/spf13/viper"

	"github.com/wagoodman/dive/dive"
)

type Options struct {
	Ci           bool
	Image        string
	Source       dive.ImageSource
	IgnoreErrors bool
	ExportFile   string
	ExportTree   bool
	CiConfig     *viper.Viper
	BuildArgs    []string
}
