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

export interface IMessage {
  action: "join" | "direct" | "broadcast";
}

export interface IBroadcastMessage extends IMessage {
  content: string;
}

export interface IDirectMessage extends IMessage {
  content: string;
  receiver: String;
}