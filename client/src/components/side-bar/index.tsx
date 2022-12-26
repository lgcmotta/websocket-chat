import { FC, useEffect } from "react";
import { websocketClient } from "../../api/ws";
import { SideBarDiv, Separator, OnlineMembersDiv, OnlineMembersParagraph } from "./divs";
import NicknameBar from "../nickname-bar";
import OnlineMembers from "../online-memebrs";

const SideBar: FC = () => {

  useEffect(() => {
    const interval = setInterval(() => {
      if (!websocketClient.isConnected()) return;
      websocketClient.requestMembersList()
    }, 5000);
    return () => clearInterval(interval);
  }, [])

  return (
    <SideBarDiv>
      <NicknameBar />
      <Separator />
      <OnlineMembersParagraph />
      <OnlineMembersDiv>
        <OnlineMembers />
      </OnlineMembersDiv>
    </SideBarDiv>
  )
}

export default SideBar