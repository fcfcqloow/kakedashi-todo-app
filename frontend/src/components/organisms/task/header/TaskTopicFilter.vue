<script lang="ts" setup>
import { useAppStore } from '@/stores/app';
import TopicSelector from '@/components/organisms/topic/TopicSelector.vue';
import { Topic } from '@/core/Topic';

type Props = {
  topics?: Topic[];
  filterTopics: Topic[];
};
type Emit = {
  (e: 'update:filterTopics', topics: Topic[]): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>
<template>
  <topic-selector
    :selected="props.filterTopics"
    @update:selected="emits('update:filterTopics', $event)"
    placeholder="Topic Filter"
    :topics="props.topics"
    :style="`
      background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
      width: 50px;
      height: 56px;
      overflow: scroll;
      margin: 10px;
    `"
  />
</template>