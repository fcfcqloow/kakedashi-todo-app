<script setup lang="ts">
import { getCurrentInstance, ref, onMounted, reactive } from 'vue';
import { ILogsRepository } from '@/core/IF/repository';
import LogsExpandList from '@/components/organisms/log/LogsExpandList.vue';
import MainTemplate from '@/components/template/MainTemplate.vue';
import { useAppStore } from '@/stores/app';

const app = getCurrentInstance();
const $logsRepository: ILogsRepository = app?.appContext.config.globalProperties?.$logsRepository;
const { subscribeGlobalLoadingPromise } = useAppStore();

const logsMap = reactive<Record<string, string[]|undefined>>({});
const dateList = ref<string[]>([]);

const logsJob = {
  syncDateList : (year?: string, month?: string) => subscribeGlobalLoadingPromise($logsRepository.listDate({ year, month }).then(response => dateList.value = response)),
  appendLogs   : (date: string) => {
    if (logsMap[date]) return;
    return $logsRepository.get(date).then((response) => logsMap[date] = response);
  },
};

onMounted(() => {
  logsJob.syncDateList();
});


</script>

<template>
  <main-template title="Logs Monitor">
    <logs-expand-list
      :date-list="dateList"
      :logsMap="logsMap"
      @open="logsJob.appendLogs"
      @update:date="logsJob.syncDateList($event.year, $event.month)"
    />
  </main-template>
</template>