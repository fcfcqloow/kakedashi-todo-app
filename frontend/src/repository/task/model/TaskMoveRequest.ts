import { TaskJson } from "./TaskJson";

export interface TaskMoveRequest {
   from: string;
   to: string;
   index: number;
   task: TaskJson;
}