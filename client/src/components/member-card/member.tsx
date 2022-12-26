import { FC } from "react";

interface IMemberNameProps {
  nickname: string
}

const MemberName: FC<IMemberNameProps> = ({ nickname }) => {
  return (
    <p className="ml-2 mr-2">{nickname}</p>
  )
}

const Status: FC = () => {
  return (
    <div className="bg-[#84cc16] rounded-full w-3 h-3 animate-pulse" />
  )
}

export { MemberName, Status }

