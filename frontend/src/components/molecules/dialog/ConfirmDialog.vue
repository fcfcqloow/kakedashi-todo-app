<script lang="ts" setup>
type Props = {
  modelValue: boolean;
  title?: string;
};
type Emit = {
  (e: 'update:modelValue', modelValue: boolean): void;
  (e: 'cancel'): void;
  (e: 'ok'): void;
};


const props = defineProps<Props>();
const emits = defineEmits<Emit>();
</script>

<template>
  <v-dialog
      v-model="props.modelValue"
      width="auto"
    >
      <v-card>
        <v-card-text>{{ props.title }}</v-card-text>
        <v-card-actions>
          <v-spacer></v-spacer>
          <v-btn color="green-darken-1" variant="text" @click="emits('cancel')">
            cancel
          </v-btn>
          <v-btn color="green-darken-1" variant="text" @click="emits('ok')">
            ok
          </v-btn>
        </v-card-actions>
      </v-card>
    </v-dialog>
</template>