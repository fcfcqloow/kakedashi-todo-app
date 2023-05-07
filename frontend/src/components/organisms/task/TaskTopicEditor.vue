<script lang="ts" setup>
import { ref, reactive, computed, watch } from 'vue';
import { Task } from '@/core/Task';
import TopicSelector from '@/components/organisms/topic/TopicSelector.vue';
import { Topic } from '@/core/Topic';

type Props = {
  dialog: boolean;
  task: Task;
  topics: Topic[];
  topicMap: Record<string, Topic>;
};

type Emit = {
  (e: 'update:dialog', dialog: boolean): void;
  (e: 'submit', task: Task): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
const selected = ref<Topic[]>([]);
watch(() => props.task, () => {
  selected.value = props.task.topics.map(id => props.topicMap[id]);
});
</script>

<template>
   <v-dialog :model-value="props.dialog" @update:model-value="emits('update:dialog', $event)" overlay-color="transparent" width="400px">
      <v-card>
        <v-card-title><span class="text-h5">Task Topics</span></v-card-title>
        <v-card-subtitle>Task ID: {{ props.task.id }}</v-card-subtitle>
        <v-card-subtitle>Task: {{ props.task.value }}</v-card-subtitle>
        <v-card-text>
          <topic-selector
            :topics="props.topics"
            v-model:selected="selected"
          />
        </v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="emits('update:dialog', false);"
          >
            Close
          </v-btn>
          <v-btn
            color="blue-darken-1"
            variant="text"
            @click="emits('update:dialog', false); emits('submit', { ...task, topics : Topic.topicsToIds(selected) });"
          >
            Save
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>

<style>
.v-overlay__scrim {
  background: rgba(0,0,0, 0)
}

</style>