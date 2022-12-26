import { FC, useState } from "react"
import { useChatContext } from "../../context/chat-context"
import { EmojiClickData } from 'emoji-picker-react';
import { IMember } from "../../models/member";
import { websocketClient } from "../../api/ws";
import { ActionsBarDiv, MessageContentDiv, MessageInputDiv } from "./divs";
import { EmojisBar } from "./emojis";
import { SendMessageButton } from "./button";
import { ReceiversLabel } from "./receivers";
import { MessageTextArea } from "./input";

const MessageInput: FC = () => {
  const { state } = useChatContext()
  const { everyone } = state;

  const [receivers, setReceivers] = useState<IMember[]>([everyone]);

  const [content, setContent] = useState<string>("");

  const handleMessageContentChange = (content: string) => setContent(content);

  const handleMentionsChange = (mentions: IMember[]) => setReceivers(mentions);

  const handleEmojiClick = (emojiData: EmojiClickData) => setContent(prev => prev + emojiData.emoji);

  const handleSendMessage = () => {
    if (content == "") return;

    if (!websocketClient.isConnected()) {
      window.alert("You're not connected to the chat");
      return;
    }

    if (receivers.includes(everyone)) {
      websocketClient.broadcast({
        action: "broadcast",
        content: content ?? ""
      })
      setContent("");
      return;
    }

    websocketClient.direct({
      action: "direct",
      receivers: receivers.map(receiver => receiver.connectionId),
      content: content ?? ""
    });

    setContent("");
    setReceivers([everyone]);
  }

  return (
    <MessageInputDiv>
      <ActionsBarDiv>
        <EmojisBar onEmojiSelected={handleEmojiClick} />
        <ReceiversLabel members={receivers} />
      </ActionsBarDiv>
      <MessageContentDiv>
        <MessageTextArea content={content}
          onContentChanged={handleMessageContentChange}
          onMentionsChanged={handleMentionsChange} />
        <SendMessageButton onClick={handleSendMessage} />
      </MessageContentDiv>
    </MessageInputDiv>
  )
}

export default MessageInput