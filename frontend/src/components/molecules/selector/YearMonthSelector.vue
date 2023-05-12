<script lang="ts" setup>
import { ref } from 'vue';
import YearPicker from '@/components/atoms/YearPicker.vue';
import MonthPicker from '@/components/atoms/MonthPicker.vue';

type Props = {};
type Emit = {
  (e: 'update:date', date: { year: string, month: string } ): void;
};
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const year = ref('');
const month = ref('');
</script>

<template>
  <v-row style="margin: 10px;" align="center" justify="center">
    <v-col cols="2">
      <year-picker :model-value="year" @update:model-value="year = $event; emits('update:date', { year, month });" />
    </v-col>
    <v-col cols="2">
      <month-picker :model-value="month" @update:model-value="month = $event; emits('update:date', { year, month });" />
    </v-col>
  </v-row>
</template>