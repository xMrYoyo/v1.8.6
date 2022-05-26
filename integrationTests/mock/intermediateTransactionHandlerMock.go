package mock

import (
	"github.com/ElrondNetwork/elrond-go-core/data"
	"github.com/ElrondNetwork/elrond-go-core/data/block"
)

// IntermediateTransactionHandlerMock -
type IntermediateTransactionHandlerMock struct {
	AddIntermediateTransactionsCalled        func(txs []data.TransactionHandler) error
	GetNumOfCrossInterMbsAndTxsCalled        func() (int, int)
	CreateAllInterMiniBlocksCalled           func() []*block.MiniBlock
	VerifyInterMiniBlocksCalled              func(body *block.Body) error
	SaveCurrentIntermediateTxToStorageCalled func()
	CreateBlockStartedCalled                 func()
	CreateMarshalizedDataCalled              func(txHashes [][]byte) ([][]byte, error)
	GetAllCurrentFinishedTxsCalled           func() map[string]data.TransactionHandler
	RemoveProcessedResultsCalled             func(key []byte) [][]byte
	InitProcessedResultsCalled               func(key []byte)
	intermediateTransactions                 []data.TransactionHandler
}

// RemoveProcessedResults -
func (ith *IntermediateTransactionHandlerMock) RemoveProcessedResults(key []byte) [][]byte {
	if ith.RemoveProcessedResultsCalled != nil {
		return ith.RemoveProcessedResultsCalled(key)
	}
	return nil
}

// InitProcessedResults -
func (ith *IntermediateTransactionHandlerMock) InitProcessedResults(key []byte) {
	if ith.InitProcessedResultsCalled != nil {
		ith.InitProcessedResultsCalled(key)
	}
}

// GetCreatedInShardMiniBlock -
func (ith *IntermediateTransactionHandlerMock) GetCreatedInShardMiniBlock() *block.MiniBlock {
	return &block.MiniBlock{}
}

// GetAllCurrentFinishedTxs -
func (ith *IntermediateTransactionHandlerMock) GetAllCurrentFinishedTxs() map[string]data.TransactionHandler {
	if ith.GetAllCurrentFinishedTxsCalled != nil {
		return ith.GetAllCurrentFinishedTxsCalled()
	}
	return nil
}

// CreateMarshalizedData -
func (ith *IntermediateTransactionHandlerMock) CreateMarshalizedData(txHashes [][]byte) ([][]byte, error) {
	if ith.CreateMarshalizedDataCalled == nil {
		return nil, nil
	}
	return ith.CreateMarshalizedDataCalled(txHashes)
}

// AddIntermediateTransactions -
func (ith *IntermediateTransactionHandlerMock) AddIntermediateTransactions(txs []data.TransactionHandler) error {
	if ith.AddIntermediateTransactionsCalled == nil {
		ith.intermediateTransactions = append(ith.intermediateTransactions, txs...)
		return nil
	}
	return ith.AddIntermediateTransactionsCalled(txs)
}

// GetIntermediateTransactions -
func (ith *IntermediateTransactionHandlerMock) GetIntermediateTransactions() []data.TransactionHandler {
	return ith.intermediateTransactions
}

// Clean -
func (ith *IntermediateTransactionHandlerMock) Clean() {
	ith.intermediateTransactions = make([]data.TransactionHandler, 0)
}

// GetNumOfCrossInterMbsAndTxs -
func (ith *IntermediateTransactionHandlerMock) GetNumOfCrossInterMbsAndTxs() (int, int) {
	if ith.GetNumOfCrossInterMbsAndTxsCalled == nil {
		return 0, 0
	}
	return ith.GetNumOfCrossInterMbsAndTxsCalled()
}

// CreateAllInterMiniBlocks -
func (ith *IntermediateTransactionHandlerMock) CreateAllInterMiniBlocks() []*block.MiniBlock {
	if ith.CreateAllInterMiniBlocksCalled == nil {
		return nil
	}
	return ith.CreateAllInterMiniBlocksCalled()
}

// VerifyInterMiniBlocks -
func (ith *IntermediateTransactionHandlerMock) VerifyInterMiniBlocks(body *block.Body) error {
	if ith.VerifyInterMiniBlocksCalled == nil {
		return nil
	}
	return ith.VerifyInterMiniBlocksCalled(body)
}

// SaveCurrentIntermediateTxToStorage -
func (ith *IntermediateTransactionHandlerMock) SaveCurrentIntermediateTxToStorage() {
	if ith.SaveCurrentIntermediateTxToStorageCalled == nil {
		return
	}
	ith.SaveCurrentIntermediateTxToStorageCalled()
}

// CreateBlockStarted -
func (ith *IntermediateTransactionHandlerMock) CreateBlockStarted() {
	if len(ith.intermediateTransactions) > 0 {
		ith.intermediateTransactions = make([]data.TransactionHandler, 0)
	}
	if ith.CreateBlockStartedCalled != nil {
		ith.CreateBlockStartedCalled()
	}
}

// IsInterfaceNil returns true if there is no value under the interface
func (ith *IntermediateTransactionHandlerMock) IsInterfaceNil() bool {
	return ith == nil
}
