import { FC, PropsWithChildren } from "react";

const MemberCardDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="ml-4 mr-4 p-2 m-1 w-11/12 items-center flex-row flex justify-center justify-between rounded-md border-2 border-white bg-[#4b5563]">
      {children}
    </div>
  )
}

const StatusDiv: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="ml-2 flex w-3 h-3 mr-2 items-center flex-row flex justify-center">
      {children}
    </div>
  )
}

export { MemberCardDiv, StatusDiv }