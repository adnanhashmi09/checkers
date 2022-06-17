package types

import (
	"testing"

	"github.com/adnanhashmi09/checkers/x/checkers/rules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

const (
	alice = "cosmos1dyp3ggmle7gxfu3qu34m3zjctqt52lx06gppg9"
	bob = "cosmos1trz6t6926ypuvecpq2av4xmxsv4rw6fjyl7jyr"
)
func GetStoredGame1() *StoredGame {
	return &StoredGame{
		Creator: alice,
		Black:   bob,
		Red:     alice,
		Index:   "1",
		Game:    rules.New().String(),
		Turn:    "b",
	}
}

func TestCanGetAddressCreator(t *testing.T) {
    aliceAddress, err1 := sdk.AccAddressFromBech32(alice)
    creator, err2 := GetStoredGame1().GetCreatorAddress()
    require.Equal(t, aliceAddress, creator)
    require.Nil(t, err1)
    require.Nil(t, err2)
}

func TestGetAddressWrongCreator(t *testing.T) {
    storedGame := GetStoredGame1()
    storedGame.Creator = "cosmos148vd4qw4lqht5nz23m25kaxxkm4j4xg92s6a09"
    creator, err := storedGame.GetCreatorAddress()
    require.Nil(t, creator)
    require.EqualError(t,
        err,
        "creator address is invalid: cosmos148vd4qw4lqht5nz23m25kaxxkm4j4xg92s6a09: decoding bech32 failed: invalid checksum (expected 2s6a08 got 2s6a09)")
    require.EqualError(t, storedGame.Validate(), err.Error())
}

func TestCanGetAddressBlack(t *testing.T) {
	bobAddress, err1 := sdk.AccAddressFromBech32(bob)
	black, err2 := GetStoredGame1().GetBlackAddress()
	require.Equal(t, bobAddress, black)
	require.Nil(t, err2)
	require.Nil(t, err1)
}

func TestGetAddressWrongBlack(t *testing.T) {
	storedGame := GetStoredGame1()
	storedGame.Black = "cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h"
	black, err := storedGame.GetBlackAddress()
	require.Nil(t, black)
	require.EqualError(t,
		err,
		"black address is invalid: cosmos1xyxs3skf3f4jfqeuv89yyaqvjc6lffavxqhc8h: decoding bech32 failed: invalid checksum (expected xqhc8g got xqhc8h)")
	require.EqualError(t, storedGame.Validate(), err.Error())
}
