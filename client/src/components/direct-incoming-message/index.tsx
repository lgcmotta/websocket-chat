import { format } from "date-fns"

const DirectIncomingMessage = () => {
  const text = "Hello!"
  const now = new Date()
  const time = format(now, "HH:mm:ss")
  const from = "George"
  return (
    <div className="w-full">
      <div className="m-2
                      ml-4
                      pt-2
                      pb-2
                      pl-4
                      pr-4
                      bg-[#4f46e5]
                      rounded-lg
                      w-1/3">
        <div className="flex flex-row">
          <div className="w-full items-center flex flex-row justify-start">
            <span className="right-0 text-xs">{from}
              <span className="pl-1 italic font-sans">@direct</span>
            </span>
          </div>
          <div className="w-full items-center flex flex-row justify-end">
            <span className="right-0 text-xs">{time}</span>
          </div>
        </div>
        <div className="pt-2">
          <span className="text-md whitespace-pre-wrap">
            {text}
          </span>
        </div>
      </div>
    </div>
  )
}

export default DirectIncomingMessage