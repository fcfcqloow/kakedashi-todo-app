import { TaskRepository, TopicRepository } from '@/repository';
import { HandshakeRepository } from '@/repository/handshake/handshake';
import { LogsRepository } from '@/repository/logs/LogRepository';
import { SettingsRepository } from '@/repository/settings/SettingsRepository';
import type { App, Plugin } from 'vue';

export const RepositoryPlugin: Plugin = {
  install: (app: App, options?: any) => {
    app.config.globalProperties.$taskRepository = new TaskRepository("http://localhost:8080");
    app.config.globalProperties.$topicRepository = new TopicRepository("http://localhost:8080");
    app.config.globalProperties.$settingsRepository = new SettingsRepository("http://localhost:8080");
    app.config.globalProperties.$handshakeRepository = new HandshakeRepository("ws://localhost:8080");
    app.config.globalProperties.$logsRepository = new LogsRepository("http://localhost:8080");
  },
};

export default RepositoryPlugin;
