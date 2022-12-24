import { IBroadcastMessage, IDirectMessage, IJoinMessage, IMessage } from "../models/message"

export default class ChatWebSocketClient {
  socket: WebSocket | undefined
  connection: string
  connected: boolean = false
  constructor() {
    this.connection = import.meta.env.VITE_WEBSOCKET_URL || ""

    if (this.connection == "") {
      throw new Error("websocket url was not provided")
    }
  }

  onMessageReceived(callback: (message: any) => void) {
    if (this.socket == undefined) return

    this.socket.onmessage = callback;
  }

  connect() {
    this.socket = new WebSocket(this.connection)
    this.socket.onopen = (e) => {
      this.connected = true
      window.alert("connected!")
    }
  }

  isConnected(): boolean {
    return this.connected
  }

  join(message: IJoinMessage) {
    const join = JSON.stringify(message)
    this.publish(join)
  }

  broadcast(message: IBroadcastMessage) {
    const broadcast = JSON.stringify(message)
    this.publish(broadcast)
  }

  direct(message: IDirectMessage) {
    const direct = JSON.stringify(message)
    this.publish(direct)
  }

  private publish(message: string) {
    if (this.socket == undefined) return
    this.socket.send(message)
  }
}

export const websocketClient: ChatWebSocketClient = new ChatWebSocketClient();