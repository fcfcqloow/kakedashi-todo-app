<script lang="ts" setup>
import LogsPanel from './LogsPanel.vue';
import LogsList from './LogsList.vue';
import { reactive } from 'vue';

type Props = {
  dateList: string[];
  logsMap: Record<string, string[]|undefined>;
};
type Emit = {
  (e: 'open', date: string): void;
};
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>

<template>
  <v-expansion-panels>
    <logs-panel v-for="item in props.dateList" :date="item">
      <logs-list
        :logs="props.logsMap[item] || []"
        @on:mount="emits('open', item)"
      />
    </logs-panel>
  </v-expansion-panels>
</template>