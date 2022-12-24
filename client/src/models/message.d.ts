import { IMember } from "./member";

export interface IMessageReceived {
  sender: IMember;
  receiver: IMember;
  content: string;
  receivedAt: string;
  type: "broadcast" | "direct";
}

export interface IMessageReceivedProps {
  message: IMessageReceived;
}

type Action = "join" | "broadcast" | "direct";
export interface IMessage {
  action: Action;
}

export interface IJoinMessage extends IMessage{
  action: "join";
  nickname: string;
}

export interface IBroadcastMessage extends IMessage{
  action: "broadcast";
  content: string;
}

export interface IDirectMessage extends IMessage{
  action: "direct";
  content: string;
  receiver: String;
}