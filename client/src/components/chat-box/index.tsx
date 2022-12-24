import MessagesBox from "../messages-box"
import MessageInput from "../message-input"
import IncomingMessage from "../incoming-message"
import { useChatContext } from "../../context/chat-context"
import { useEffect } from "react"
import { websocketClient } from "../../api/ws"
import { IMessageReceived } from "../../models/message"
import MessagesContainer from "../messages-container"
import OutgoingMessage from "../outgoing-message"
import SystemMessage from "../system-message"

const ChatBox = () => {
  const { state, setState } = useChatContext()
  const { myself, messages } = state;

  useEffect(() => {
    if (myself.nickname == "" || websocketClient.isConnected()) return;

    websocketClient.connect()
    websocketClient.onMessageReceived(onMessageReceived)
    setTimeout(() => {
      websocketClient.join({ action: "join", nickname: myself.nickname })
    }, 3000)

  }, [myself])

  const onMessageReceived = (event: any) => {
    const message = JSON.parse(event.data) as IMessageReceived
    const { receiver } = message
    if (receiver.nickname == myself.nickname && myself.connectionId == "") {
      setState(prev => {
        return { ...prev, myself: { ...prev.myself, connectionId: receiver.connectionId }, messages: [...prev.messages, message] }
      })
    } else {
      setState(prev => {
        return { ...prev, messages: [...prev.messages, message] }
      })
    }
  }

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
    <MessagesContainer>
      <MessagesBox>
        {renderMessages()}
      </MessagesBox>
      <MessageInput />
    </MessagesContainer>
  )
}

export default ChatBox