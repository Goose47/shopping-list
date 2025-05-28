<template>
  <button class="button" @click="$emit('click')">
    <slot></slot>
  </button>
</template>

<script setup>
import { color } from '@/theme/colors.js'
</script>

<style scoped lang="scss">
.button {
  width: 100%;
  padding: 12px;
  font-size: 16px;
  border: none;
  border-radius: 12px;
  background-color: v-bind('color("button_color")');
  color: v-bind('color("button_text_color")');
  cursor: pointer;
  transition: background-color 0.2s ease;

  &:hover {
    opacity: 0.9;
  }
}
</style>