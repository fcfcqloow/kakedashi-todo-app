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
    @click="emits('click', $event)"
    :style="`background-color: ${props.backgroundColor};`"
    icon="mdi-restore"
  />
</template>