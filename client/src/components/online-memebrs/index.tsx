import { FC } from "react";
import { useChatContext } from "../../context/chat-context";
import MemberCard from "../member-card";

const OnlineMembers: FC = () => {
  const { state } = useChatContext()

  const { members } = state

  const renderMembers = (): JSX.Element[] => {
    if (!members || members.length == 0) return []

    return members.map(member => {
      return <MemberCard key={member.connectionId} member={member} />;
    });
  }

  return (
    <>
      {renderMembers()}
    </>
  )
}

export default OnlineMembers;