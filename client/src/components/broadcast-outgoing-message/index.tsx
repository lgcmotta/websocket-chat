import { format } from "date-fns"

const BroadcastOutgoingMessage = () => {
  const text = "Hello!"
  const now = new Date()
  const time = format(now, "HH:mm:ss")
  const from = "Me"
  return (
    <div className="w-full">
      <div className="m-2
                      mr-4
                      pt-2
                      pb-2
                      pl-4
                      pr-4
                      bg-[#4b5563]
                      rounded-lg
                      w-1/3
                      float-right">
        <div className="flex flex-row">
          <div className="w-full items-center flex flex-row justify-start">
            <span className="right-0 text-xs">{time}</span>
          </div>
          <div className="w-full items-center flex flex-row justify-end">
            <span className="right-0 text-xs">{from}
              <span className="pl-1 italic font-sans">@everyone</span>
            </span>
          </div>
        </div>
        <div className="pt-2 float-right">
          <span className="text-md whitespace-pre-wrap ">
            {text}
          </span>
        </div>
      </div>
    </div>
  )
}

export default BroadcastOutgoingMessage