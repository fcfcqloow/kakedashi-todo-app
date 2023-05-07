<script lang="ts" setup>
import { onMounted, getCurrentInstance, ref, watch } from 'vue';
import { debug } from '@/aop/logger';
import { ITodoRepository } from '@/core/IF/repository';
import { useAppStore } from '@/stores/app';

const app = getCurrentInstance();
const $taskRepository: ITodoRepository = app?.appContext.config.globalProperties?.$taskRepository;

const { settings } = useAppStore();
const timeoutjob = ref(0);
const requestPermission = () => {
  if (!('Notification' in window)) {
    return;
  }

  debug('Notification.permission', Notification.permission);
  if (['denied', 'granted'].includes(Notification.permission)) {
    return;
  }
  
  return Notification.requestPermission();
};

const notificationInterval = (interval: number) => {
  debug('Next Notification Sec', 1000 * settings.notificationIntervalSec);
  timeoutjob.value = setTimeout(() => {
    $taskRepository.get({}).then((response) => {
      const now = Date.now();
      const timeover = response.todo.filter(task => task.deadLine && task.deadLine.getTime() <= now);
      if (timeover) {
        new Notification(`timeout ${timeover.length}\n` + timeover.map(task => task.value).join('\n'))
      }
    });
    notificationInterval(1000 * settings.notificationIntervalSec);
  }, interval) as any as number;
};

onMounted(async () => {
  await requestPermission();
  notificationInterval(1000 * settings.notificationIntervalSec);
});

watch(() => settings.notificationIntervalSec, () => {
  if (timeoutjob.value) {
    debug('clear timeout', timeoutjob.value);
    clearTimeout(timeoutjob.value);
  }
  notificationInterval(1000 * settings.notificationIntervalSec);
});

</script>

<template>

</template>