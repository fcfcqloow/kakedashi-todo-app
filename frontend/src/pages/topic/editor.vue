<script lang="ts" setup>
import { onMounted, getCurrentInstance, ref } from 'vue';
import { Color } from '@/core/Color';
import { Topic } from '@/core/Topic';
import { ITopicRepository } from '@/core/IF/repository';
import MainTemplate from '@/components/template/MainTemplate.vue';
import TopicEditor from '@/components/organisms/topic/TopicEditor.vue';
import { useAppStore } from '@/stores/app';

const { subscribeGlobalLoadingPromise, showSnackbar } = useAppStore();
const app = getCurrentInstance();
const $topicRepository: ITopicRepository = app?.appContext.config.globalProperties?.$topicRepository;

const topics = ref<Array<Topic>|undefined>(undefined);

const topicJob = {
  loadingWrap :(promise?: Promise<any>, errMsg?: string) => {
    subscribeGlobalLoadingPromise(promise?.catch(() => showSnackbar(errMsg ?? '', Color.ERROR)), topicJob.syncTopics);
  },
  syncTopics : () => $topicRepository?.get().then((response: Topic[]) => topics.value = response),
  add        : (topic: Topic) => topicJob.loadingWrap($topicRepository?.add(topic), 'failed to add topic'),
  remove     : (topic: Topic) => topicJob.loadingWrap($topicRepository?.remove(topic), 'failed to remove topic'),
  update     : (topic: Topic) => topicJob.loadingWrap($topicRepository?.update(topic), 'failed to update topic'),
};


onMounted(() => {
  subscribeGlobalLoadingPromise(topicJob.syncTopics());
});
</script>

<template>
  <main-template title="Topic Editor">
    <topic-editor
      :topics="topics ?? []"
      @create:topic="topicJob.add"
      @update:topic="topicJob.update"
      @remove:topic="topicJob.remove"
    />
  </main-template>
</template>

  <style scoped>
.border-radius {
  width: 50px;
  height: 50px;
  line-height: 50px;
  background-color: green;
  border-radius: 50%;
  color: #fff;
  text-align: center;
  display: inline-block; _display: inline;
  margin: 10px;
  cursor: pointer;
  overflow: hidden;
}
.border-radius p {
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}
</style>
