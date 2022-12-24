import { FC } from "react";
import { IMessageReceivedProps } from "../../models/message";


const SystemMessage: FC<IMessageReceivedProps> = ({ message }) => {
  const { content, sender, receiver } = message
  return (
    <div className="m-2 bg-[#737373] rounded-md w-1/4 flex flex-col items-center justify-center ">
      <span className="text-xs">from {sender.nickname} to {receiver.nickname}</span>
      <span className="text-base">{content}</span>
    </div>
  )
}

export default SystemMessage