import { FC } from "react";
import { IMember } from "../../models/member";

interface IReceiversProps {
  members: IMember[]
}

const ReceiversLabel: FC<IReceiversProps> = ({ members }) => {
  const receivers = members.map(member => member.nickname).join(', ')

  return (
    <p className="ml-2 text-sm">To: {receivers}</p>
  )
}

export { ReceiversLabel }