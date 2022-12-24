import { ChatContextProvider } from "./context/chat-context"
import Chat from "./page";

function App() {
  return (
    <ChatContextProvider>
      <Chat />
    </ChatContextProvider>
  );
}


export default App;