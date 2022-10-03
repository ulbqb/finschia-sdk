package tendermint

import (
	"github.com/line/lbm-sdk/x/ibc/light-clients/07-tendermint/types"
)

// Name returns the IBC client name
func Name() string {
	return types.SubModuleName
}
