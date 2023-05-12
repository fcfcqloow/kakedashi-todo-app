export class Task {
  public id: string;
  public value: string;
  public topics: string[];
  public deadLine?: Date;
  public priority: number;
  public doneDate?: Date;
  public created: Date;
  public constructor(args: {
    id: string;
    value: string;
    topics: string[];
    deadLine?: Date;
    priority?: number;
    doneDate?: Date;
    created: Date;
  }) {
    this.id = args.id;
    this.value = args.value;
    this.topics = args.topics;
    this.deadLine = args.deadLine;
    this.priority = args.priority || 0;
    this.doneDate = args.doneDate;
    this.created = args.created;
  }

  static createEmptyTask() {
    return new Task({
      id       : `${Date.now()}_${Math.random().toString(32).substring(2)}`,
      value    : '',
      deadLine : new Date(),
      created  : new Date(),
      topics   : [],
    });
  }

}