package cmd_test

import (
	"testing"

	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	require "github.com/stretchr/testify/require"

	cmd "github.com/pumpkinzomb/cosmos-sandbox/cmd/gaiad/cmd"
	app "github.com/pumpkinzomb/cosmos-sandbox/gaia/app"
)

func TestRootCmdConfig(t *testing.T) {
	rootCmd, _ := cmd.NewRootCmd()
	rootCmd.SetArgs([]string{
		"config",          // Test the config cmd
		"keyring-backend", // key
		"test",            // value
	})

	require.NoError(t, svrcmd.Execute(rootCmd, app.DefaultNodeHome))
}
