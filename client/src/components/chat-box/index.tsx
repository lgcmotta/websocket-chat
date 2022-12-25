import { MessagesContainer, MessagesBox } from "./divs"
import { FC, useEffect } from "react"
import { useChatContext } from "../../context/chat-context"
import { websocketClient } from "../../api/ws"
import { IMessageReceived } from "../../models/message"
import { isMembersList, isMessage } from "../../utils/type-checks"
import { IConnectedMembers } from "../../models/member"
import MessageInput from "../message-input"
import ChatMessages from "../chat-messages"

const ChatBox: FC = () => {
  const { state, setState } = useChatContext()
  const { myself } = state;

  useEffect(() => {
    if (myself.nickname == "" || websocketClient.isConnected()) return;

    websocketClient.connect()
    websocketClient.onMessageReceived(onMessageReceived)
    setTimeout(() => {
      websocketClient.join({ action: "join", nickname: myself.nickname })
    }, 3000)

  }, [myself])

  const onMessageReceived = (event: any) => {
    const received = JSON.parse(event.data)

    if (isMessage(received)) {
      const message = received as IMessageReceived
      handleMessage(message)
      return;
    }

    if (isMembersList(received)) {
      const connectedMembers = received as IConnectedMembers;

      setState(prev => {
        return { ...prev, members: connectedMembers.members }
      })
    }
  }

  const handleMessage = (message: IMessageReceived) => {
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

  return (
    <MessagesContainer>
      <MessagesBox>
        <ChatMessages />
      </MessagesBox>
      <MessageInput />
    </MessagesContainer>
  )
}

export default ChatBox