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
            class="back-button"
            :class="{'back-button__active': !!selectedList}"
            @click="unselectList"
        >
          <arrow-left-icon/>
        </app-button>
        <div
            class="list-item__text"
            :class="{ 'list-item__text__shrink': selectedList && selectedList.id === list.id }"
            @click="selectList(list)"
        >
          {{ list.name }}
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
            class="list-item__text"
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

const filteredItems = computed(() => {
  if (!selectedList.value) {
    return []
  }
  return selectedList.value.items
})
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
  justify-content: space-between;

  &__text {
    width: 100%;
    padding: 12px;
    border-radius: 12px;
    background-color: v-bind('color("secondary_bg_color")');
    transition: width 0.3s ease;

    &__shrink {
      width: calc(85%);
    }
  }
}

.back-button {
  transition: all 0.3s ease;
  width: 0 !important;
  min-height: 48px;
  padding: 0 !important;

  &__active {
    display: block;
    width: 10%;
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