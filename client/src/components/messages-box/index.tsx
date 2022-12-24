import { FC, PropsWithChildren } from "react"

const MessagesBox: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="h-full
                    w-full
                    pt-4
                    flex
                    flex-col
                    items-center
                    rounded-md
                    border-b-2
                    border-white
                    bg-[#18181b]
                    overflow-x-hidden
                    overflow-y-auto">
      {children}
    </div>
  )
}

export default MessagesBox;