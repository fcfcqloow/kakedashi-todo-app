<script setup lang="ts">
import { ref, getCurrentInstance, onMounted, computed } from 'vue';

import TaskBoard from '@/components/organisms/task/TaskBoard.vue';
import TaskAddDialog from '@/components/organisms/task/TaskAddDialog.vue';
import TaskSortButton from '@/components/organisms/task/header/TaskSortButton.vue';
import TaskAddButton from '@/components/organisms/task/header/TaskAddButton.vue';
import TaskTopicFilterButton from '@/components/organisms/task/header/TaskTopicFilterButton.vue';
import TaskTitleSearchButton from '@/components/organisms/task/header/TaskTitleSearchButton.vue';
import MainTemplate from '@/components/template/MainTemplate.vue';

import { IHandshakeRepository, ITodoRepository, ITopicRepository } from '@/core/IF/repository';
import { Task } from '@/core/Task';
import { Tasks } from '@/core/Tasks';
import { Topic } from '@/core/Topic';
import { Color } from '@/core/Color';
import { useAppStore } from '@/stores/app';
import TaskRestoreButton from '@/components/organisms/task/header/TaskRestoreButton.vue';

const { subscribeGlobalLoadingPromise, showSnackbar, settings } = useAppStore();
const app = getCurrentInstance();
const $taskRepository: ITodoRepository = app?.appContext.config.globalProperties?.$taskRepository;
const $topicRepository: ITopicRepository = app?.appContext.config.globalProperties?.$topicRepository;
const $handshakeRepository: IHandshakeRepository = app?.appContext.config.globalProperties?.$handshakeRepository;

const isOpenAddDialog = ref(false);
const filterTopics = ref<Topic[]>([]);
const topics = ref<Array<Topic>>([]);
const tasks = ref<Tasks|undefined>();
const search = ref('');
const displayTasks = computed(() => {
  return tasks.value?.
    filter((tk) => filterTopics.value.length == 0 || filterTopics.value.some(tp => tk.topics.includes(tp.id || ''))).
    filter((tk) => !search || tk.value.includes(search.value));
});
const topicJob = {
  syncTopics : () => $topicRepository?.get().then((response: Topic[]) => topics.value = response),
};

const taskJob = {
  loadingWrap : (promise: Promise<any>, opt: { errorMsg: string }) =>
    subscribeGlobalLoadingPromise(
      // with errorhandling
      promise.catch(() => showSnackbar(opt.errorMsg, Color.ERROR)), 
      // show updated value
      taskJob.syncTasks,
    ),
  syncTasks   : () => $taskRepository?.get({ limit : settings.taskLimit }).then((response: Tasks) => tasks.value = response),
  sort        : (task: Task) =>
    taskJob.loadingWrap($taskRepository?.sort(), { errorMsg : 'failed to add task' }),
  add         : (task: Task) =>
    taskJob.loadingWrap($taskRepository?.add(task).then(() => $handshakeRepository.sendTaskLog(task, 'add')), { errorMsg : 'failed to add task' }),
  update      : (task: Task) => 
    taskJob.loadingWrap($taskRepository?.update(task).then(() => $handshakeRepository.sendTaskLog(task, 'update')), { errorMsg : 'failed to update task' }),
  remove      : (task: Task) => 
    taskJob.loadingWrap($taskRepository?.remove(task).then(() => $handshakeRepository.sendTaskLog(task, 'remove')), { errorMsg : 'failed to remove task' }),
  move        : (from: string, to: string, index: number, task: Task) =>
    taskJob.loadingWrap($taskRepository?.move(from, to, index, task).then(() => $handshakeRepository.sendTaskLog(task, to)), { errorMsg : 'failed to move task' }),
  restore     : () =>
    taskJob.loadingWrap($taskRepository?.restore(), { errorMsg : 'panic!!!!! restore!!!' }),
};

onMounted(() => {
  document.addEventListener('keydown', (e: KeyboardEvent) => {
    if (e.shiftKey && e.code === 'KeyN') {
      isOpenAddDialog.value = true;
      return;
    }
  });
  subscribeGlobalLoadingPromise(
    Promise.all([
      taskJob.syncTasks(),
      topicJob.syncTopics(),
    ])
  );
});

</script>

<template>
  <main-template title="TODO List">
    <template #header>
      <v-spacer />
      <task-restore-button @click="taskJob.restore" />
      <task-title-search-button v-model:search="search" />
      <task-topic-filter-button v-model:filter-topics="filterTopics" :topics="topics" />
      <task-sort-button @click="taskJob.sort" />
      <task-add-button @click="isOpenAddDialog = true" />
    </template>

    <task-board
      height="calc(100vh - 95px)"
      :model-value="displayTasks"
      :topics="topics"
      @get:task="taskJob.syncTasks"
      @add:task="taskJob.add"
      @remove:task="taskJob.remove"
      @move:task="taskJob.move"
      @update:task="taskJob.update"
    />
  </main-template>

  <task-add-dialog
    v-model="isOpenAddDialog"
    :topics="topics"
    :color-map="Topic.topicsToColorMap(topics)"
    @create:task="taskJob.add"
  />
</template>

<style>
.height100 {
  height: 100%;
}
v-select__selections {
  overflow: scroll;
  flex-wrap: nowrap;
}
.v-chip {
  overflow: initial;
}
.v-input__details {
  display: none;
}
</style>
