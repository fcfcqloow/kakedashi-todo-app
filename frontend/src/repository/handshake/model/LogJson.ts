export type WebsocketJsonValue = {
  valueType: string;
  created: number;
  json: unknown;
}

export type WebsocketJson = {
  messageType: string;
  value: WebsocketJsonValue;
}