import { IBroadcastMessage, IDirectMessage, IJoinMessage, IMessage } from "../models/message"
import  WebSocketClient, { MessageHandler } from "@mitz-it/websocket-client"

export default class ChatWebSocketClient {
  private client: WebSocketClient

  connection: string
  connected: boolean = false
  constructor() {
    const connection = import.meta.env.VITE_WEBSOCKET_URL || "";

    if (connection == "") {
      throw new Error("websocket url was not provided")
    }
    this.connection = connection
    this.client = new WebSocketClient(connection)
  }

  onMessageReceived<TMessage>(handler: MessageHandler<TMessage>) {
    if (this.client == undefined) return

    this.client.addMessageHandler(handler)
  }

  connect() {
    this.client.connect(() => {
      window.alert("connected!")
    })
  }

  isConnected(): boolean {
    return this.client.isConnected()
  }

  join(message: IJoinMessage) {
    this.publish<IJoinMessage>(message)
  }

  broadcast(message: IBroadcastMessage) {
    this.publish<IBroadcastMessage>(message)
  }

  direct(message: IDirectMessage) {
    this.publish<IDirectMessage>(message)
  }

  requestMembersList() {
    if (this.client == undefined) return
    this.client.publish({ action: "members" })
  }

  private publish<TMessage>(message: TMessage) {
    if (this.client == undefined) return
    this.client.publish<TMessage>(message)
  }
}

export const websocketClient: ChatWebSocketClient = new ChatWebSocketClient();