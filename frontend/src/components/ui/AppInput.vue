<template>
  <div class="text-input-wrapper">
    <input
        class="text-input"
        v-model="model"
        :placeholder="placeholder"
        ref="input"
    />
  </div>
</template>

<script setup>
import {ref} from 'vue'
import {color} from '@/theme/colors.js'

const model = defineModel() // заменяет v-model / props+emit

const props = defineProps({
  placeholder: {type: String, default: 'Введите текст'}
})

const input = ref(null)
const focus = () => {
  input.value.focus()
}
defineExpose({ focus })
</script>

<style scoped lang="scss">
.text-input-wrapper {
  width: 100%;
  display: flex;
  align-items: center;
  margin-bottom: 10px;
  padding: 0 12px;
  height: 48px;
  border-radius: 12px;
  background-color: v-bind('color("secondary_bg_color")');
  transition: background-color 0.25s ease;
}

.text-input {
  flex: 1;
  width: 100%;
  border: none;
  background: transparent;
  outline: none;
  font-size: 16px;
  color: v-bind('color("primary_text_color")');

  &::placeholder {
    color: v-bind('color("hint_text_color")');
    opacity: 0.6;
  }
}
</style>