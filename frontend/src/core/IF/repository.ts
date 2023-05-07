import { Settings } from "../Settings";
import { Task } from "../Task";
import { Tasks } from "../Tasks";
import { Topic } from '../Topic';

export type ITodoRepository = {
  add    : (task: Task) => Promise<void>;
  remove : (task: Task) => Promise<void>;
  update : (task: Task) => Promise<void>;
  done   : (task: Task) => Promise<void>;
  get    : (options?: { limit?: number, offset?: number, target?: string }) =>  Promise<Tasks>;
  move   : (from: string, to: string, index: number, task: Task) => Promise<void>;
  sort   : () => Promise<void>;
};


export type ITopicRepository = {
  get    : () => Promise<Array<Topic>>;
  add    : (topic: Topic) => Promise<void>;
  remove : (topic: Topic) => Promise<void>;
  update : (topic: Topic) => Promise<void>;
};


export type ISettingsRepository = {
  get    : () => Promise<Settings>;
  update : (settings: Settings) => Promise<void>;
};

export type IHandshakeRepository = {
  sendTaskLog: (task: Task, operation: 'move'|'add'|'remove'|'update'|'done') => void;
};


export type ILogsRepository = {
  get      : (date: string) => Promise<string[]>;
  listDate : () => Promise<string[]>;
};

