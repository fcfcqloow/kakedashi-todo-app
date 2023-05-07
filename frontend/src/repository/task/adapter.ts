import { Task } from "@/core/Task";
import { Topic } from "@/core/Topic";
import { TaskJson } from "./model/TaskJson";
import { TopicJson } from "./model/TopicJson";

export const fe2BeTask = (task: Task): TaskJson => {
    return {
      ...task,
      deadLine : task.deadLine?.getTime() || 0,
      done     : task.doneDate?.getTime() || 0,
      created  : task.created?.getTime() || 0,
    }
  };
  
export const fe2BeTopic = (topic: Topic): TopicJson => {
    return {
      id      : topic.id || '',
      value   : topic.value,
      color   : topic.color,
      created : topic.created.getTime(),
      done    : topic.done?.getTime(),
    };
};
  
export const be2FeTask = (task: TaskJson): Task => {
    return {
        ...task,
        created  : task.created && new Date(task.created) || new Date(),
        deadLine : task.deadLine && new Date(task.deadLine) || undefined,
        doneDate : task.done && new Date(task.done) || undefined,
        priority : task.priority ?? 0,
    };
};