import { FC } from "react";
import ChatBox from "../components/chat-box";
import MembersList from "../components/members-list";

const Chat: FC = () => {
  return (
    <>
      <ChatBox />
      <MembersList />
    </>
  )
}

export default Chat;