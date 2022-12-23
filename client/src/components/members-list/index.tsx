import React, { useState } from "react";
import { IMember } from "../../models/member";
import MemberCard from "../member-card";
import Nickname from "../nickname";

const MembersList = () => {
  const [members, setMembers] = useState<IMember[]>([])


  return (
    <div className="w-1/4 items-start flex-col flex justify-start rounded-md border-2 border-white overflow-hidden">
      <Nickname />
      <div className="mt-4 ml-4 mr-4 w-11/12 border-2 rounded-full opacity-80"></div>
      <p className="pt-4 ml-4 mr-4">Online Members:</p>
      <div className="w-full overflow-y-auto overflow-x-hidden h-full">
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
        <MemberCard />
      </div>

    </div>
  )
}

export default MembersList