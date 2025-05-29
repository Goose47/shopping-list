<template>
  <app-page title="Списки">
    <transition-group
        name="lists"
        tag="div"
        class="list"
    >
      <div
          class="list-item"
          v-for="list in filteredLists"
          :key="list.id"
      >
        <app-button
            class="button__sliding"
            :class="{
              'button__sliding__active': !!selectedList,
              'button__sliding__left': !!selectedList,
            }"
            @click="unselectList"
        >
          <arrow-left-icon/>
        </app-button>
        <div
            class="list-item__content"
            :class="{
              'list-item__content__shrink': selectedList && selectedList.id === list.id,
              'list-item__content__active': isListCompleted(list),
            }"
            @click="selectList(list)"
            @touchstart="onTouchStart($event, list.id)"
            @touchmove="onTouchMove($event, list.id)"
            @touchend="onTouchEnd(list.id)"
        >
          {{ list.name }}
        </div>
        <div
            class="list-item__buttons"
            :style="getSwipeStyleForContainer(list.id)"
        >
          <app-button
            :style="getSwipeStyleForButton(list.id)"
          >
            <arrow-left-icon/>
          </app-button>
          <app-button
            :style="getSwipeStyleForButton(list.id)"
          >
            <arrow-left-icon/>
          </app-button>
        </div>
      </div>
    </transition-group>

    <transition-group
        name="items"
        tag="div"
        class="list"
    >
      <div
          class="list-item"
          v-for="(item, idx) in filteredItems"
          :key="item.id"
          :style="{ transitionDelay: (idx * 100) + 'ms' }"
      >
        <div
            class="list-item__content"
            :class="{
              'list-item__content__active': item.checked,
            }"
            @click="checkItem(item)"
        >
          {{ item.name }}
        </div>
      </div>
    </transition-group>

<!--    <app-button class="button">Добавить список</app-button>-->
  </app-page>
</template>

<script setup>
import {color} from '@/theme/colors.js'
import {ref, computed} from "vue";
import testData from './testData.json'
import ArrowLeftIcon from "@/components/icons/ArrowLeftIcon.vue";

const lists = ref([])
const getLists = () => {
  lists.value = testData
}
getLists()

const filteredLists = computed(() => {
  if (!selectedList.value) {
    return lists.value
  }
  return lists.value.filter(l => {
    return l.id === selectedList.value.id
  })
})

const selectedList = ref(null)
const selectList = (list) => {
  selectedList.value = list
}
const unselectList = () => {
  selectedList.value = null
}

const isListCompleted = (list) => {
  return list.items.filter(i => !i.checked).length === 0
}

const filteredItems = computed(() => {
  if (!selectedList.value) {
    return []
  }
  return selectedList.value.items
})

const checkItem = (item) => {
  item.checked = !item.checked
}

// ttouches.
const touchStartX = ref(0)
const touchCurrentX = ref(0)
const swipedListId = ref(null)
const swipeOffset = ref(0)  // сдвиг в px

const onTouchStart = (event, id) => {
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
    swipeOffset.value = Math.max(deltaX, -150) // лимит сдвига -150px
    swipedListId.value = id
  }
}

const onTouchEnd = (id) => {
  if (swipeOffset.value < -80) {
    // Если свайп достаточно большой, фиксируем кнопку показаной
    swipeOffset.value = -120
    swipedListId.value = id
  } else {
    // Если свайп маленький, сбрасываем
    swipeOffset.value = 0
    swipedListId.value = null
  }
}

const getSwipeStyleForContainer = (id) => {
  if (swipedListId.value === id) {
    return {
      'width': `${-swipeOffset.value}px`,
      'min-width': `${-swipeOffset.value}px`,
      transition: swipeOffset.value === 0 ? 'width 0.3s ease' : 'none',
    }
  }
  return {
    width: '0',
    transition: 'width 0.3s ease',
  }
}

const getSwipeStyleForButton = (id) => {
  if (swipedListId.value === id) {
    return {
      'min-height': '48px',
      'max-width': '48px',
      'margin-left': '12px',
    }
  }
  return {
    'min-height': '48px',
    padding: '0',
    transition: 'width 0.3s ease',
  }
}

</script>

<style scoped lang="scss">
.list {
  list-style: none;
  padding: 0;
  margin-bottom: 24px;
}

.list-item {
  margin-bottom: 10px;
  display: flex;
  width: 100%;
  min-height: 48px;
  max-height: 48px;
  justify-content: space-between;

  &__content {
    width: 100%;
    padding: 12px;
    border-radius: 12px;
    background-color: v-bind('color("secondary_bg_color")');
    transition: width 0.3s ease;

    &__shrink {
      width: calc(100% - 60px) !important;
    }
    &__active {
      background-color: v-bind('color("accent_text_color")');
    }
  }

  &__buttons {
    display: flex;
    height: 48px;
  }
}

.button__sliding {
  width: 0 !important;
  padding: 0 !important;
  transition: all 0.3s ease;
  min-height: 48px;

  &__left {
    margin-right: 12px;
  }

  &__active {
    min-width: 48px;
  }
}

.lists-move,
.lists-enter-active,
.lists-leave-active {
  transition: all 0.2s ease;
}
.lists-enter-from,
.lists-leave-to {
  opacity: 0;
  transform: translateX(100px);
}
.lists-leave-active {
  position: absolute;
}

.items-move,
.items-enter-active,
.items-leave-active {
  transition: all 0.3s ease;
}
.items-enter-from,
.items-leave-to {
  opacity: 0;
  transform: translateX(100px);
}
.items-leave-active {
  position: absolute;
}
</style>