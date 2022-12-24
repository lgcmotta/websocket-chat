import { FC, createRef, useState, RefObject } from "react"
import { useChatContext } from "../../context/chat-context"
import TextInput from 'react-autocomplete-input';
import EmojiPicker, { Theme } from 'emoji-picker-react';
import search from "@jukben/emoji-search"
import 'react-autocomplete-input/dist/bundle.css';
import "./index.css"
import { IMember } from "../../models/member";
import { websocketClient } from "../../api/ws";

const MessageInput: FC = () => {
  const textRef = createRef<HTMLTextAreaElement>()
  const navRef = createRef<HTMLDivElement>();
  const everyone = { nickname: "@everyone", connectionId: "" }


  const [showEmojis, setShowEmojis] = useState<boolean>(false)
  const [receiver, setReceiver] = useState<IMember>(everyone)
  const [content, setContent] = useState<string>("")
  const { state } = useChatContext()
  const { members, myself } = state;

  const onSelect = (receiver: string) => {
    const nickname = receiver.replace("@", "")
    const to = members.find(member => member.nickname == nickname)
    if (to) {
      setReceiver(to)
    }

  }

  const onChange = (e: string) => {
    setContent(e)

    if (!receiver || e.includes(`@${receiver.nickname}`)) return;

    setReceiver(everyone)
  }


  const handleShowEmoji = (e: any) => {
    setShowEmojis(prev => !prev)
  }

  const handleSend = () => {
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

  }

  const clearContent = () => setContent("")

  return (
    <div className="p-2 w-full h-48 items-center flex flex-col justify-center">
      <div ref={navRef} className="pb-2 flex flex-row justify-start items-center w-full">
        <button className="bg-transparent border-2 rounded-md" onClick={handleShowEmoji}>
          {search("smile")[0].char}
        </button>
        <p className="ml-2 text-sm">To: {receiver.nickname}</p>
        {showEmojis && <EmojiPicker theme={Theme.DARK} width="400px" />}
      </div>
      <div className="w-full h-48 items-center flex flex-row justify-center">
        <TextInput
          spacer=""
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
          onSelect={onSelect}
          onChange={onChange}
          options={{
            "@": members.filter(m => m.connectionId != myself.connectionId).map(m => m.nickname),
          }}
          ref={textRef} />
        <button className="ml-2 w-1/12 h-full bg-[#818cf8]" onClick={handleSend}>Send</button>
      </div>
    </div>
  )
}

export default MessageInput