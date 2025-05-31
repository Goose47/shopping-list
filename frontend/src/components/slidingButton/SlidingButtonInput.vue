<template>
  <div class="wrapper">
    <div
        class="content"
        :style="getContentStyle()"
    >
      <slot name="content"></slot>
    </div>
    <div
        class="right-section"
        :style="getRightSectionStyle()"
    >
      <slot name="right"></slot>
    </div>
  </div>
</template>

<script setup>
import {ref} from "vue";
import {color} from "../../theme/colors.js";

const props = defineProps({
  showRightSection: { required: false, type: Boolean, default: false },
})

const getContentStyle = () => {
  const marginRight = props.showRightSection ? '60px' : '0'

  return {
    marginRight: marginRight,
    transition: 'margin-right 0.3s ease',
  }
}

const getRightSectionStyle = () => {
  const right = props.showRightSection ? '16px' : '-1000px'

  return {
    right: right,
    transition: 'right 0.25s ease',
  }
}
</script>

<style scoped lang="scss">
.wrapper {
  display: flex;
}

.content {
  width: 100%;
  transition: margin-left 0.3s ease;
}

.right-section {
  position: absolute;

  height: 48px;
  width: 48px;
}
</style>