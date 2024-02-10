package tracers

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/core/vm"
	"github.com/ethereum/go-ethereum/params"
)

// BlockchainLogger is used to collect traces during chain processing. It's a similar
// interface as the go-ethereum's `core.BlockchainLogger` but adapted to Sei particularities.
//
// The method all starts with OnSei... to avoid confusion with the go-ethereum's `core.BlockchainLogger`
// interface and allow one to implement both interfaces in the same struct.
type BlockchainLogger interface {
	vm.EVMLogger
	state.StateLogger
	OnSeiBlockchainInit(chainConfig *params.ChainConfig)
	// OnSeiBlockStart is called before executing `block`.
	// `td` is the total difficulty prior to `block`.
	// `skip` indicates processing of this previously known block
	// will be skipped. OnBlockStart and OnBlockEnd will be emitted to
	// convey how chain is progressing. E.g. known blocks will be skipped
	// when node is started after a crash.
	OnSeiBlockStart(hash []byte, size uint64, b *types.Header)
	OnSeiBlockEnd(err error)

	// FIXME: What about OnSeiGenesisBlock/State, should we put something right here? It seems
	// it could be the best appealing ways to get our hands on the snapshot of the state
	// at the EVM "genesis" block (maybe the name should be different since it's not the genesis
	// block really, more the genesis state of the EVM).
}
type CtxBlockchainLoggerKeyType string

const CtxBlockchainLoggerKey = CtxBlockchainLoggerKeyType("evm_and_state_logger")

func SetCtxBlockchainLogger(ctx sdk.Context, logger BlockchainLogger) sdk.Context {
	return ctx.WithContext(context.WithValue(ctx.Context(), CtxBlockchainLoggerKey, logger))
}

func GetCtxBlockchainLogger(ctx sdk.Context) BlockchainLogger {
	rawVal := ctx.Context().Value(CtxBlockchainLoggerKey)
	if rawVal == nil {
		return nil
	}
	logger, ok := rawVal.(BlockchainLogger)
	if !ok {
		return nil
	}
	return logger
}
