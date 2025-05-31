<template>
  <li
      class="list-item"
      :style="{ transitionDelay: delay + 'ms' }"
  >
    <div
      class="left-section"
      :style="getLeftSectionStyle()"
    >
      <slot name="left"></slot>
    </div>
    <div
        class="content"
        :class="{
          'content__active': active,
        }"
        @touchstart="onTouchStart"
        @touchmove="onTouchMove"
        @touchend="onTouchEnd"
        @mousedown="onTouchStart"
        @mousemove="onTouchMove"
        @mouseup="onTouchEnd"
        :style="getContentStyle()"
        @click.prevent="onClick"
    >
      <slot name="content"></slot>
    </div>
    <div
        class="right-section"
        :style="getRightSectionStyle()"
    >
      <slot name="right"></slot>
    </div>
  </li>
</template>

<script setup>
import {ref} from "vue";
import {color} from "@/theme/colors.js";

const props = defineProps({
  active: { required: true, type: Boolean },
  showLeftSection: { required: false, type: Boolean, default: false },
  delay: { required: false, type: Number, default: 0 },
})

// Touch events handling.
const touchStartX = ref(0)
const touchCurrentX = ref(0)
const swipeOffset = ref(0)  // сдвиг в px
const isMouseDown = ref(false)
const wasSwipe = ref(false)

const getClientX = (event) => {
  if (event.type.startsWith('touch')) {
    return event.changedTouches[0].clientX
  } else {
    return event.clientX
  }
}

const emit = defineEmits(['selected'])
const onClick = (event) => {
  if (wasSwipe.value) {
    event.stopImmediatePropagation()
    event.stopPropagation()
    event.preventDefault()
    return
  }

  emit('selected')
}

const onTouchStart = (event) => {
  event.preventDefault()
  event.stopPropagation()
  if (event.type === 'mousedown') isMouseDown.value = true
  const clientX = getClientX(event)
  touchStartX.value = clientX
  touchCurrentX.value = clientX
  swipeOffset.value = 0
}

const onTouchMove = (event) => {
  if (event.type === 'mousemove' && !isMouseDown.value) return
  touchCurrentX.value = getClientX(event)
  const deltaX = touchCurrentX.value - touchStartX.value

  if (Math.abs(deltaX) > 10) {
    wasSwipe.value = true
  }
  if (deltaX < 0) { // свайп влево
    swipeOffset.value = Math.max(deltaX, -180) // лимит сдвига
  }
}

const maxOffset = -76
const onTouchEnd = () => {
  setTimeout(() => {
    wasSwipe.value = false
  }, 0)
  isMouseDown.value = false
  if (swipeOffset.value < -60) {
    // Если свайп достаточно большой, фиксируем кнопку показаной
    swipeOffset.value = maxOffset
  } else {
    // Если свайп маленький, сбрасываем
    swipeOffset.value = 0
  }
}

const getContentStyle = () => {
  const marginLeft = props.showLeftSection ? '60px' : '0'
  const marginRight = `${-swipeOffset.value - 12}px`

  return {
    marginLeft: marginLeft,
    marginRight: marginRight,
    transition: 'margin-left 0.25s ease, margin-right 0.3s ease',
  }
}
const getLeftSectionStyle = () => {
  const left = props.showLeftSection ? '16px' : '-1000px'

  return {
    left: left,
    transition: 'left 0.25s ease',
  }
}
const getRightSectionStyle = () => {
  const right = swipeOffset.value ? `${-swipeOffset.value-60}px` : '-1000px'

  return {
    right: right,
    transition: 'right 0.25s ease',
  }
}
</script>

<style scoped lang="scss">
.list-item {
  margin-bottom: 10px;
  display: flex;
  width: 100%;
  min-height: 48px;
  max-height: 48px;
  justify-content: space-between;
}

.content {
  width: 100%;
  padding: 12px;
  border-radius: 12px;
  background-color: v-bind('color("secondary_bg_color")');
  transition: margin-left 0.3s ease;

  &__active {
    background-color: v-bind('color("accent_text_color")');
  }
}

.right-section {
  display: flex;
  position: absolute;

  height: 48px;
  width: 48px;
}

.left-section {
  display: flex;
  position: absolute;

  height: 48px;
  width: 48px;
}
</style>