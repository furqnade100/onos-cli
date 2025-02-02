// Copyright 2019-present Open Networking Foundation.
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

package e2t

import (
	"github.com/onosproject/onos-lib-go/pkg/cli"
	loglib "github.com/onosproject/onos-lib-go/pkg/logging/cli"
	"github.com/spf13/cobra"
)

const (
	configName     = "e2t"
	defaultAddress = "onos-e2t:5150"
)

// init initializes the command line
func init() {
	cli.InitConfig(configName)
}

// Init is a hook called after cobra initialization
func Init() {
	// noop for now
}

// GetCommand returns the root command for the RAN service
func GetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "e2t {get,add,remove,watch} [args]",
		Short: "ONOS e2t subsystem commands",
	}

	cli.AddConfigFlags(cmd, defaultAddress)
	cmd.AddCommand(cli.GetConfigCommand())
	cmd.AddCommand(loglib.GetCommand())
	cmd.AddCommand(getGetCommand())
	cmd.AddCommand(getWatchCommand())
	return cmd
}

func getGetCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "get {connections,subscriptions,subscription} [args]",
		Short:   "Get command",
		Aliases: []string{"list"},
	}

	cmd.AddCommand(getGetChannelsCommand())
	cmd.AddCommand(getGetChannelCommand())
	cmd.AddCommand(getGetSubscriptionsCommand())
	cmd.AddCommand(getGetSubscriptionCommand())
	return cmd
}

func getWatchCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "watch {subscriptions} [args]",
		Short: "Watch command",
	}

	cmd.AddCommand(getWatchSubscriptionsCommand())
	cmd.AddCommand(getWatchChannelsCommand())
	return cmd
}
