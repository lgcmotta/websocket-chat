import {
  FC,
  createContext,
  useState,
  PropsWithChildren,
  Dispatch,
  SetStateAction,
  useContext,
  useMemo
} from "react";
import { IMember } from "../models/member";
import { IMessageReceived } from "../models/message";

interface IChatContextProps {
  state: IChatState;
  setState: Dispatch<SetStateAction<IChatState>>
}

interface IChatState {
  members: IMember[],
  myself: IMember,
  messages: IMessageReceived[]
  everyone: IMember
}

const everyone: IMember = { nickname: "everyone", connectionId: "" }

const initialState: IChatState = {
  members: [],
  myself: {
    connectionId: "",
    nickname: ""
  },
  messages: [],
  everyone: everyone
}


const ChatContext = createContext<IChatContextProps>({} as IChatContextProps)

const useChatContext = () => useContext(ChatContext)

function useStateSelector<T>(selector: (state: IChatState) => T) {
  const { state } = useChatContext();

  return selector(state);
}

const ChatContextProvider: FC<PropsWithChildren> = ({ children }) => {

  const [state
    , setState] = useState<IChatState>(initialState)

  return (
    useMemo(() => <ChatContext.Provider value={{ state, setState }} children={children} />, [state])
  )
}

export { ChatContext, ChatContextProvider, useChatContext, useStateSelector };