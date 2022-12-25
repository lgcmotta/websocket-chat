import { FC, useState } from "react"
import search from "@jukben/emoji-search"
import EmojiPicker, { EmojiClickData, Theme } from "emoji-picker-react";

interface IEmojisBarProps {
  onEmojiSelected: (emojiData: EmojiClickData) => void;
}

const EmojisBar: FC<IEmojisBarProps> = ({ onEmojiSelected }) => {
  const [showEmojis, setShowEmojis] = useState<boolean>(false)
  const handleEmojiButtonClick = () => setShowEmojis(prev => !prev)
  const handleEmojiSelected = (emojiData: EmojiClickData) => {
    onEmojiSelected(emojiData)
    setShowEmojis(false)
  }
  return (
    <>
      <EmojisButton onEmojiButtonClick={handleEmojiButtonClick} />
      {showEmojis && <EmojiPicker theme={Theme.DARK} width="400px" onEmojiClick={handleEmojiSelected} />}
    </>
  )
}


interface IEmojisButtonProps {
  onEmojiButtonClick: () => void;
}

const EmojisButton: FC<IEmojisButtonProps> = ({ onEmojiButtonClick }) => {
  return (
    <button className="bg-transparent border-2 rounded-md" onClick={onEmojiButtonClick}>
      {search("smile")[0].char}
    </button>
  )
}

export { EmojisBar }