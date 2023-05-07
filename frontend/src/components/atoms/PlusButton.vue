<script lang="ts" setup>
type Props = {
  size?: number|string;
  backgroundColor?: string;
};
type Emit = {
  (e: 'click', date: Date): void;
};


const props = defineProps<Props>();
const emits = defineEmits<Emit>();
</script>

<template>
  <v-btn
    :size="props.size ?? 'large'"
    variant="outlined"
    icon="mdi-plus"
    @click="emits('click', $event)"
    :style="`background-color: ${props.backgroundColor};`"
  />
</template>