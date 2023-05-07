<script lang="ts" setup>
import { ref } from 'vue';
import { Topic } from '@/core/Topic';
import DropDeleteIcon from '@/components/atoms/DropDeleteIcon.vue';
import TopicToolTip from '@/components/organisms/topic/TopicToolTip.vue';

type Props = {
  topics: Array<Topic>;
  showDeleteIcon: boolean;
  height: string;
  width: string;
};
type Emit = {
  (e: 'update:showDeleteIcon', showDeleteIcon: boolean): void;
  (e: 'start:drag', topic: Topic): void;
  (e: 'end:drag'): void;
  (e: 'delete', topic: Topic): void;
};
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const draggingTopic = ref<Topic|undefined>();

</script>

<template>
  <v-sheet color="rgba(255, 255, 255, 0.1)" :height="props.height" :width="props.width" style="overflow: scroll;">
    <drop-delete-icon
      :disabled="!props.showDeleteIcon"
      :size="80"
      custom-style="position: absolute; top: 8%; left: 50%;"
      @drop="draggingTopic && emits('delete', draggingTopic); draggingTopic = undefined"
    />
    <topic-tool-tip
      v-for="item in props.topics"
      :key="item.id"
      :topic="item"
      @start:drag="emits('update:showDeleteIcon', true); draggingTopic = $event"
      @end:drag="emits('update:showDeleteIcon', false); draggingTopic = undefined"
    />
  </v-sheet>
</template>