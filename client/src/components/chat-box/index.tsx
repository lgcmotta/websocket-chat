import MessageInput from "../message-input"
import IncomingMessage from "../incoming-message"
import { useSelector } from "../../context/chat-context"

const ChatBox = () => {
  const messages = useSelector(state => state.messages)
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
        {/* TODO: map app messages */}
        {messages.map((message, index) => <IncomingMessage key={index} message={message} />)}
      </div>
      <MessageInput />
    </div>
  )
}

export default ChatBox