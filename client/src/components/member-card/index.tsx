import { FC, Fragment } from "react"
import { IMemberCardProps } from "../../models/member"
import { MemberCardDiv, StatusDiv } from "./divs";
import { MemberName, Status } from "./member";

const MemberCard: FC<IMemberCardProps> = ({ member }) => {
  const { nickname } = member;

  if (!nickname || nickname == "") {
    return (
      <Fragment />
    )
  }

  return (
    <MemberCardDiv>
      <MemberName nickname={nickname} />
      <StatusDiv>
        <Status />
      </StatusDiv>
    </MemberCardDiv>
  )
}

export default MemberCard

