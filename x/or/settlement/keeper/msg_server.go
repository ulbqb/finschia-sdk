package keeper

import (
	"context"
	"encoding/hex"

	sdk "github.com/Finschia/finschia-sdk/types"
	"github.com/Finschia/finschia-sdk/x/or/settlement/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (s msgServer) StartChallenge(c context.Context, req *types.MsgStartChallenge) (*types.MsgStartChallengeResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)

	startState, _ := hex.DecodeString("e4cff37938153c13e1d5347e0889c4bf846099c7e6e6b4946a7d898f7eb56b37") // start state of bc (demo application)
	// TODO:
	// - get start state from rollup module
	// - validate if rollup name exist
	// - validate if challenger exist in rollup
	// - validate if defender exist in rollup
	// - validate if block height exist in rollup blocks
	challenge := types.Challenge{
		L:                   0,
		R:                   req.StepCount,
		AssertedStateHashes: map[uint64][]byte{0: startState[:]},
		DefendedStateHashes: map[uint64][]byte{0: startState[:]},
		Challenger:          req.From,
		Defender:            req.To,
		BlockHeight:         req.BlockHeight,
		RollupName:          req.RollupName,
	}

	if s.Keeper.HasChallenge(ctx, challenge.ID()) {
		return nil, types.ErrChallengeExists
	}
	s.Keeper.SetChallenge(ctx, challenge.ID(), challenge)

	return &types.MsgStartChallengeResponse{
		ChallengeId: challenge.ID(),
	}, nil
}

func (s msgServer) NsectChallenge(c context.Context, req *types.MsgNsectChallenge) (*types.MsgNsectChallengeResponse, error) {
	panic("implement me")
}

func (s msgServer) FinishChallenge(c context.Context, req *types.MsgFinishChallenge) (*types.MsgFinishChallengeResponse, error) {
	panic("implement me")
}
