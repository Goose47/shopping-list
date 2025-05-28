<template>
  <div class="page">
    <h1 class="title">{{ title }}</h1>
    <slot/>
  </div>
</template>

<script setup>
import { color } from '@/theme/colors.js'

const props = defineProps({
  title: { required: true, type: String },
})
</script>

<style scoped lang="scss">
.title {
  font-size: 24px;
  margin-bottom: 16px;
  color: v-bind('color("section_header_text_color")');
}
.page {
  padding: 16px;
  background-color: v-bind('color("bg_color")');
  color: v-bind('color("text_color")');
  min-height: 100vh;
}
</style>