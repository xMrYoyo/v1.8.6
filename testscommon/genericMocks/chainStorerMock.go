package genericMocks

import (
	"github.com/ElrondNetwork/elrond-go/dataRetriever"
	"github.com/ElrondNetwork/elrond-go/storage"
)

// ChainStorerMock -
type ChainStorerMock struct {
	Miniblocks    *StorerMock
	Transactions  *StorerMock
	Rewards       *StorerMock
	Unsigned      *StorerMock
	Logs          *StorerMock
	MetaHdrNonce  *StorerMock
	ShardHdrNonce *StorerMock
	Others        *StorerMock
}

// NewChainStorerMock -
func NewChainStorerMock(epoch uint32) *ChainStorerMock {
	return &ChainStorerMock{
		Miniblocks:    NewStorerMockWithEpoch(epoch),
		Transactions:  NewStorerMockWithEpoch(epoch),
		Rewards:       NewStorerMockWithEpoch(epoch),
		Unsigned:      NewStorerMockWithEpoch(epoch),
		Logs:          NewStorerMockWithEpoch(epoch),
		MetaHdrNonce:  NewStorerMockWithEpoch(epoch),
		ShardHdrNonce: NewStorerMockWithEpoch(epoch),
		Others:        NewStorerMockWithEpoch(epoch),
	}
}

// CloseAll -
func (sm *ChainStorerMock) CloseAll() error {
	return nil
}

// AddStorer -
func (sm *ChainStorerMock) AddStorer(_ dataRetriever.UnitType, _ storage.Storer) {
	panic("not supported")
}

// GetStorer -
func (sm *ChainStorerMock) GetStorer(unitType dataRetriever.UnitType) storage.Storer {
	if unitType == dataRetriever.MiniBlockUnit {
		return sm.Miniblocks
	}
	if unitType == dataRetriever.TransactionUnit {
		return sm.Transactions
	}
	if unitType == dataRetriever.RewardTransactionUnit {
		return sm.Rewards
	}
	if unitType == dataRetriever.UnsignedTransactionUnit {
		return sm.Unsigned
	}
	if unitType == dataRetriever.TxLogsUnit {
		return sm.Logs
	}
	if unitType == dataRetriever.MetaHdrNonceHashDataUnit {
		return sm.MetaHdrNonce
	}
	if unitType == dataRetriever.ShardHdrNonceHashDataUnit {
		return sm.ShardHdrNonce
	}

	// According to: dataRetriever/interface.go
	if unitType > dataRetriever.ShardHdrNonceHashDataUnit {
		return sm.MetaHdrNonce
	}

	return sm.Others
}

// Has -
func (sm *ChainStorerMock) Has(_ dataRetriever.UnitType, _ []byte) error {
	return nil
}

// Get -
func (sm *ChainStorerMock) Get(unitType dataRetriever.UnitType, key []byte) ([]byte, error) {
	return sm.GetStorer(unitType).Get(key)
}

// Put -
func (sm *ChainStorerMock) Put(unitType dataRetriever.UnitType, key []byte, value []byte) error {
	return sm.GetStorer(unitType).Put(key, value)
}

// GetAll -
func (sm *ChainStorerMock) GetAll(unitType dataRetriever.UnitType, keys [][]byte) (map[string][]byte, error) {
	storer := sm.GetStorer(unitType)

	data := make(map[string][]byte)
	for _, key := range keys {
		buff, err := storer.Get(key)
		if err != nil {
			return nil, err
		}

		data[string(key)] = buff
	}

	return data, nil
}

// SetEpochForPutOperation -
func (sm *ChainStorerMock) SetEpochForPutOperation(_ uint32) {
	panic("not supported")
}

// GetAllStorers -
func (sm *ChainStorerMock) GetAllStorers() map[dataRetriever.UnitType]storage.Storer {
	return map[dataRetriever.UnitType]storage.Storer{
		dataRetriever.MiniBlockUnit:             sm.Miniblocks,
		dataRetriever.TransactionUnit:           sm.Transactions,
		dataRetriever.RewardTransactionUnit:     sm.Rewards,
		dataRetriever.UnsignedTransactionUnit:   sm.Unsigned,
		dataRetriever.TxLogsUnit:                sm.Logs,
		dataRetriever.MetaHdrNonceHashDataUnit:  sm.MetaHdrNonce,
		dataRetriever.ShardHdrNonceHashDataUnit: sm.ShardHdrNonce,
	}
}

// Destroy -
func (sm *ChainStorerMock) Destroy() error {
	return nil
}

// IsInterfaceNil -
func (sm *ChainStorerMock) IsInterfaceNil() bool {
	return sm == nil
}
