package trie

import (
	"github.com/multiversx/mx-chain-go/common"
)

// MockStatistics -
type MockStatistics struct {
	WaitForSnapshotsToFinishCalled func()
}

// SnapshotFinished -
func (m *MockStatistics) SnapshotFinished() {
}

// NewSnapshotStarted -
func (m *MockStatistics) NewSnapshotStarted() {
}

// WaitForSnapshotsToFinish -
func (m *MockStatistics) WaitForSnapshotsToFinish() {
	if m.WaitForSnapshotsToFinishCalled != nil {
		m.WaitForSnapshotsToFinishCalled()
	}
}

// AddTrieStats -
func (m *MockStatistics) AddTrieStats(_ common.TrieStatisticsHandler) {
}
