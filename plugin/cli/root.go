package cli

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	// external
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	// "github.com/sniperkit/snk.golang.cobra"
	// "github.com/sniperkit/snk.golang.viper"

	// internal - core
	conf "github.com/sniperkit/snk.golang.dirwalk/pkg/config"
	out "github.com/sniperkit/snk.golang.dirwalk/pkg/output"
)

// RootCmd is the root command for limo
var RootCmd = &cobra.Command{
	Use:   conf.ProgramName,
	Short: "A CLI for formating import packages.",
	Long:  fmt.Sprintf(`%s allows you to formating import packages and group them by section in your code.`, conf.ProgramName),
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPreRun with args: %v\n", args)
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PreRun with args: %v\n", args)
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd Run with args: %v\n", args)
	},
	PostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PostRun with args: %v\n", args)
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Inside rootCmd PersistentPostRun with args: %v\n", args)
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func setRuntimeLimit(max int) {
	switch {
	case numCPUs > max:
		runtime.GOMAXPROCS(max)
	case max == -1:
		return
	default:
		runtime.GOMAXPROCS(numCPUs)
	}

}

func init() {

	numCPUs := runtime.NumCPU()

	flags := RootCmd.PersistentFlags()

	flags.StringVarP(&options.output, "output", "o", "color", "output type")
	flags.StringVarP(&options.dirConf, "conf-dir", "c", ".goimpi.yaml", "write config to prefix dir, default XGDB Base directory.")

	// config file
	configInfo := fmt.Sprintf("config file (default is $HOME/.%s.yaml)", strings.ToLower(conf.ProgramName))
	flags.PersistentFlags().StringVar(&cfgFile, "conf", "", configInfo)
	flags.BoolVarP(&options.writeConf, "write-conf", "w", true, "write config file to xgd dir")

	flags.IntVarP(&options.numCPUs, "max-procs", "m", numCPUs, "max parallel processes. (default: number of cores of the local machine).")

	flags.BoolVarP(&options.dryMode, "dry-mode", "n", false, "don't make any changes; perform checks only")
	flags.BoolVarP(&options.version, "version", "v", false, "display version")
	flags.BoolVarP(&options.debug, "debug", "d", false, "debug mode")

	viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	viper.SetDefault("license", "MIT")

}

func getConfiguration() (*conf.Config, error) {
	if configuration == nil {
		var err error
		if configuration, err = conf.ReadConfig(); err != nil {
			return nil, err
		}
	}
	return configuration, nil
}

func getOutput() out.Output {
	o := out.ForName(options.output)
	oc, err := getConfiguration()
	if err == nil {
		o.Configure(oc.GetOutput(options.output))
	}
	return o
}

func fatalOnError(err error) {
	if err != nil {
		if options.debug {
			log.WithFields(logFields{
				"prog-name": conf.ProgramName,
			}).Fatalln("error", err.Error())
		} else {
			getOutput().Fatal(err.Error())
		}
	}
}
