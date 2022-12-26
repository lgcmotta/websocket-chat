import { FC, PropsWithChildren, RefObject } from "react"

const MessagesContainer: FC<PropsWithChildren> = ({ children }) => {
  return (
    <div className="w-full
                    h-full
                    mr-4
                    rounded-md
                    border-2
                    border-white
                    items-center
                    flex
                    flex-col
                    justify-between 
                    justify-center">
      {children}
    </div>
  )
}

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

const ScrollToLastMessage: FC<{ scrollRef: RefObject<HTMLDivElement> }> = ({ scrollRef }) => {
  return (
    <div ref={scrollRef} />
  )
}

export { MessagesContainer, MessagesBox, ScrollToLastMessage }