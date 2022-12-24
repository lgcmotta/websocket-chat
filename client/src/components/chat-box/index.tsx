import MessageInput from "../message-input"
import IncomingMessage from "../incoming-message"
import { useChatContext } from "../../context/chat-context"
import { useEffect } from "react"
import { websocketClient } from "../../api/ws"
import { IMessageReceived } from "../../models/message"

const ChatBox = () => {
  const { state, setState } = useChatContext()
  const { myself, messages } = state;

  useEffect(() => {
    if (myself.nickname == "" || websocketClient.isConnected()) return;

    websocketClient.connect()
    websocketClient.onMessageReceived(event => {
      const message = JSON.parse(event.data) as IMessageReceived
      const { sender } = message
      if (sender.nickname == myself.nickname && myself.connectionId == "") {
        setState(prev => {
          return { ...prev, myself: { ...prev.myself, connectionId: sender.connectionId }, messages: [...prev.messages, message] }
        })
      } else {
        setState(prev => {
          return { ...prev, messages: [...prev.messages, message] }
        })
      }
    })
    setTimeout(() => {
      websocketClient.join({ action: "join", nickname: myself.nickname })
    }, 5000)

  }, [myself])


  return (
    <div className="w-full
                    h-full
                    mr-4
                    rounded-md
                    border-2
                    border-white
                    items-center
                    flex
                    flex-col
                    justify-between 
                    justify-center">
      <div className="h-full
                      w-full
                      pt-4
                      flex
                      flex-col
                      items-center 
                      rounded-md
                      border-b-2
                      border-white
                      bg-[#18181b]
                      overflow-x-hidden
                      overflow-y-auto">
        {messages.map((message, index) => <IncomingMessage key={index} message={message} />)}
      </div>
      <MessageInput />
    </div>
  )
}

export default ChatBox