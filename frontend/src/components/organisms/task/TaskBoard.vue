<script setup lang="ts">
import { ref, computed } from 'vue';
import { debug } from '@/aop/logger';
import { Task } from '@/core/Task';
import { ITasks } from '@/core/Tasks';
import { Topic } from '@/core/Topic';
import { Tasks, TasksKeys } from '@/core/Tasks';
import TaskCard from '@/components/organisms/task/TaskCard.vue';
import { useAppStore } from '@/stores/app';
import TaskTopicEditor from './TaskTopicEditor.vue';

const { settings } = useAppStore();

type Props = {
  modelValue: Tasks;
  topics: Topic[];
  height?: number|string;
};
type Emit = {
  (e: 'update:modelValue', modelValue: Tasks): void;
  (e: 'move:task', from: typeof TasksKeys[number], to: typeof TasksKeys[number], index: number, task: Task): void;
  (e: 'update:task', task: Task): void;
  (e: 'add:task', task: Task): void;
  (e: 'remove:task', task: Task): void;
};

const props = withDefaults(defineProps<Props>(), {
  modelValue : () => new Tasks({
    todo       : [],
    processing : [],
    pending    : [],
    done       : [],
  }),
});
const emits = defineEmits<Emit>();

const isOpenTopicEditor = ref(false);
const topicColorMap = computed(() => Topic.topicsToColorMap(props.topics));
const topicNameMap = computed(() => Topic.topicsToNameMap(props.topics));
const topicMap = computed(() => Topic.topicsToMap(props.topics));
const editingTopic = ref<Task|undefined>();
const pickingTask = ref<{ task: Task; index: number; key: keyof ITasks; element: any }|undefined>();
const overElement = ref<any>();
const overLineElement = ref<any>();
const findTargetElement = (target: any, targetClass: string): any => {
  if (!target.parentElement) {
    return undefined;
  }

  if ((target.classList as DOMTokenList)?.contains(targetClass)) {
    return target;
  }


  return findTargetElement(target.parentElement, targetClass);
};

const startDrag = (key: keyof ITasks, index: number, task: Task, event: DragEvent) => {
  debug('startDrag', key, task.value);
  pickingTask.value = { task, index, key, element :  event.target };
  pickingTask.value.element.classList.add('active');
};

const endDrag = (key: keyof ITasks, index: number, task: Task, event: DragEvent) => {
  debug('endDrag', key, task.value);
  (event.target as any).classList.remove('active');
  pickingTask.value = undefined;
  overElement.value?.classList?.remove('over');
  overElement.value = undefined;
  overLineElement.value?.classList?.remove('line-over');
  overLineElement.value = undefined;
};

const enterDrag = (event: DragEvent) => {
  debug('enterDrag');
  const foundElement = findTargetElement(event.target, 'drag');

  if (overElement.value?.classList && overElement.value.id !== foundElement.id) {
    overElement.value.classList.remove('over');
    overElement.value = undefined;
  }

  if (!foundElement.classList.contains('over')) {
    overElement.value = foundElement;
    overElement.value.classList.add('over');
  }

};

const leaveDrag = (event: DragEvent) => {
};

const drop = (key: keyof ITasks, index: number, task: Task) => {
  debug('drop', task.value);
  if (pickingTask.value) {
    emits('move:task', pickingTask.value.key, key, index, pickingTask.value.task);
    pickingTask.value = undefined;
  }
};

const dropLine = (key: keyof ITasks) => {
  debug('dropLine');
  if (pickingTask.value) {
    emits('move:task', pickingTask.value.key, key, props.modelValue[key].length - 1, pickingTask.value.task);
    pickingTask.value = undefined;
  }
};

const overLine = (event: any) => {
  debug('overLine');
  if (event.target?.getBoundingClientRect().y < overElement.value?.getBoundingClientRect().y) {
    overElement.value.classList.remove('over');
    overElement.value = undefined;
  }

  const foundElement = findTargetElement(event.target, 'line');
  if (overLineElement.value) {
    overLineElement.value.classList.remove('line-over');
    overLineElement.value = undefined;
  }

  if (foundElement.classList.contains('line')) {
    foundElement.classList.add('line-over');
    overLineElement.value = foundElement;
  }
};

</script>

<template>
  <v-container fluid>
    <v-row>
      <v-col v-for="key in TasksKeys" :key="key" cols="3">
        <div
          :id="key"
          @dragover.prevent
          @drop="dropLine(key)"
          @dragover="overLine"
          draggable="false"
          class="line"
        >
          <v-card
            class="mx-auto pa-2 transparent"
            style="overflow: scroll;"
            min-height="100px"
            :height="props.height"
          >
            <task-topic-editor
              v-if="editingTopic"
              v-model:dialog="isOpenTopicEditor"
              :task="editingTopic"
              :topics="topics ?? []"
              :topic-map="topicMap ?? {}"
              @submit="emits('update:task', $event)"
            />
            <v-list style="overflow: visible;" >
              <v-sheet border rounded>
                <v-row justify="space-between" style="width: 100%;">
                  <v-col cols="7">
                    <v-list-subheader>
                      <span style="font-size: 18px; line-height: 20px;">{{ key.toUpperCase() ?? '' }}</span>
                    </v-list-subheader>
                  </v-col>
                </v-row>
              </v-sheet>
              <v-sheet border rounded style="overflow: visible; padding: 10px;" min-height="100px">
                <div v-for="(item, index) in props.modelValue[key]"
                  :id="item.value"
                  :key="item.id"
                  class="drag"
                  :draggable="!Boolean((item as any).ignoreDrag)"
                  @dragstart="($event: DragEvent) => startDrag(key, index, item, $event)"
                  @dragend="($event: DragEvent) => endDrag(key, index, item, $event)"
                  @dragenter="enterDrag"
                  @dragleave="leaveDrag"
                  @dragover.prevent
                  @drop.prevent="() => drop(key, index, item)"
                >
                  <task-card
                    style="margin: 10px;"
                    :is-deletable="true"
                    :key="item.id"
                    :model-value="item"
                    :color="(settings.topicColorFlag && topicColorMap && item.topics?.length && topicColorMap[item.topics[0]]) || undefined"
                    @delete="emits('remove:task', item)"
                    @apply="emits('update:task', $event)"
                  >
                    <template #bottom>
                      <p style="text-align: center; padding-bottom: 5px; cursor: pointer;" @click="isOpenTopicEditor = true; editingTopic = item">
                        {{  topicNameMap && item.topics?.map((topicId) => topicNameMap && topicNameMap[topicId]).join(' / ') || 'No Topics' }}
                      </p>
                    </template>
                  </task-card>
                </div>
              </v-sheet>
            </v-list>
          </v-card>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<style scoped>
.drag {
  cursor: move;
}
.drag.active {
  visibility: hidden;
  transition: visibility linear 0.1s;
}
.drag.over {
  border-top: thick solid purple;
}
.line {}
.line-over {
  border: 1px solid rgb(0, 81, 128);
}
</style>
