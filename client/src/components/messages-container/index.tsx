import { FC, PropsWithChildren } from "react"

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

export default MessagesContainer