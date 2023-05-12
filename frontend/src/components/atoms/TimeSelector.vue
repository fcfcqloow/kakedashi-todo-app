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
const value = computed(() => props.modelValue || '00:00');
</script>

<template>
  <input
    :value="value"
    @input="(event: any) => emits('update:modelValue', event.target.value)"
    type="time"
    required
  >
</template>