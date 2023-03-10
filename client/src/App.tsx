import Chat from "./components/chat";
import SideBar from "./components/side-bar";
import { ChatContextProvider } from "./context/chat-context"

function App() {
  return (
    <ChatContextProvider>
      <Chat />
      <SideBar />
    </ChatContextProvider>
  );
}


export default App;