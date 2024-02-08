package staking

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/multiversx/mx-chain-core-go/core"
	"github.com/multiversx/mx-chain-core-go/data/transaction"
	"github.com/multiversx/mx-chain-go/config"
	"github.com/multiversx/mx-chain-go/node/chainSimulator"
	"github.com/multiversx/mx-chain-go/node/chainSimulator/components/api"
	"github.com/multiversx/mx-chain-go/node/chainSimulator/configs"
	"github.com/multiversx/mx-chain-go/vm"
	"github.com/stretchr/testify/require"
)

// Test description
// All test cases will do a stake transaction and wait till the new node is jailed
// testcase1 -- unJail transaction will be sent when staking v3.5 is still action --> node status should be `new` after unjail
// testcase2 -- unJail transaction will be sent when staking v4 step1 is action --> node status should be `auction` after unjail
// testcase3 -- unJail transaction will be sent when staking v4 step2 is action --> node status should be `auction` after unjail
// testcase4 -- unJail transaction will be sent when staking v4 step3 is action --> node status should be `auction` after unjail
func TestChainSimulator_ValidatorJailUnJail(t *testing.T) {
	if testing.Short() {
		t.Skip("this is not a short test")
	}

	t.Run("staking ph 4 is not active", func(t *testing.T) {
		testChainSimulatorJailAndUnJail(t, 4, "new")
	})

	t.Run("staking ph 4 step 1 active", func(t *testing.T) {
		testChainSimulatorJailAndUnJail(t, 5, "auction")
	})

	t.Run("staking ph 4 step 2 active", func(t *testing.T) {
		testChainSimulatorJailAndUnJail(t, 6, "auction")
	})

	t.Run("staking ph 4 step 3 active", func(t *testing.T) {
		testChainSimulatorJailAndUnJail(t, 7, "auction")
	})
}

func testChainSimulatorJailAndUnJail(t *testing.T, targetEpoch int32, nodeStatusAfterUnJail string) {
	startTime := time.Now().Unix()
	roundDurationInMillis := uint64(6000)
	roundsPerEpoch := core.OptionalUint64{
		HasValue: true,
		Value:    20,
	}

	numOfShards := uint32(3)

	cs, err := chainSimulator.NewChainSimulator(chainSimulator.ArgsChainSimulator{
		BypassTxSignatureCheck:   false,
		TempDir:                  t.TempDir(),
		PathToInitialConfig:      defaultPathToInitialConfig,
		NumOfShards:              numOfShards,
		GenesisTimestamp:         startTime,
		RoundDurationInMillis:    roundDurationInMillis,
		RoundsPerEpoch:           roundsPerEpoch,
		ApiInterface:             api.NewNoApiInterface(),
		MinNodesPerShard:         3,
		MetaChainMinNodes:        3,
		NumNodesWaitingListMeta:  1,
		NumNodesWaitingListShard: 1,
		AlterConfigsFunction: func(cfg *config.Configs) {
			cfg.EpochConfig.EnableEpochs.StakingV4Step1EnableEpoch = 5
			cfg.EpochConfig.EnableEpochs.StakingV4Step2EnableEpoch = 6
			cfg.EpochConfig.EnableEpochs.StakingV4Step3EnableEpoch = 7

			cfg.EpochConfig.EnableEpochs.MaxNodesChangeEnableEpoch[2].EpochEnable = 7

			newNumNodes := cfg.SystemSCConfig.StakingSystemSCConfig.MaxNumberOfNodesForStake + 8 // 8 nodes until new nodes will be placed on queue
			configs.SetMaxNumberOfNodesInConfigs(cfg, newNumNodes, numOfShards)

			cfg.RatingsConfig.ShardChain.RatingSteps.ConsecutiveMissedBlocksPenalty = 100
			cfg.RatingsConfig.ShardChain.RatingSteps.HoursToMaxRatingFromStartRating = 1
			cfg.RatingsConfig.MetaChain.RatingSteps.ConsecutiveMissedBlocksPenalty = 100
			cfg.RatingsConfig.MetaChain.RatingSteps.HoursToMaxRatingFromStartRating = 1
		},
	})
	require.Nil(t, err)
	require.NotNil(t, cs)
	defer func() {
		_ = cs.Close()
	}()

	metachainNode := cs.GetNodeHandler(core.MetachainShardId)
	err = cs.GenerateBlocks(30)
	require.Nil(t, err)

	_, blsKeys, err := chainSimulator.GenerateBlsPrivateKeys(1)
	require.Nil(t, err)

	mintValue := big.NewInt(0).Mul(oneEGLD, big.NewInt(3000))
	walletKeyBech, err := cs.GenerateAndMintWalletAddress(core.AllShardId, mintValue)
	require.Nil(t, err)

	walletKey, err := metachainNode.GetCoreComponents().AddressPubKeyConverter().Decode(walletKeyBech)
	require.Nil(t, err)

	txDataField := fmt.Sprintf("stake@01@%s@%s", blsKeys[0], mockBLSSignature)
	txStake := generateTransaction(walletKey, 0, vm.ValidatorSCAddress, minimumStakeValue, txDataField, gasLimitForStakeOperation)
	stakeTx, err := cs.SendTxAndGenerateBlockTilTxIsExecuted(txStake, maxNumOfBlockToGenerateWhenExecutingTx)
	require.Nil(t, err)
	require.NotNil(t, stakeTx)

	// wait node to be jailed
	err = cs.GenerateBlocksUntilEpochIsReached(4)
	require.Nil(t, err)

	decodedBLSKey, _ := hex.DecodeString(blsKeys[0])
	status := getBLSKeyStatus(t, metachainNode, decodedBLSKey)
	require.Equal(t, "jailed", status)

	// do an unjail transaction
	unJailValue, _ := big.NewInt(0).SetString("2500000000000000000", 10)
	txUnJailDataField := fmt.Sprintf("unJail@%s", blsKeys[0])
	txUnJail := generateTransaction(walletKey, 1, vm.ValidatorSCAddress, unJailValue, txUnJailDataField, gasLimitForStakeOperation)

	err = cs.GenerateBlocksUntilEpochIsReached(targetEpoch)
	require.Nil(t, err)

	unJailTx, err := cs.SendTxAndGenerateBlockTilTxIsExecuted(txUnJail, maxNumOfBlockToGenerateWhenExecutingTx)
	require.Nil(t, err)
	require.NotNil(t, unJailTx)
	require.Equal(t, transaction.TxStatusSuccess, unJailTx.Status)

	err = cs.GenerateBlocks(1)
	require.Nil(t, err)

	status = getBLSKeyStatus(t, metachainNode, decodedBLSKey)
	require.Equal(t, "staked", status)

	err = metachainNode.GetProcessComponents().ValidatorsProvider().ForceUpdate()
	require.Nil(t, err)
	validatorsStatistics, err := metachainNode.GetFacadeHandler().ValidatorStatisticsApi()
	require.Nil(t, err)
	require.Equal(t, nodeStatusAfterUnJail, validatorsStatistics[blsKeys[0]].ValidatorStatus)
}
