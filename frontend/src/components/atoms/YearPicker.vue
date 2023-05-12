<script lang="ts" setup>
import { computed } from 'vue';

type Props = {
  modelValue: string;
};
type Emit = {
  (e: 'update:modelValue', value: string): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const items = computed(() => {
  const result: Array<string> = [''];
  for (let i = 0; i < 50; i++) {
    result.push(`${2023 + i}`);
  }

  return result;
});
</script>

<template>
  <v-autocomplete
    label="Years"
    :model-value="props.modelValue"
    :items="items"
    @update:model-value="emits('update:modelValue', $event)"
  />
</template>