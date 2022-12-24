import { FC } from "react";
import { format, parseISO } from "date-fns";
import { IMessageReceivedProps } from "../../models/message";

const OutgoingMessage: FC<IMessageReceivedProps> = ({ message }) => {
  const { receiver, content, receivedAt, type } = message;

  const from = type == "broadcast" ? "@everyone" : "@direct"

  return (
    <div className="w-full">
      <div className={"m-2 mr-4 pt-2 pb-2 pl-4 pr-4 rounded-lg w-1/3 float-right bg-[#52525b]"}>
        <div className="flex flex-row">
          <div className="w-full items-center flex flex-row justify-start">
            <span className="right-0 text-xs">{format(parseISO(receivedAt), "HH:mm:ss")}</span>
          </div>
          <div className="w-full items-center flex flex-row justify-end">
            <span className="right-0 text-xs">{receiver.nickname}
              <span className="pl-1 italic font-sans">{from}</span>
            </span>
          </div>
        </div>
        <div className="pt-2 float-right">
          <span className="text-md whitespace-pre-wrap ">
            {content}
          </span>
        </div>
      </div>
    </div>
  )
}

export default OutgoingMessage