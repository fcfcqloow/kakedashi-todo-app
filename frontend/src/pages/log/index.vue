<script setup lang="ts">
import { getCurrentInstance, ref, onMounted, reactive } from 'vue';
import { ILogsRepository } from '@/core/IF/repository';
import LogsExpandList from '@/components/organisms/log/LogsExpandList.vue';
import MainTemplate from '@/components/template/MainTemplate.vue';

const app = getCurrentInstance();
const $logsRepository: ILogsRepository = app?.appContext.config.globalProperties?.$logsRepository;

const logsMap = reactive<Record<string, string[]|undefined>>({});
const dateList = ref<string[]>([]);
onMounted(() => {
  $logsRepository.listDate().then(response => dateList.value = response);
});

const setLogs = (date: string) => {
  if (logsMap[date]) return;
  return $logsRepository.get(date).then((response) => logsMap[date] = response);
};
</script>

<template>
  <main-template title="Logs Monitor">
    <logs-expand-list
      :date-list="dateList"
      :logsMap="logsMap"
      @open="setLogs"
    />
  </main-template>
</template>