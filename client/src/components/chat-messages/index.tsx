import { FC } from "react";
import { useChatContext } from "../../context/chat-context";
import IncomingMessage from "../incoming-message";
import OutgoingMessage from "../outgoing-message";
import SystemMessage from "../system-message";

const ChatMessages: FC = () => {
  const { state, } = useChatContext()
  const { myself, messages } = state;

  const renderMessages = (): JSX.Element[] => {
    if (myself.connectionId == "") return []

    return messages.map((message, index) => {
      if (message.type == "system") {
        return <SystemMessage key={index} message={message} />
      }

      if (message.sender.connectionId == myself.connectionId) {
        return <OutgoingMessage key={index} message={message} />
      }

      return <IncomingMessage key={index} message={message} />
    });
  }

  return (
    <>
      {renderMessages()}
    </>
  )
}

export default ChatMessages;