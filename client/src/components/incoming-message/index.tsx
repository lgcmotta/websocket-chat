import { format, parseISO } from "date-fns"
import { FC } from "react"
import { IMessageReceivedProps } from "../../models/message"

const IncomingMessage: FC<IMessageReceivedProps> = ({ message }) => {
  const { sender, content, receivedAt, type } = message;

  const { color, receiver } = type == "broadcast"
    ? { color: "#818cf8", receiver: "@everyone" }
    : { color: "#6366f1", receiver: "@direct" };

  return (
    <div className="w-full">
      <div className={"m-2 ml-4 pt-2 pb-2 pl-4 pr-4 rounded-lg w-1/3 " + `bg-[${color}]`}>
        <div className="flex flex-row">
          <div className="w-full items-center flex flex-row justify-start">
            <span className="right-0 text-xs">{sender.nickname}
              <span className="pl-1 italic font-sans">to {receiver}</span>
            </span>
          </div>
          <div className="w-full items-center flex flex-row justify-end">
            <span className="right-0 text-xs">{format(parseISO(receivedAt), "HH:mm:ss")}</span>
          </div>
        </div>
        <div className="pt-2">
          <span className="text-md whitespace-pre-wrap">
            {content}
          </span>
        </div>
      </div>
    </div>
  )
}

export default IncomingMessage