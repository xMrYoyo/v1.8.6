package fee

import (
	"github.com/ElrondNetwork/elrond-go-core/core/check"
	"github.com/ElrondNetwork/elrond-go/config"
	"github.com/ElrondNetwork/elrond-go/process"
	"github.com/ElrondNetwork/elrond-go/process/economics"
)

// ArgsNewFeeComputer holds the arguments for constructing a feeComputer
type ArgsNewFeeComputer struct {
	BuiltInFunctionsCostHandler economics.BuiltInFunctionsCostHandler
	EconomicsConfig             config.EconomicsConfig
	EnableEpochsConfig          config.EnableEpochs
	TxVersionChecker            process.TxVersionCheckerHandler
}

func (args *ArgsNewFeeComputer) check() error {
	if check.IfNil(args.BuiltInFunctionsCostHandler) {
		return process.ErrNilBuiltInFunctionsCostHandler
	}
	if check.IfNil(args.TxVersionChecker) {
		return process.ErrNilTransactionVersionChecker
	}

	return nil
}
