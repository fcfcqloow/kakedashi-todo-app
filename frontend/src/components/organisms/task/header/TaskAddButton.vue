<script lang="ts" setup>
import PlusButton from '@/components/atoms/PlusButton.vue';
import { useAppStore } from '@/stores/app';

type Props = {
};
type Emit = {
  (e: 'click'): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();

</script>
<template>
  <v-tooltip
    text="Shift + N"
    location="bottom"
  >
    <template v-slot:activator="{ props }">
      <plus-button
        v-bind="props"
        size="small"
        @click="emits('click')"
        :background-color="settings.mode === 'dark' ? 'black' : 'white'"
        :style="`margin: 10px;`"
      />
    </template>
  </v-tooltip>
</template>