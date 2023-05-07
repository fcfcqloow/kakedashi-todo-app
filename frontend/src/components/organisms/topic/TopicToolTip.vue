<script lang="ts" setup>
import { ref } from 'vue';
import type { Topic } from '@/core/Topic';

type Props = {
  topic: Topic;
};
type Emit = {
  (e: 'dbclick:topic'): void;
  (e: 'start:drag', topic: Topic): void;
  (e: 'end:drag', topic: Topic): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const isOpen = ref(false);
</script>

<template>
  <div
    v-bind="props"
    class="border-radius"
    draggable="true"
    :style="`background-color: ${props.topic?.color};`"
    @dblclick="emits('dbclick:topic')"
    @dragstart="emits('start:drag', props.topic)"
    @dragend="emits('end:drag', props.topic)"
    >
      <p>{{ props.topic?.value[0] }}</p>
    </div>
</template>

<style scoped>
.border-radius {
  width: 50px;
  height: 50px;
  line-height: 50px;
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
