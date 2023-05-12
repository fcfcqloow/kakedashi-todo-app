export class Settings {
  public taskLimit: number;
  public sortable: boolean;
  public topicColorFlag: boolean;
  public mode: 'dark'| 'white';
  public topicListType: 'list'|'icon';
  public notificationIntervalSec: number;
  public stopNotification: boolean;
  public myColor: string;
  constructor(args: {
    taskLimit?: number;
    sortable: boolean;
    topicColorFlag: boolean;
    mode: 'dark'| 'white';
    topicListType: 'list'|'icon';
    notificationIntervalSec: number;
    myColor: string;
    stopNotification: boolean;
  }) {
    this.taskLimit = args.taskLimit ?? 1_000;
    this.sortable = args.sortable;
    this.topicColorFlag = args.topicColorFlag;
    this.mode = args.mode || 'dark';
    this.topicListType = args.topicListType;
    this.notificationIntervalSec = Math.max(args.notificationIntervalSec, 30);
    this.myColor = args.myColor || '#E0E000';
    this.stopNotification = args.stopNotification || false;
  }
};
