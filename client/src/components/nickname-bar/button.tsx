import { FC, MouseEventHandler } from "react"

interface INicknameButtonProps {
  label: string;
  onClick: MouseEventHandler<HTMLButtonElement>;
}

const NicknameButton: FC<INicknameButtonProps> = ({ label, onClick }) => {
  return (
    <button className="w-full mt-2 bg-[#818cf8] rounded-md border-2" onClick={onClick}>
      {label}
    </button>
  )
}

export { NicknameButton }