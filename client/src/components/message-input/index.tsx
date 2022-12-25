import { FC, useState } from "react"
import { useChatContext } from "../../context/chat-context"
import TextInput from 'react-autocomplete-input';
import { EmojiClickData } from 'emoji-picker-react';
import { IMember } from "../../models/member";
import { websocketClient } from "../../api/ws";
import { ActionsBarDiv, MessageContentDiv, MessageInputDiv } from "./divs";
import { EmojisBar } from "./emojis";
import 'react-autocomplete-input/dist/bundle.css';
import "./index.css"

const MessageInput: FC = () => {
  const { state } = useChatContext()
  const { members, myself, everyone } = state;

  const [receivers, setReceivers] = useState<IMember[]>([everyone])
  const [content, setContent] = useState<string>("")

  const availableMentions = [
    everyone.nickname,
    ...members.filter(member => member.nickname != myself.nickname).map(member => member.nickname)
  ]

  const handleMention = (content: string) => {

    const mentions = findMentionsInContent(content)

    setReceivers(mentions)
  }

  const handleMessageContentChange = (content: string) => {
    setContent(content)

    const mentions = findMentionsInContent(content)

    setReceivers(mentions)
  }

  const findMentionsInContent = (content: string): IMember[] => {
    const mentionsRegex = /\B@\w+/g;

    const matches = content.match(mentionsRegex)

    if (!matches || matches == null) return [everyone];

    if (matches.some(match => match.replace("@", "") == everyone.nickname)
      || (members.length > 2 && members
        .filter(member => member.connectionId != myself.connectionId)
        .every(member => matches.some(m => m.replace("@", "") == member.nickname)))) {
      return [everyone]
    }

    const mentions = members.filter(member => matches.some(m => m.replace("@", "") == member.nickname))

    return mentions
  }

  const handleSendMessage = () => {
    if (!websocketClient.isConnected()) {
      window.alert("You're not connected to the chat")
      return;
    }

    const message = content

    if (receivers.includes(everyone)) {
      websocketClient.broadcast({
        action: "broadcast",
        content: message ?? ""
      })
      clearContent()
      return;
    }

    websocketClient.direct({
      action: "direct",
      receivers: receivers.map(receiver => receiver.connectionId),
      content: message ?? ""
    })


    clearContent()

    setReceivers([everyone])
  }

  const handleEmojiClick = (emojiData: EmojiClickData) => {
    setContent(prev => prev + emojiData.emoji)
  }

  const clearContent = () => setContent("")

  return (
    <MessageInputDiv>
      <ActionsBarDiv>
        <EmojisBar onEmojiSelected={handleEmojiClick} />
        <p className="ml-2 text-sm">To: {receivers.map(receiver => receiver.nickname).join(', ')}</p>
      </ActionsBarDiv>
      <MessageContentDiv>
        <TextInput className="h-full w-full p-4 rounded-md border-2 border-white bg-[#4b5563] text-sm"
          value={content}
          offsetY={-30}
          trigger={["@"]}
          onSelect={handleMention}
          onChange={handleMessageContentChange}
          options={{ "@": availableMentions }} />
        <button className="ml-2 w-1/12 h-full bg-[#818cf8]" onClick={handleSendMessage}>Send</button>
      </MessageContentDiv>
    </MessageInputDiv>
  )
}

export default MessageInput