import { Settings } from "@/core/Settings";
import { SettingsJson } from "./model/SettingsJson";

export const toFeSettings = (settingJson: SettingsJson): Settings => new Settings({
  ...settingJson,
  topicColorFlag          : settingJson.topicColor,
  notificationIntervalSec : settingJson.notificationIntervalSec,
});

export const toBeSettings = (settings: Settings): SettingsJson => ({
  taskLimit               : settings.taskLimit,
  sortable                : settings.sortable,
  topicColor              : settings.topicColorFlag,
  mode                    : settings.mode,
  topicListType           : settings.topicListType,
  notificationIntervalSec : settings.notificationIntervalSec,
});