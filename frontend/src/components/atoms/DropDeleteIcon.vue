1<script lang="ts" setup>
import { ref } from 'vue';

type Props = {
  customStyle?: string;
  size?: number;
  disabled?: boolean;
};
type Emit = {
  (e: 'drop', event: any): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const isOver = ref(false);

</script>

<template>
  <div
    v-if="!props.disabled"
    class="wrapper"
    :style="props.customStyle"
  >
    <v-icon
      color="red"
      :icon="isOver ?  'mdi-delete-empty' : 'mdi-delete'"
      :size="props.size ?? 50"
      @dragenter="isOver = true;"
      @dragleave="isOver = false;"
      @dragover.prevent
      @drop="emits('drop', $event)"
    />
  </div>
</template>

<style scoped>
.wrapper {}
</style>
