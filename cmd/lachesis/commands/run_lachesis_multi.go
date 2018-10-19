// +build multi

// This version will be built when tag MULTI is used in go build
//
package commands

import (
	"github.com/andrecronje/lachesis/src/lachesis"
	"github.com/jinzhu/copier"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func runLachesis(cmd *cobra.Command, args []string) error {

	config := NewDefaultCLIConfig()

	if err := bindFlagsLoadViper(cmd, config); err != nil {
		return err
	}

	err := viper.Unmarshal(config)
	if err != nil {
		return err
	}

	config2 := &CLIConfig{
		Lachesis:   lachesis.LachesisConfig{},
		ProxyAddr:  "127.0.0.1:1338",
		ClientAddr: "127.0.0.1:1339",
		Inapp:      false,
	}
	copier.Copy(&config2.Lachesis, &config.Lachesis)
	config2.Lachesis.BindAddr = "127.0.0.2:1200"
	config2.ProxyAddr = "127.0.0.2:9000"

	go runSingleLachesis(config2)


	config2 := &CLIConfig{
		Lachesis:   lachesis.LachesisConfig{},
		ProxyAddr:  "127.0.0.1:1338",
		ClientAddr: "127.0.0.1:1339",
		Inapp:      false,
	}
	copier.Copy(&config2.Lachesis, &config.Lachesis)
	config2.Lachesis.BindAddr = "127.0.0.2:1200"
	config2.ProxyAddr = "127.0.0.2:9000"

	go runSingleLachesis(config2)



	config3 := &CLIConfig{
		Lachesis:   lachesis.LachesisConfig{},
		ProxyAddr:  "127.0.0.1:1338",
		ClientAddr: "127.0.0.1:1339",
		Inapp:      false,
	}
	copier.Copy(&config3.Lachesis, &config.Lachesis)
	config3.Lachesis.BindAddr = "127.0.0.3:1200"
	config3.ProxyAddr = "127.0.0.3:9000"

	go runSingleLachesis(config3)


	

	return runSingleLachesis(config)
}
