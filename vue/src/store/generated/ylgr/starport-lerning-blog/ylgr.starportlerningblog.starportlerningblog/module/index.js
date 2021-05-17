// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry } from "@cosmjs/proto-signing";
import { Api } from "./rest";
import { MsgCreateComment } from "./types/starportlerningblog/tx";
import { MsgUpdateComment } from "./types/starportlerningblog/tx";
import { MsgDeleteComment } from "./types/starportlerningblog/tx";
import { MsgCreatePost } from "./types/starportlerningblog/post";
const types = [
    ["/ylgr.starportlerningblog.starportlerningblog.MsgCreateComment", MsgCreateComment],
    ["/ylgr.starportlerningblog.starportlerningblog.MsgUpdateComment", MsgUpdateComment],
    ["/ylgr.starportlerningblog.starportlerningblog.MsgDeleteComment", MsgDeleteComment],
    ["/ylgr.starportlerningblog.starportlerningblog.MsgCreatePost", MsgCreatePost],
];
const registry = new Registry(types);
const defaultFee = {
    amount: [],
    gas: "200000",
};
const txClient = async (wallet, { addr: addr } = { addr: "http://localhost:26657" }) => {
    if (!wallet)
        throw new Error("wallet is required");
    const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
    const { address } = (await wallet.getAccounts())[0];
    return {
        signAndBroadcast: (msgs, { fee = defaultFee, memo = null }) => memo ? client.signAndBroadcast(address, msgs, fee, memo) : client.signAndBroadcast(address, msgs, fee),
        msgCreateComment: (data) => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgCreateComment", value: data }),
        msgUpdateComment: (data) => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgUpdateComment", value: data }),
        msgDeleteComment: (data) => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgDeleteComment", value: data }),
        msgCreatePost: (data) => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgCreatePost", value: data }),
    };
};
const queryClient = async ({ addr: addr } = { addr: "http://localhost:1317" }) => {
    return new Api({ baseUrl: addr });
};
export { txClient, queryClient, };
