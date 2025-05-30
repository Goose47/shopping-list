<template>
  <app-page title="Списки">
    <swipeable-list-component>
      <swipeable-list-item-component
          v-for="(list, idx) in filteredLists"
          :key="list.id"
          :show-left-section="!!selectedList"
          :active="isListCompleted(list)"
          :delay="idx * 50"
      >
        <template #left>
          <app-button @click="unselectList">
            <arrow-left-icon/>
          </app-button>
        </template>

        <template #content>
          <div @click="selectList(list)">
            {{ list.name }}
          </div>
        </template>

        <template #right>
          <app-button @click="removeList(list)">
            <cross-icon/>
          </app-button>
        </template>
      </swipeable-list-item-component>
    </swipeable-list-component>

    <swipeable-list-component>
      <swipeable-list-item-component
          v-for="(item, idx) in filteredItems"
          :key="item.id"
          :delay="idx * 50"
          :active="item.checked"
       >
        <template #left>
          <app-button @click="unselectList">
            <arrow-left-icon/>
          </app-button>
        </template>

        <template #content>
          <div @click="checkItem(item)">
            {{ item.name }}
          </div>
        </template>

        <template #right>
          <app-button @click="removeItem(item)">
            <cross-icon/>
          </app-button>
        </template>
      </swipeable-list-item-component>
    </swipeable-list-component>
  </app-page>
</template>

<script setup>
import {color} from '@/theme/colors.js'
import {ref, computed} from "vue";
import testData from './testData.json'
import ArrowLeftIcon from "@/components/icons/ArrowLeftIcon.vue";
import CrossIcon from "@/components/icons/CrossIcon.vue";
import SwipeableListComponent from "../../../components/swipeableList/SwipeableListComponent.vue";
import SwipeableListItemComponent from "../../../components/swipeableList/SwipeableListItemComponent.vue";

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

const removeList = (list) => {
  lists.value = lists.value.filter(l => l.id !== list.id)
  if (selectedList.value) {
    selectedList.value = null
  }
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

const removeItem = (item) => {
  selectedList.value.items = selectedList.value.items.filter(i => i.id !== item.id)
}
</script>

<style scoped lang="scss">
</style>