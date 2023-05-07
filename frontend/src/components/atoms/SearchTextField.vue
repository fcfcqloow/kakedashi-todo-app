<script lang="ts" setup>

type Props = {
  search: string
};
type Emit = {
  (e: 'update:search', search: string): void;
};


const props = defineProps<Props>();
const emits = defineEmits<Emit>();
</script>

<template>
  <v-text-field
    :model-value="props.search"
    @update:model-value="emits('update:search', $event)"
    label="🔍 Search"
    single-line
    hide-details
  />
</template>