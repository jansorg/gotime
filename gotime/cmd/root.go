package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/text/message"

	"github.com/jansorg/gotime/gotime/config"
	"github.com/jansorg/gotime/gotime/context"
	"github.com/jansorg/gotime/gotime/i18n"
	"github.com/jansorg/gotime/gotime/query"
	"github.com/jansorg/gotime/gotime/store"
)

var ctx context.GoTimeContext
var configFile string

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().String("data-dir", "", "data directory (default is $HOME/.gotime)")
	RootCmd.PersistentFlags().StringVarP(&configFile, "config", "c", "", "config file (default is $HOME/.gotime.yaml)")

	RootCmd.PersistentFlags().String("cpu-profile", "", "create a cpu profile for performance measurement")
	RootCmd.Flag("cpu-profile").Hidden = true
	RootCmd.Flags().String("mem-profile", "", "create a memory profile for performance measurement")
	RootCmd.Flag("mem-profile").Hidden = true

	newProjectsCommand(&ctx, RootCmd)
	newTagsCommand(&ctx, RootCmd)
	newFramesCommand(&ctx, RootCmd)
	newCreateCommand(&ctx, RootCmd)
	newRemoveCommand(&ctx, RootCmd)
	newStartCommand(&ctx, RootCmd)
	newStopCommand(&ctx, RootCmd)
	newCancelCommand(&ctx, RootCmd)
	newReportCommand(&ctx, RootCmd)
	newImportCommand(&ctx, RootCmd)
	newStatusCommand(&ctx, RootCmd)
	newInvoiceCommand(&ctx, RootCmd)
	newConfigCommand(&ctx, RootCmd)
	// hidden command
	newCompletionCommand(&ctx, RootCmd)

	if err := viper.BindPFlag(config.KeyDataDir, RootCmd.PersistentFlags().Lookup("data-dir")); err != nil {
		fatal(err)
	}
}

func fatal(err ...interface{}) {
	fmt.Println(append([]interface{}{"Error: "}, err...)...)
	os.Exit(1)
}

func initConfig() {
	config.SetDefaults()
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	// setup config dir if it doesn't exist
	dataDir := viper.GetString(config.KeyDataDir)
	if _, err := os.Stat(dataDir); os.IsNotExist(err) {
		if err := os.MkdirAll(dataDir, 0700); err != nil {
			fatal(err)
		}
	}

	if err := viper.ReadInConfig(); !os.IsNotExist(err) {
		// fatal(err)
	}

	dataStore, err := store.NewStore(dataDir)
	if err != nil {
		fatal(err)
	}

	ctx.Store = dataStore
	ctx.StoreHelper = store.NewStoreHelper(dataStore)
	ctx.Query = query.NewStoreQuery(dataStore)
	ctx.Language = i18n.FindPreferredLanguages()
	ctx.LocalePrinter = message.NewPrinter(ctx.Language)
	ctx.Locale = i18n.FindLocale(ctx.Language)
	ctx.DurationPrinter = i18n.NewDurationPrinter(ctx.Language)
	ctx.DateTimePrinter = i18n.NewDateTimePrinter(ctx.Language)
}

const (
	//language=BASH
	bash_completion_func = `__gotime_projects_get()
{
		  local IFS=$'\n'
#          if [[ $COMP_CWORD -eq 2 ]]; then
            local projects="$(gotime projects -f name | sed -e 's/ /\\\\ /g')"
            COMPREPLY=($(compgen -W "$projects" -- ${cur}))
#          else
#            tags="$(tags | sed -e 's/ /\\\\ /g' | awk '$0="+"$0')"
#            COMPREPLY=($(compgen -W "$tags" -- ${cur}))
#          fi
}

__gotime_get_projects()
{
    if [[ ${#nouns[@]} -eq 0 ]]; then
        __gotime_projects_get ""
	else
	    __gotime_projects_get ${nouns[${#nouns[@]} -1]}
    fi
    if [[ $? -eq 0 ]]; then
        return 0
    fi
}

__custom_func() {
    case ${last_command} in
        gotime_start)
            __gotime_get_projects
            return
            ;;
        *)
            ;;
    esac
}
`
)

var RootCmd = &cobra.Command{
	Use:                    "gotime",
	Short:                  "gotime is a command line application to track time.",
	Version:                "1.0.0",
	BashCompletionFunction: bash_completion_func,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		cpuProfile, _ := cmd.Flags().GetString("cpu-profile")
		if cpuProfile != "" {
			log.Println("creating a cpu profile...")
			f, err := os.Create(cpuProfile)
			if err != nil {
				log.Fatal("could not create CPU profile: ", err)
			}
			if err := pprof.StartCPUProfile(f); err != nil {
				log.Fatal("could not start CPU profile: ", err)
			}
		}
	},
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		cpuProfile, _ := cmd.Flags().GetString("cpu-profile")
		if cpuProfile != "" {
			pprof.StopCPUProfile()
		}

		memProfile, _ := cmd.Flags().GetString("mem-profile")
		if memProfile != "" {
			log.Println("creating a mem profile...")
			f, err := os.Create(memProfile)
			if err != nil {
				log.Fatal("could not create memory profile: ", err)
			}
			runtime.GC() // get up-to-date statistics
			if err := pprof.WriteHeapProfile(f); err != nil {
				log.Fatal("could not write memory profile: ", err)
			}
			f.Close()
		}
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
