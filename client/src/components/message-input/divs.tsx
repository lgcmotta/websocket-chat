import { FC, PropsWithChildren } from "react"

const MessageInputDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="p-2 w-full h-48 items-center flex flex-col justify-center">
      {children}
    </div>
  )
}

const ActionsBarDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="pb-2 flex flex-row justify-start items-center w-full">
      {children}
    </div>
  )
}

const MessageContentDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="w-full h-48 items-center flex flex-row justify-center">
      {children}
    </div>
  )
}

export { MessageInputDiv, ActionsBarDiv, MessageContentDiv }