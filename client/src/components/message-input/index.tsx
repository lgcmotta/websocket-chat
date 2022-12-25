import { FC, createRef, useState, RefObject } from "react"
import { useChatContext } from "../../context/chat-context"
import TextInput from 'react-autocomplete-input';
import EmojiPicker, { EmojiClickData, Theme } from 'emoji-picker-react';
import search from "@jukben/emoji-search"
import 'react-autocomplete-input/dist/bundle.css';
import "./index.css"
import { IMember } from "../../models/member";
import { websocketClient } from "../../api/ws";

const MessageInput: FC = () => {
  const textRef = createRef<HTMLTextAreaElement>()
  const navRef = createRef<HTMLDivElement>();
  const { state } = useChatContext()
  const { members, myself, everyone } = state;

  const [showEmojis, setShowEmojis] = useState<boolean>(false)
  const [receiver, setReceiver] = useState<IMember>(everyone)
  const [content, setContent] = useState<string>("")

  const mentions = [everyone.nickname, ...members.filter(member => member.nickname != myself.nickname).map(member => member.nickname)]

  const handleMention = (receiver: string) => {
    const re = /\B@\w+/g;

    const matches = receiver.match(re)

    if (!matches || matches == null) return;

    const to = members.find(member => member.nickname == matches[0].replace("@", ""))

    if (to) {
      setReceiver(to)
    }
  }

  const handleMessageContentChange = (e: string) => {
    setContent(e)

    if (!receiver || e.includes(`@${receiver.nickname}`)) return;

    setReceiver(everyone)
  }

  const handleSendMessage = () => {
    const ref = ((textRef.current as any).refInput as RefObject<HTMLTextAreaElement>)
    if (!websocketClient.isConnected()) {
      window.alert("You're not connected to the chat")
      return;
    }

    const message = ref.current?.value

    if (receiver.nickname == everyone.nickname) {
      websocketClient.broadcast({
        action: "broadcast",
        content: message ?? ""
      })
      clearContent()
      return;
    }

    websocketClient.direct({
      action: "direct",
      receiver: receiver.connectionId,
      content: message ?? ""
    })

    clearContent()

    setReceiver(everyone)
  }

  const handleShowEmoji = () => setShowEmojis(prev => !prev)

  const handleEmojiClick = (emojiData: EmojiClickData) => {
    setContent(prev => prev + emojiData.emoji)
    setShowEmojis(false)
  }

  const clearContent = () => setContent("")

  return (
    <div className="p-2 w-full h-48 items-center flex flex-col justify-center">
      <div ref={navRef} className="pb-2 flex flex-row justify-start items-center w-full">
        <button className="bg-transparent border-2 rounded-md" onClick={handleShowEmoji}>
          {search("smile")[0].char}
        </button>
        <p className="ml-2 text-sm">To: {receiver.nickname}</p>
        {showEmojis && <EmojiPicker theme={Theme.DARK} width="400px" onEmojiClick={handleEmojiClick} />}
      </div>
      <div className="w-full h-48 items-center flex flex-row justify-center">
        <TextInput

          className="h-full
                    w-full
                    p-4
                    rounded-md
                    border-2
                    border-white
                    bg-[#4b5563]
                    text-sm"
          value={content}
          offsetY={-30}
          trigger={["@"]}
          onSelect={handleMention}
          onChange={handleMessageContentChange}
          options={{
            "@": mentions,
          }}
          ref={textRef} />
        <button className="ml-2 w-1/12 h-full bg-[#818cf8]" onClick={handleSendMessage}>Send</button>
      </div>
    </div>
  )
}

export default MessageInput