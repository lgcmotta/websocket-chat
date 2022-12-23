import { IMember } from "./member";


export enum MessageType {
  BROADCAST = "broadcast",
  DIRECFT = "direct"
}

export interface IMessageReceived {
  sender: IMember;
  receiver: IMember;
  content: string;
  receivedAt: Date;
  type: MessageType;
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