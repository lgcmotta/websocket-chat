import { FC } from "react";
import { useChatContext } from "../../context/chat-context";
import { IMember } from "../../models/member";
import TextInput from 'react-autocomplete-input';
import 'react-autocomplete-input/dist/bundle.css';
import "./index.css"

interface IMessageTextAreaProps {
  content: string
  onContentChanged: (content: string) => void
  onMentionsChanged: (mentions: IMember[]) => void
}

const MessageTextArea: FC<IMessageTextAreaProps> = ({ content, onContentChanged, onMentionsChanged }) => {
  const { state } = useChatContext();
  const { members, myself, everyone } = state;

  const availableMentions = [
    everyone.nickname,
    ...members.filter(member => member.nickname != myself.nickname).map(member => member.nickname)
  ]

  const handleMessageContentChange = (content: string) => {
    const mentions = findMentionsInContent(content)

    onContentChanged(content)

    onMentionsChanged(mentions)
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

  return (
    <TextInput className="h-full w-full p-4 rounded-md border-2 border-white bg-[#4b5563] text-sm"
      value={content}
      offsetY={-30}
      trigger={["@"]}
      onChange={handleMessageContentChange}
      options={{ "@": availableMentions }} />
  )
}

export { MessageTextArea }