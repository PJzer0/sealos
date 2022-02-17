// Copyright © 2021 sealos.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/fanux/sealos/cmd/sealos/boot"
	"github.com/fanux/sealos/pkg/types/contants"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "sealos",
	Short: "simplest way install kubernetes tools.",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(boot.OnBootOnDie)
	rootCmd.PersistentFlags().StringVar(&boot.CmdFlag.Root.ConfigDir, "config-dir", contants.DefaultConfigPath, "config dir (default is $HOME/.sealos)")
	rootCmd.PersistentFlags().StringVar(&boot.CmdFlag.Root.ConfigFile, "config-file", contants.DefaultConfigFile, "config file (default is /config.yaml )")

	rootCmd.PersistentFlags().BoolVar(&boot.CmdFlag.Root.Debug, "debug", false, "enable debug logger")
	rootCmd.PersistentFlags().BoolVar(&boot.CmdFlag.Root.ShowPatch, "path", false, "enable show code path")
}