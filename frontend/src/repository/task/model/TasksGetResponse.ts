import { TaskJson } from "./TaskJson";

export interface TasksGetResponse {
    todo: TaskJson[];
    done: TaskJson[];
    pending: TaskJson[];
    processing: TaskJson[];
}