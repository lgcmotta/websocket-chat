import { FC, MouseEventHandler, createRef } from "react";
import { useChatContext } from "../../context/chat-context";
import { NicknameButton } from "./button";
import { NicknameBarDiv, NicknameButtonDiv } from "./divs";
import { NicknameHeader } from "./header";
import { NicknameInput } from "./input";

const NicknameBar: FC = () => {
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
    <NicknameBarDiv >
      <NicknameHeader text={title} />
      <NicknameInput ref={inputRef} />
      <NicknameButtonDiv>
        <NicknameButton label={buttonText} onClick={onClick} />
      </NicknameButtonDiv>
    </NicknameBarDiv>
  )
}

export default NicknameBar