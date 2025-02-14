package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/go-logr/logr"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/gkwa/anyhobbit/core"
	"github.com/gkwa/anyhobbit/internal/logger"
)

var (
	cfgFile   string
	outFile   string
	verbose   int
	logFormat string
	quiet     bool
	cliLogger logr.Logger
)

var rootCmd = &cobra.Command{
	Use:   "anyhobbit",
	Short: "A tool for generating Renovate configs with various presets",
	Long:  `A tool for generating Renovate configs with various presets for different update strategies.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cliLogger.IsZero() {
			cliLogger = logger.NewConsoleLogger(verbose, logFormat == "json")
		}

		ctx := logr.NewContext(context.Background(), cliLogger)
		cmd.SetContext(ctx)
	},
}

var renovateCmd = &cobra.Command{
	Use:   "renovate",
	Short: "Renovate subcommands",
	Long:  "Renovate subcommands for generating different Renovate configuration presets",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return fmt.Errorf("unknown command %q for %q", args[0], cmd.CommandPath())
	},
}

func createCommand(name, short, long string) {
	cmd := &cobra.Command{
		Use:   name,
		Short: short,
		Long:  long,
		RunE: func(cmd *cobra.Command, args []string) error {
			return core.GenerateConfig(name)
		},
	}
	renovateCmd.AddCommand(cmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.anyhobbit.yaml)")
	rootCmd.PersistentFlags().StringVarP(&outFile, "outfile", "o", ".renovaterc.json", "output file path")
	rootCmd.PersistentFlags().CountVarP(&verbose, "verbose", "v", "increase verbosity")
	rootCmd.PersistentFlags().StringVar(&logFormat, "log-format", "", "json or text (default is text)")
	rootCmd.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "suppress output")

	if err := viper.BindPFlag("outfile", rootCmd.PersistentFlags().Lookup("outfile")); err != nil {
		fmt.Printf("Error binding outfile flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		fmt.Printf("Error binding verbose flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("log-format", rootCmd.PersistentFlags().Lookup("log-format")); err != nil {
		fmt.Printf("Error binding log-format flag: %v\n", err)
		os.Exit(1)
	}
	if err := viper.BindPFlag("quiet", rootCmd.PersistentFlags().Lookup("quiet")); err != nil {
		fmt.Printf("Error binding quiet flag: %v\n", err)
		os.Exit(1)
	}

	rootCmd.AddCommand(renovateCmd)

	for name, cmd := range commands {
		createCommand(name, cmd.short, cmd.long)
	}
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".anyhobbit")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	outFile = viper.GetString("outfile")
	logFormat = viper.GetString("log-format")
	verbose = viper.GetInt("verbose")
}

func LoggerFrom(ctx context.Context, keysAndValues ...interface{}) logr.Logger {
	if cliLogger.IsZero() {
		cliLogger = logger.NewConsoleLogger(verbose, logFormat == "json")
	}
	newLogger := cliLogger
	if ctx != nil {
		if l, err := logr.FromContext(ctx); err == nil {
			newLogger = l
		}
	}
	return newLogger.WithValues(keysAndValues...)
}
