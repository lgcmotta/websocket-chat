import { createRef, FC, MouseEventHandler, useEffect } from "react";
import { useChatContext } from "../../context/chat-context";

const Nickname: FC = () => {
  const { state, setState } = useChatContext()
  const { myself } = state;

  const { title, buttonText } = myself.nickname == ""
    ? { title: "Choose your nickname:", buttonText: "Join!" }
    : { title: "Logged as:", buttonText: "Change nickname" };

  const inputRef = createRef<HTMLInputElement>();

  const onClick: MouseEventHandler<HTMLButtonElement> = (event) => {
    if (inputRef.current?.value != "") {
      const nickname = inputRef.current?.value as string;
      setState(prev => {
        const { myself } = prev
        const connectionId = !myself ? "" : myself.connectionId;
        return { ...prev, myself: { connectionId, nickname: nickname } }
      })
    }
  }

  return (
    <div className="mt-2 w-full items-start flex-col flex justify-start">
      <p className="ml-4 mr-4">{title}</p>
      <input className="mr-4
                        ml-4
                        mt-2
                        p-2
                        w-11/12
                        items-center
                        flex-col
                        flex 
                        justify-start 
                        rounded-md 
                        border-2 
                        border-white 
                        text-left
                        bg-[#4b5563]"
        ref={inputRef}
        type="text" />
      <div className="pr-4 pl-4 pt-2 w-full items-center flex-col flex justify-center">
        <button className="w-full mt-2 bg-[#7c3aed] rounded-md border-2"
          onClick={onClick}>{buttonText}</button>
      </div>
    </div>
  )
}

export default Nickname