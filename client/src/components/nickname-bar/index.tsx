import { FC, MouseEventHandler, createRef, useEffect } from "react";
import { websocketClient } from "../../api/ws";
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

  const handleSetNickname: MouseEventHandler<HTMLButtonElement> = () => {
    if (inputRef.current?.value == "") return;
    const nickname = inputRef.current?.value as string;
    setState(prev => {
      const { myself } = prev
      return { ...prev, myself: { ...myself, nickname: nickname } }
    })
  }

  return (
    <NicknameBarDiv >
      <NicknameHeader text={title} />
      <NicknameInput innerRef={inputRef} />
      <NicknameButtonDiv>
        <NicknameButton label={buttonText} onClick={handleSetNickname} />
      </NicknameButtonDiv>
    </NicknameBarDiv>
  )
}

export default NicknameBar