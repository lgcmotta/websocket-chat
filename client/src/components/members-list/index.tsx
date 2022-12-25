import { FC, useEffect } from "react";
import { websocketClient } from "../../api/ws";
import { useChatContext } from "../../context/chat-context";
import MemberCard from "../member-card";
import Nickname from "../nickname";

const MembersList: FC = () => {
  const { state } = useChatContext()

  const { members } = state

  useEffect(() => {
    const interval = setInterval(() => {
      if (!websocketClient.isConnected()) return;
      websocketClient.requestMembersList()
    }, 5000);
    return () => clearInterval(interval);
  }, [])


  const renderMembers = (): JSX.Element[] => {
    if (!members || members.length == 0) return []

    return members.map(member => {
      return <MemberCard key={member.connectionId} member={member} />;
    });
  }

  return (
    <div className="w-1/4 items-start flex-col flex justify-start rounded-md border-2 border-white overflow-hidden">
      <Nickname />
      <div className="mt-4 ml-4 mr-4 w-11/12 border-2 rounded-full opacity-80"></div>
      <p className="pt-4 ml-4 mr-4">Online Members:</p>
      <div className="w-full overflow-y-auto overflow-x-hidden h-full">
        {renderMembers()}
      </div>
    </div>
  )
}

export default MembersList