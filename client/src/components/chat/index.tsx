import { createRef, FC, useEffect } from "react"
import { useChatContext } from "../../context/chat-context"
import { websocketClient } from "../../api/ws"
import { IMessageReceived } from "../../models/message"
import { isMembersList, isMessage } from "../../utils/type-checks"
import { IConnectedMembers } from "../../models/member"
import { MessagesContainer, MessagesBox, ScrollToLastMessage } from "./divs"
import MessageInput from "../message-input"
import ChatMessages from "../chat-messages"

const Chat: FC = () => {
  const { state, setState } = useChatContext()
  const { myself } = state;
  const scrollRef = createRef<HTMLDivElement>()

  useEffect(() => {
    if (myself.nickname == "") return;

    if (!websocketClient.isConnected()) {
      websocketClient.connect()
      websocketClient.onMessageReceived(onMessageReceived)
    }

    setTimeout(() => {
      websocketClient.join({ action: "join", nickname: myself.nickname })
    }, 3000)

  }, [myself.nickname])

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
        return {
          ...prev,
          myself: {
            ...prev.myself,
            connectionId: receiver.connectionId
          },
          messages: [...prev.messages, message]
        }
      })
    } else {
      setState(prev => {
        return { ...prev, messages: [...prev.messages, message] }
      })
    }
  }

  return (
    <MessagesContainer>
      <MessagesBox >
        <ChatMessages scrollRef={scrollRef} />
        <ScrollToLastMessage scrollRef={scrollRef} />
      </MessagesBox>
      <MessageInput />
    </MessagesContainer>
  )
}

export default Chat