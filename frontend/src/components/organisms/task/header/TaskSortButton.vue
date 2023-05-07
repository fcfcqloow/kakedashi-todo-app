<script lang="ts" setup>
import { useAppStore } from '@/stores/app';
import SortButton from '@/components/atoms/SortButton.vue';

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
    v-if="settings.sortable"
    text="Sort Todo (Deadline + Priority)"
    location="bottom"
  >
    <template v-slot:activator="{ props }">
      <sort-button 
        v-bind="props"
        size="small"
        @click="emits('click')"
        :background-color="settings.mode === 'dark' ? 'black' : 'white'"
        :style="`margin: 10px;`"
      />
    </template>
  </v-tooltip>
</template>