import axios from 'axios';

import type { ITodoRepository, ITopicRepository } from "@/core/IF/repository";
import type { Task } from "@/core/Task";
import { Tasks } from "@/core/Tasks";
import { Topic } from "@/core/Topic";
import type { TaskAddRequest } from './model/TaskAddRequest';
import type { TasksGetResponse } from './model/TasksGetResponse';
import type { TaskMoveRequest } from './model/TaskMoveRequest';
import type { TaskPatchRequest } from './model/TaskPatchRequest';
import type { TaskDeleteRequest } from './model/TaskDeleteRequest';
import type { TopicAddRequest } from './model/TopicAddRequest';
import type { TopicJson } from './model/TopicJson';
import type { TopicUpdateRequest } from './model/TopicUpdateRequest';
import type { TopicRemoveRequest } from './model/TopicRemoveRequest';
import { be2FeTask, fe2BeTask, fe2BeTopic } from './adapter';

export class TaskRepository implements ITodoRepository  {
  private baseURL: string;
  constructor(baseURL: string) {
    this.baseURL = baseURL;  
  }

  async get(options?: { limit?: number, offset?: number, target?: string }): Promise<Tasks> {
    const axiosResponse = await axios.get(`${this.baseURL}/todo`, {
      params : {
        limit: options?.limit,
        offset: options?.offset,
        target: options?.target,
        // sort : true,
      },
    });
    const response = axiosResponse.data as TasksGetResponse;
    return Promise.resolve(new Tasks({
      todo       : response.todo.map(be2FeTask),
      processing : response.processing.map(be2FeTask),
      done       : response.done.map(be2FeTask),
      pending    : response.pending.map(be2FeTask),
    }));
  }

  async sort() {
    await axios.patch(`${this.baseURL}/todo/sort`)
  }
  async update(task: Task): Promise<void> {
    await axios.patch(`${this.baseURL}/todo`, <TaskPatchRequest>fe2BeTask(task))
  }
  async add(task: Task): Promise<void> {
    await axios.put(`${this.baseURL}/todo`, <TaskAddRequest>fe2BeTask(task))
  }
  async remove(task: Task): Promise<void> {
    await axios.delete(`${this.baseURL}/todo`, { data : <TaskDeleteRequest>fe2BeTask(task) })
  }
  async done(task: Task): Promise<void> { 
    throw new Error('no implement') 
  }
  async move(from: string, to: string, index: number, task: Task): Promise<void> {
    await axios.post(`${this.baseURL}/todo`, <TaskMoveRequest>{
      from,
      to,
      index,
      task : fe2BeTask(task),
    });
  }
  async restore() {
    await axios.post(`${this.baseURL}/todo/restore/done`);
  }
};

export class TopicRepository implements ITopicRepository  {
  private baseURL: string;
  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  async get() {
    const axiosResponse = await axios.get(`${this.baseURL}/topic`);
    const topicJson = axiosResponse.data as Array<TopicJson>;
    const toFeTopic = (beTopic: TopicJson): Topic => new Topic({
      id      : beTopic.id,
      value   : beTopic.value,
      color   : beTopic.color,
      created : new Date(beTopic.created),
      done    : (beTopic.done ?? -1) > 0 ?  new Date(beTopic.done!) : undefined,

    });
    return Promise.resolve(topicJson.map(toFeTopic));
  }

  async add(topic: Topic) {
   await axios.put(`${this.baseURL}/topic`, <TopicAddRequest>{ ...fe2BeTopic(topic), id: Math.random().toString(32).substring(2) + Date.now() })
  }
  async remove(topic: Topic) {
    await axios.delete(`${this.baseURL}/topic`, { data : <TopicRemoveRequest>fe2BeTopic(topic) });
  }
  async update(topic: Topic) { 
    await axios.patch(`${this.baseURL}/topic`, <TopicUpdateRequest>fe2BeTopic(topic))
  }
};