import { FC, RefObject } from "react"

interface INicknameInputProps {
  ref: RefObject<HTMLInputElement>
}

const NicknameInput: FC<INicknameInputProps> = ({ ref }) => {
  return (
    <input className="mr-4
                        ml-4
                        mt-2
                        p-2
                        w-11/12
                        items-center
                        flex-col
                        flex 
                        justify-start 
                        rounded-md 
                        border-2 
                        border-white 
                        text-left
                        bg-[#4b5563]"
      ref={ref}
      type="text" />
  )
}

export { NicknameInput }