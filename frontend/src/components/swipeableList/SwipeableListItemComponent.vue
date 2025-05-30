<template>
  <div
      class="list-item"
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
        @touchstart="onTouchStart($event)"
        @touchmove="onTouchMove($event, 1)"
        @touchend="onTouchEnd(1)"
        :style="getContentStyle(1)"
    >
      <slot name="content"></slot>
    </div>
    <div
        class="right-section"
        :style="getRightSectionStyle(1)"
    >
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {color} from "../../theme/colors.js";

const props = defineProps({
  active: { required: true, type: Boolean },
  showLeftSection: { required: false, type: Boolean, default: false },
})

// Touch events handling.
const touchStartX = ref(0)
const touchCurrentX = ref(0)
const swipedListId = ref(null)
const swipeOffset = ref(0)  // сдвиг в px

const onTouchStart = (event) => {
  const touch = event.changedTouches[0]
  touchStartX.value = touch.clientX
  touchCurrentX.value = touch.clientX
  swipedListId.value = null
  swipeOffset.value = 0
}

const onTouchMove = (event, id) => {
  const touch = event.changedTouches[0]
  touchCurrentX.value = touch.clientX
  let deltaX = touchCurrentX.value - touchStartX.value

  if (deltaX < 0) { // свайп влево
    swipeOffset.value = Math.max(deltaX, -180) // лимит сдвига
    swipedListId.value = id
  }
}

const maxOffset = -76
const onTouchEnd = (id) => {
  if (swipeOffset.value < -60) {
    // Если свайп достаточно большой, фиксируем кнопку показаной
    swipeOffset.value = maxOffset
    swipedListId.value = id
  } else {
    // Если свайп маленький, сбрасываем
    swipeOffset.value = 0
    swipedListId.value = null
  }
}

const getContentStyle = (id) => {
  const marginLeft = props.showLeftSection ? '60px' : '0'
  const marginRight = swipedListId.value === id ? `${-swipeOffset.value - 12}px` : '0'

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
const getRightSectionStyle = (id) => {
  const right = swipedListId.value === id ? '16px' : '-1000px'

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