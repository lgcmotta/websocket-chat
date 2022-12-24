import ChatBox from "./components/chat-box";
import MembersList from "./components/members-list";
import { ChatContextProvider } from "./context/chat-context"

(window as any).global = window;

function App() {
  return (
    <ChatContextProvider>
      <ChatBox />
      <MembersList />
    </ChatContextProvider>
  );
}


export default App;