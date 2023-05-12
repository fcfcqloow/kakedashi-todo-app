<script lang="ts" setup>
import { ref, watch, reactive } from 'vue';
import { Task } from '@/core/Task';
import { useAppStore } from '@/stores/app';
import { getFontColor } from '@/core/Color';

type Props = {
  modelValue: Task;
  isCreatePanel?: boolean;
  isDeletable?: boolean;
  color?: string;
  textAreaRows?: number;
};

type Emit = {
  (e: 'delete'): void;
  (e: 'apply', task: Task): void;
  (e: 'update:modelValue', task: Task): void;
};

const colorMap = [
  '#655',
  '#855',
  '#A55',
  '#C55',
  '#F55',
];

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
const { showCalender } = useAppStore();

const isEdited = ref(false);
const beforeValue = ref('');
const task = reactive(props.modelValue);

const clickDate = () => {
  const callback = (date: Date) => {
    task.deadLine = date;
    emits('apply', task);
  };
  showCalender(task.deadLine, callback);
};
// 編集内容を管理するWatch
watch(() => task.value, (newValue, oldValue) => {
  // 変更前と同じ内容なら編集中ではない
  if (newValue === beforeValue.value) {
    isEdited.value = false;
    beforeValue.value = '';
    return;
  }

  // もし、編集中でない場合に新しい場合が来たら編集中にする
  if (!isEdited.value) {
    beforeValue.value = oldValue;
    isEdited.value = true;
  }
});
</script>

<template>
  <v-card
    :color="props.color"
    elevation="3"
    style="overflow: visible; font-size: 1vh; position: relative; z-index 5;"
  >
    <v-btn
      v-if="props.isDeletable"
      size="compact"
      color="error"
      icon="mdi-alpha-x-circle-outline"
      style="position: absolute; right: -10px; top: -10px;"
      @click="emits('delete')"
    />

    <v-badge
      class="priority"
      :color="colorMap[Number(task.priority)]"
      :content="task.priority"
      inline
    />

    <div @click="clickDate" 
      :style="`
        color: ${props.color && getFontColor(props.color)};
        font-size: 15px;
      `"
    >
      {{ props.modelValue.deadLine?.toLocaleString() }}
    </div>

    <v-textarea
      :model-value="task.value"
      :style="`overflow: scroll; color: ${props.color && getFontColor(props.color)}` "
      auto-grow
      :rows="`${props.textAreaRows ?? 0}`"
      @update:model-value="task.value = $event"
    />

    <v-row justify="space-between" v-if="isEdited && !props.isCreatePanel">
      <v-col>
        <v-btn
          style="left: 5px; bottom: 5px;"
          icon="mdi-close"
          color="red-accent-3"
          @click="task.value = beforeValue"
        />
      </v-col>
      <v-col>
        <v-btn
          position="absolute"
          style="right: 5px; bottom: 5px;"
          color="light-green-accent-3"
          icon="mdi-check-circle-outline"
          @click="() =>{
            emits('apply', task);
            isEdited = false;
            beforeValue = '';
          }"
        />
      </v-col>
    </v-row>
    <slot name="bottom" />
  </v-card>
</template>


<style>
.date-picker {
  position: absolute;
  overflow: visible;
}
.action-row {
  display: flex;
  flex-direction: column;
  align-items: center;
  width: 100%;
}
.priority {
  position: absolute;
  right: 5px;
  top: 2px;
  font-size: 15px;
}
</style>
