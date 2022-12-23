const MessageInput = () => {
  return (
    <div className="p-2 w-full h-48 items-center flex flex-row justify-center">
      <textarea className="h-full
                          w-full
                          p-4
                          rounded-md
                          border-2
                          border-white
                          bg-[#4b5563]"/>
      <button className="ml-2 w-1/12 h-full bg-[#7c3aed]">Send</button>
    </div>
  )
}

export default MessageInput