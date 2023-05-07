import { Task } from "./Task";

type Todo = Task[];
type Processing = Task[];
type Done = Task[];
type Pending = Task[];

export interface ITasks {
  todo       : Todo,
  processing : Processing,
  done       : Done,
  pending    : Pending,
};
export const TasksKeys: Array<keyof ITasks>  = ['todo', 'processing', 'pending', 'done'];
export class Tasks implements ITasks {
  public todo       : Todo;
  public processing : Processing;
  public done       : Done;
  public pending    : Pending;

  public constructor(args: ITasks) {
    this.todo = args.todo;
    this.processing = args.processing;
    this.done = args.done;
    this.pending = args.pending;
  }

  filter(ho: (task: Task) => boolean): Tasks {
    const result = Tasks.createEmpty();
    result.todo = this.todo.filter(ho);
    result.processing = this.processing.filter(ho);
    result.pending = this.pending.filter(ho);
    result.done = this.done.filter(ho);
    return result;
  }

  static createEmpty(): Tasks {
    return new Tasks({
      done       : [],
      todo       : [],
      pending    : [],
      processing : [],
    });
  }
}