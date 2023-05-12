<script lang="ts" setup>
type Props = {
  modelValue: string;
};
type Emit = {
  (e: 'update:modelValue', value: string): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const items = ['', '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12'];
</script>

<template>
  <v-autocomplete
    label="Month"
    :model-value="props.modelValue"
    :items="items"
    @update:model-value="emits('update:modelValue', $event)"
  />
</template>