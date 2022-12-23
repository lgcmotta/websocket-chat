import BroadcastIncomingMessage from "../broadcast-incoming-message"
import DirectIncomingMessage from "../direct-incoming-message"
import MessageInput from "../message-input"
import BroadcastOutgoingMessage from "../broadcast-outgoing-message"
import DirectOutgoingMessage from "../direct-outgoing-message"

const ChatBox = () => {
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
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
        <DirectIncomingMessage />
        <BroadcastOutgoingMessage />
        <BroadcastIncomingMessage />
        <DirectOutgoingMessage />
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
        <BroadcastIncomingMessage />
        <BroadcastOutgoingMessage />
      </div>
      <MessageInput />
    </div>
  )
}

export default ChatBox