import { Task } from "@/core/Task";
import { IHandshakeRepository } from "@/core/IF/repository";
import { WebsocketJson } from "./model/LogJson";
import { debug } from "@/aop/logger";
import { fe2BeTask } from "../task/adapter";

export class HandshakeRepository implements IHandshakeRepository {
  baseURL: string;
  ws!: WebSocket;
  constructor(baseURL: string) {
    this.baseURL = baseURL;
    this.connect();
  }

  private connect() {
    this.ws = new WebSocket(`${this.baseURL}/ws`);
  }

  sendTaskLog(task: Task, operation: string) {
    const request :WebsocketJson = {
      messageType : 'log',
      value       : {
        valueType : 'task',
        created   : Date.now(),
        json      : { operation, task : fe2BeTask(task) },
      },
    };
    debug('websocket send', JSON.stringify(request));

    this.ws.send(JSON.stringify(request));
  }
}