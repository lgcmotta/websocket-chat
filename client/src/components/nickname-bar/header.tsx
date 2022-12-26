import { FC } from "react";

interface INicknameHeaderProps {
  text: string
}

const NicknameHeader: FC<INicknameHeaderProps> = ({ text }) => {
  return (
    <p className="ml-4 mr-4">{text}</p>
  )
}

export { NicknameHeader }