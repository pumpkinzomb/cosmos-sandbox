package client

import (
	govclient "github.com/cosmos/cosmos-sdk/x/gov/client"

	"github.com/pumpkinzomb/cosmos-sandbox/x/evmos/incentives/client/cli"
	"github.com/pumpkinzomb/cosmos-sandbox/x/evmos/incentives/client/rest"
)

var (
	RegisterIncentiveProposalHandler = govclient.NewProposalHandler(cli.NewRegisterIncentiveProposalCmd, rest.RegisterIncentiveProposalRESTHandler)
	CancelIncentiveProposalHandler   = govclient.NewProposalHandler(cli.NewCancelIncentiveProposalCmd, rest.CancelIncentiveProposalRequestRESTHandler)
)
