import { FC, PropsWithChildren } from "react";

const SideBarDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="w-1/4 items-start flex-col flex justify-start rounded-md border-2 border-white overflow-hidden">
      {children}
    </div>
  )
}

const Separator: FC = () => {
  return (
    <div className="mt-4 ml-4 mr-4 w-11/12 border-2 rounded-full opacity-80">
    </div>
  )
}

const OnlineMembersDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="w-full overflow-y-auto overflow-x-hidden h-full">
      {children}
    </div>
  )
}

const OnlineMembersParagraph: FC = () => {
  return (
    <p className="pt-4 ml-4 mr-4">Online Members:</p>
  )
}

export { SideBarDiv, Separator, OnlineMembersDiv, OnlineMembersParagraph }