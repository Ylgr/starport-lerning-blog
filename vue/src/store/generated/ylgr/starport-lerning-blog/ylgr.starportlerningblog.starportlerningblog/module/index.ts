// THIS FILE IS GENERATED AUTOMATICALLY. DO NOT MODIFY.

import { StdFee } from "@cosmjs/launchpad";
import { SigningStargateClient } from "@cosmjs/stargate";
import { Registry, OfflineSigner, EncodeObject, DirectSecp256k1HdWallet } from "@cosmjs/proto-signing";
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

const registry = new Registry(<any>types);

const defaultFee = {
  amount: [],
  gas: "200000",
};

interface TxClientOptions {
  addr: string
}

interface SignAndBroadcastOptions {
  fee: StdFee,
  memo?: string
}

const txClient = async (wallet: OfflineSigner, { addr: addr }: TxClientOptions = { addr: "http://localhost:26657" }) => {
  if (!wallet) throw new Error("wallet is required");

  const client = await SigningStargateClient.connectWithSigner(addr, wallet, { registry });
  const { address } = (await wallet.getAccounts())[0];

  return {
    signAndBroadcast: (msgs: EncodeObject[], { fee=defaultFee, memo=null }: SignAndBroadcastOptions) => memo?client.signAndBroadcast(address, msgs, fee,memo):client.signAndBroadcast(address, msgs, fee),
    msgCreateComment: (data: MsgCreateComment): EncodeObject => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgCreateComment", value: data }),
    msgUpdateComment: (data: MsgUpdateComment): EncodeObject => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgUpdateComment", value: data }),
    msgDeleteComment: (data: MsgDeleteComment): EncodeObject => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgDeleteComment", value: data }),
    msgCreatePost: (data: MsgCreatePost): EncodeObject => ({ typeUrl: "/ylgr.starportlerningblog.starportlerningblog.MsgCreatePost", value: data }),
    
  };
};

interface QueryClientOptions {
  addr: string
}

const queryClient = async ({ addr: addr }: QueryClientOptions = { addr: "http://localhost:1317" }) => {
  return new Api({ baseUrl: addr });
};

export {
  txClient,
  queryClient,
};
