package keeper

import (
	"github.com/louisliu2048/argus/x/bpmn/types"
)

var _ types.QueryServer = Keeper{}
