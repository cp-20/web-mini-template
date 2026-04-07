<template>
  <button :class="classes" class="button" :type="type" @click="$emit('click')">
    <slot />
  </button>
</template>

<script setup lang="ts">
import { computed } from "vue";

const props = withDefaults(
  defineProps<{
    type?: "button" | "submit";
    variant?: "primary" | "secondary" | "danger";
  }>(),
  {
    type: "button",
    variant: "primary",
  },
);

defineEmits<{
  click: [];
}>();

const classes = computed(() => {
  return {
    secondary: props.variant === "secondary",
    danger: props.variant === "danger",
  };
});
</script>
