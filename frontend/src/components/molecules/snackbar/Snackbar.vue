<script setup lang="ts">
type Props = {
  modelValue: boolean;
  timeout?: number;
  color?: string;
};
type Emit = {
  (e: 'update:modelValue', value: boolean): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>
<template>
  <v-snackbar
    :model-value="props.modelValue"
    :timeout="props.timeout"
    :color="props.color"
    @update:model-value="emits('update:modelValue', $event)"
  >
    <slot></slot>
  </v-snackbar>
</template>
