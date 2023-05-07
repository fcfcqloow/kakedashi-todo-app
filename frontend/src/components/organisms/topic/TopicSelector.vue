<script lang="ts" setup>
import { Topic } from '@/core/Topic';
import { computed } from 'vue';

type Props = {
  selected?: Topic[]; 
  topics?: Topic[];
  backgroundColor?: string;
  placeholder?: string;
};

type Emit = {
  (e: 'update:selected', selected: Topic[]): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();

const items = computed(() => props.topics?.filter(Boolean).map(tp => ({ ...tp, title : tp?.value })) || [])
const selected = computed(() => props.selected?.filter(Boolean).map(tp => ({ ...tp, title : tp?.value })) );
</script>

<template>
  <v-combobox
    clearable
    multiple
    :model-value="selected"
    :items="items"
    :bg-color="backgroundColor"
    @update:model-value="emits('update:selected', $event)"
    :placeholder="props.placeholder"
  >
    <template v-slot:selection="data">
      <v-chip
        v-bind="data.attrs"
        size="small"
        :key="data.id"
        :model-value="data.selected"
        :disabled="data.disabled"
        :color="data.item.value.color"
        @click:close="data.parent.selectItem(data.item)"
      >
      <p :style="`color: ${data.item.value.color}`">
        {{ data.item.value.value }}
      </p>
      </v-chip>
    </template>
  </v-combobox>
</template>

