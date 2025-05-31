<template>
  <transition-group
      name="list"
      tag="ul"
      class="list"
  >
    <slot/>
  </transition-group>
</template>

<script setup>
import {nextTick, ref} from "vue";

const scrollToBottom = () => {
  nextTick(() => {
    setTimeout(() => {
      window.scrollTo({
        top: document.body.scrollHeight,
        behavior: 'smooth' // плавная прокрутка, можно убрать
      })
    }, 100)
  })
}

defineExpose({ scrollToBottom })
</script>

<style lang="scss">
.list {
  list-style: none;
  padding: 0;
  margin-bottom: 24px;
}

.list-move,
.list-enter-active,
.list-leave-active {
  transition: all 0.2s ease;
}
.list-leave-active {
  transition-delay: 0s !important;
}
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(100px);
}
.list-leave-active {
  position: absolute;
  //position: relative;
}
</style>