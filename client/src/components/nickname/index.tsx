const Nickname = () => {
  return (
    <div className="mt-2 w-full items-start flex-col flex justify-start">
      <p className="ml-4 mr-4">Logged as:</p>
      <input className="mr-4
                        ml-4
                        mt-2
                        p-2
                        w-11/12
                        items-center
                        flex-col
                        flex 
                        justify-start 
                        rounded-md 
                        border-2 
                        border-white 
                        text-left
                        bg-[#4b5563]"
        type="text" />
      <div className="pr-4 pl-4 pt-2 w-full items-center flex-col flex justify-center">
        <button className="w-full mt-2 bg-[#7c3aed] rounded-md border-2">Change Nickname</button>
      </div>
    </div>
  )
}

export default Nickname