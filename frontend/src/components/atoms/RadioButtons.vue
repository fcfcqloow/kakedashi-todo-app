<script lang="ts" setup>
type Props = {
  modelValue: string|number,
  labels: (string|number)[],
  inline?: boolean,
  radioCustomStyle?: string;
};
type Emit = {
  (e: 'update:modelValue', label: string|number): void;
};

const props = defineProps<Props>();
const emits = defineEmits<Emit>();
</script>

<template>
  <fieldset>
    <div v-for="(label, idx) in labels" :key="label" :style="props.inline ? `display: inline-block; _display: inline;${props.radioCustomStyle}`: props.radioCustomStyle" >
      <input :id="`${label}`" type="radio" name="labels" :value="label" :checked="modelValue === label" @click="emits('update:modelValue', label)">
      <label :for="`${idx}`">{{ label }}</label>
    </div>
  </fieldset>
</template>

<style scoped>
p,
label {
  font: 1rem 'Fira Sans', sans-serif;
}

input {
  margin: 0.4rem;
}

</style>
