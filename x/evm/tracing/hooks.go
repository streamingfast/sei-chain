package tracing

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/tracing"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/params"
)

// BlockEvent is emitted upon tracing an incoming block.
// It contains the block as well as consensus related information.
type BlockEvent struct {
	Hash   common.Hash
	Header *types.Header
	Size   uint64
}

type (
	// OnSeiBlockchainInitHook is called when the blockchain is initialized
	// once per process and receives the chain configuration.
	OnSeiBlockchainInitHook = func(chainConfig *params.ChainConfig)
	// OnSeiBlockStart is called before executing `block`.
	// `td` is the total difficulty prior to `block`.
	// `skip` indicates processing of this previously known block
	// will be skipped. OnBlockStart and OnBlockEnd will be emitted to
	// convey how chain is progressing. E.g. known blocks will be skipped
	// when node is started after a crash.
	OnSeiBlockStartHook = func(hash []byte, size uint64, b *types.Header)
	// OnSeiBlockEnd is called after executing `block` and receives the error
	// that occurred during processing. If an `err` is received in the callback,
	// it means the block should be discarded (optimistic execution failed for example).
	OnSeiBlockEnd = func(err error)
)

// Hooks is used to collect traces during chain processing. It's a similar
// interface as the go-ethereum's `tracing.Hooks` but adapted to Sei particularities.
//
// The method all starts with OnSei... to avoid confusion with the go-ethereum's `core.BlockchainLogger`
// interface and allow one to implement both interfaces in the same struct.
type Hooks struct {
	*tracing.Hooks

	OnSeiBlockchainInit OnSeiBlockchainInitHook
	OnSeiBlockStart     OnSeiBlockStartHook
	OnSeiBlockEnd       OnSeiBlockEnd
}
