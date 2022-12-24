import { FC } from "react";
import { format, parseISO } from "date-fns";
import { IMessageReceivedProps } from "../../models/message";

const OutgoingMessage: FC<IMessageReceivedProps> = ({ message }) => {
  const { receiver, content, receivedAt, type } = message;

  const { color, from } = type == "broadcast"
    ? { color: "#4b5563", from: "@everyone" }
    : { color: "#9ca3af", from: "@direct" };

  const classes = `m-2 mr-4 pt-2 pb-2 pl-4 pr-4 bg-[${color}] rounded-lg w-1/3 float-right`
  return (
    <div className="w-full">
      <div className={classes}>
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