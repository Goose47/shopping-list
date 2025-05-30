<template>
  <transition-group
      name="list"
      tag="div"
      class="list"
  >
    <slot/>
  </transition-group>
</template>

<script setup>
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
.list-enter-from,
.list-leave-to {
  opacity: 0;
  transform: translateX(100px);
}
.list-leave-active {
  position: absolute;
}
</style>