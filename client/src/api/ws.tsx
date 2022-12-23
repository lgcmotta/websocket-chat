export default class ChatWebSocketClient {
  socket: WebSocket
  connection: string
  constructor() {
    this.connection = import.meta.env.VITE_WEBSOCKET_URL || ""

    if (this.connection == "") {
      throw new Error("websocket url was not provided")
    }

    this.socket = new WebSocket(this.connection)
  }

  onMessageReceived(callback: (message: any) => void) {
    this.socket.onmessage = callback;

  }

  onConnected(callback: (message: any) => void) {
    window.alert("connected!")
    this.socket.onopen = callback;
  }
}
