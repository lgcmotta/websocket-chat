import { FC } from "react";

interface ISendMessageButtonProps {
  onClick: () => void;
}

const SendMessageButton: FC<ISendMessageButtonProps> = ({ onClick }) => {
  return (
    <button className="ml-2 w-1/12 h-full bg-[#818cf8]" onClick={onClick}>
      Send
    </button>
  )
}

export { SendMessageButton }