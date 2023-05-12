export type SettingsJson = {
  taskLimit: number;
  sortable: boolean;
  topicColor: boolean;
  mode: 'dark'| 'white';
  topicListType: 'list' | 'icon';
  notificationIntervalSec: number;
  myColor: string;
  stopNotification: boolean;
};