import { FC, PropsWithChildren } from "react"

const NicknameBarDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="mt-2 w-full items-start flex-col flex justify-start">
      {children}
    </div>
  )
}

const NicknameButtonDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="pr-4 pl-4 pt-2 w-full items-center flex-col flex justify-center">
      {children}
    </div>
  )
}

export { NicknameBarDiv, NicknameButtonDiv }