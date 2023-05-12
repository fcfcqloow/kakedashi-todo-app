<script lang="ts" setup>
import SearchTextField from '@/components/atoms/SearchTextField.vue';
import { useAppStore } from '@/stores/app';
import { ref } from 'vue';
import { Color } from '../../../../core/Color';

type Props = {
  search: string;
};
type Emit = {
  (e: 'update:search', search: string): void;
};

const { settings } = useAppStore();
const props = defineProps<Props>();
const emits = defineEmits<Emit>();
const showField = ref(false);
</script>
<template>
  <search-text-field
    v-if="showField"
    :search="props.search"
    @update:search="emits('update:search', $event)"
    :style="`
      width: 200px;
      background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
    `"
  />
  <v-tooltip
    text="Search Task Title"
    location="bottom"
  >
    <template v-slot:activator="{ props }">
      <v-btn
        v-bind="props"
        size="small"
        variant="outlined"
        :icon="showField ? 'mdi-close-circle' : 'mdi-text-search'"
        :color="showField ? 'red' : ''"
        :style="`
          margin: 10px;
          background-color: ${settings.mode === 'dark' ? 'black' : 'white'};
        `"
        @click="showField = !showField"
      />
    </template>
  </v-tooltip>
</template>