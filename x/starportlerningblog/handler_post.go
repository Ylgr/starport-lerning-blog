package starportlerningblog

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ylgr/starport-lerning-blog/x/starportlerningblog/keeper"
	"github.com/ylgr/starport-lerning-blog/x/starportlerningblog/types"
)

func handleMsgCreatePost(ctx sdk.Context, k keeper.Keeper, msg *types.MsgCreatePost) (*sdk.Result, error) {
	k.CreatePost(ctx, *msg)

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}
