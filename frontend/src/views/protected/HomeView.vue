<template>
  <div
      v-for="color in colors"
      :style="{
        'background-color': wa.themeParams['color'],
        'color': wa.themeParams['color'],
        'width': '100px',
        'height': '40px',
        'margin-bottom': '10px',
      }"
  >
    {{color}}
  </div>
  <div>{{ tg }}</div>
</template>

<script setup>
const tg = window.Telegram
const wa = tg.WebApp
wa.MainButton.setText("Подтвердить");
wa.MainButton.show();
wa.MainButton.onClick(() => {
  wa.sendData(JSON.stringify({ action: "confirm" }));
});

// Secondary Button 1
const cancelBtn = wa.createSecondaryButton();
cancelBtn.setText("Отмена");
cancelBtn.onClick(() => wa.close());

// Secondary Button 2
const infoBtn = wa.createSecondaryButton();
infoBtn.setText("Подробнее");
infoBtn.onClick(() => {
  alert("Здесь можно показать дополнительную информацию.");
});

const colors = [
  'accent_text_color',
  'bg_color',
  'bottom_bar_bg_color',
  'button_color',
  'button_text_color',
  'destructive_text_color',
  'header_bg_color',
  'hint_color',
  'link_color',
  'secondary_bg_color',
  'section_bg_color',
  'section_header_text_color',
  'section_separator_color',
  'subtitle_text_color',
  'text_color',
]
</script>

<style scoped lang="scss">

</style>