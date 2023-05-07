<script lang="ts" setup>

type Props = {
  date: string;
};
type Emit = {
  (e: 'click', date: Date): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
const log = console.log
</script>
<template>
  <v-expansion-panel>
    <v-expansion-panel-title @click="log('click')">
      <template v-slot:default="{ expanded }">
        {{ props.date }}
      </template>
    </v-expansion-panel-title>
    <v-expansion-panel-text>
      <slot />
    </v-expansion-panel-text>
  </v-expansion-panel>
</template>