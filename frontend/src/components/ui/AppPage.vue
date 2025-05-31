<template>
  <div class="page">
    <h1 class="page__title">{{ title }}</h1>
    <div class="page__main"><slot/></div>
    <template v-if="$slots.footer">
      <div class="page__footer">
        <slot name="footer" />
      </div>
    </template>
  </div>
</template>

<script setup>
import { color } from '@/theme/colors.js'

const props = defineProps({
  title: { required: true, type: String },
})
</script>

<style scoped lang="scss">
.page {
  padding: 16px;
  background-color: v-bind('color("bg_color")');
  color: v-bind('color("text_color")');
  min-height: 100vh;
  overflow-x: hidden;
  position: relative;

  &__title {
    font-size: 24px;
    margin-bottom: 16px;
    color: v-bind('color("section_header_text_color")');
  }

  &__main {
    padding-bottom: 36px;
  }

  &__footer {
    position: fixed;
    bottom: 0;
    left: 0;
    padding: 12px 16px;
    width:100%;
    height: 72px;
    background-color: v-bind('color("bg_color")');
    border-top: 1px solid rgba(0, 0, 0, 0.1);
    z-index: 100;
  }
}
</style>