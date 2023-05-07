<script lang="ts" setup>
import { ref } from 'vue';
import { Task } from '@/core/Task';
import { Topic } from '@/core/Topic';
import { Color } from '@/core/Color';
import RadioButtons from '@/components/atoms/RadioButtons.vue';
import TaskCard from '@/components/organisms/task/TaskCard.vue';
import TopicSelector from '@/components/organisms/topic/TopicSelector.vue';
import { copyObjectKV } from '@/util/object';

type Props = {
  modelValue: boolean;
  topics?: Topic[];
  colorMap?: Record<string, Color>;
};
type Emit = {
  (e: 'update:modelValue', isOpen: boolean): void;
  (e: 'create:task', task: Task): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const newTopics = ref<Topic[]>([]);
const newTask = ref(Task.createEmptyTask());
const shortcut = (event: KeyboardEvent) => {
  if (event.shiftKey) {
    switch(event.code) {
      case 'ArrowRight':
        newTask.value.priority = Math.min((newTask.value.priority ?? 0) + 1, 4);
        return;
      case 'ArrowLeft':
        newTask.value.priority = Math.max((newTask.value.priority ?? 0) - 1, 0);
        return;
      case 'Enter':
        if (newTopics?.value?.length) {
          newTask.value.topics = newTopics.value.map((tp) => tp.id).filter((id): id is string => Boolean(id));
          // newTopics.value = [];
        }
        emits('create:task', newTask.value);
        copyObjectKV(newTask.value, Task.createEmptyTask());
      
        return;
    }
  }
};
</script>

<template>
  <v-dialog
    :model-value="props.modelValue"
    width="auto"
    @update:model-value="emits('update:modelValue', $event)"
    @keydown="shortcut"
  >
    <task-card
      v-model="newTask"
      style="width: 50vw; height: 30vh;"
      is-create-panel
    >
      <template #bottom>
        <radio-buttons
          :model-value="newTask.priority ?? 0"
          @update:model-value="newTask.priority = $event"
          inline
          radio-custom-style="padding-left: 9px"
          :labels="[0, 1, 2, 3, 4]"
        />
        <topic-selector
          v-model:selected="newTopics"
          :topics="props.topics"
        />
      </template>
    </task-card>
  </v-dialog>
</template>
