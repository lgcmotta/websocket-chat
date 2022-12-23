import { FC, createContext, useState, PropsWithChildren, Dispatch, SetStateAction, useContext } from "react";
import { IMember } from "../models/member";
import { IMessageReceived } from "../models/message";
import ChatWebSocketClient from "../api/ws";

interface IChatContextProps {
  state: IChatState;
  setState: Dispatch<SetStateAction<IChatState>>
}

interface IChatState {
  members: IMember[],
  myself: IMember | undefined,
  messages: IMessageReceived[]
}

const initialState: IChatState = {
  members: [],
  myself: undefined,
  messages: []
}

const ChatContext = createContext<IChatContextProps>({} as IChatContextProps)

const useChatContext = () => useContext(ChatContext)

const websocketClient = new ChatWebSocketClient()

function useSelector<T>(selector: (state: IChatState) => T) {
  const { state } = useChatContext();

  return selector(state);
}

const ChatContextProvider: FC<PropsWithChildren> = ({ children }) => {

  const [state, setState] = useState<IChatState>(initialState)

  websocketClient.onConnected(message => {
    console.log(message)
  })

  websocketClient.onMessageReceived(message => {
    setState(prev => {
      return { ...prev, messages: [...prev.messages, JSON.parse(message.data) as IMessageReceived] }
    })
  })

  return (
    <ChatContext.Provider value={{ state, setState }} children={children} />
  )
}

export { ChatContext, ChatContextProvider, useChatContext, useSelector };