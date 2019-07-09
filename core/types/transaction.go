package types

import (
	"encoding/json"

	"github.com/fletaio/fleta/common"
	"github.com/fletaio/fleta/common/amount"
)

// Transaction defines common transaction functions
type Transaction interface {
	json.Marshaler
	Timestamp() uint64
	Fee(loader LoaderWrapper) *amount.Amount
	Validate(p Process, loader LoaderWrapper, signers []common.PublicHash) error
	Execute(p Process, ctx *ContextWrapper, index uint16) error
}
