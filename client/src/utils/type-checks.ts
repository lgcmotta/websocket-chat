import { IConnectedMembers } from "../models/member";
import { IMessageReceived } from "../models/message";

function isMessage(message: any): message is IMessageReceived {
  return "sender" in message &&
    "receiver" in message &&
    "content" in message &&
    "receivedAt" in message &&
    "type"  in message;
}

function isMembersList(members: any): members is IConnectedMembers {
  return "members" in members;
}

export { isMessage, isMembersList }