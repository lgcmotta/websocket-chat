import { FC } from "react"
import { IMemberCardProps } from "../../models/member"

const MemberCard: FC<IMemberCardProps> = ({ member }) => {
  const { nickname } = member;

  if (!nickname || nickname == "") {
    return (
      <div />
    )
  }

  return (
    <div className="ml-4 mr-4 p-2 m-1 w-11/12 items-center flex-row flex justify-center justify-between rounded-md border-2 border-white bg-[#4b5563]">
      <p className="ml-2 mr-2">{nickname}</p>
      <div className="ml-2 flex w-3 h-3 mr-2 items-center flex-row flex justify-center">
        <div className="bg-[#84cc16] rounded-full w-3 h-3 animate-pulse" />
      </div>
    </div>
  )
}

export default MemberCard

