<template>
 <div>Hello from protected!</div>
</template>

<script setup>
const tg = window.Telegram.WebApp
tg.MainButton.setText("Подтвердить");
tg.MainButton.show();
tg.MainButton.onClick(() => {
  tg.sendData(JSON.stringify({ action: "confirm" }));
});

// Secondary Button 1
const cancelBtn = tg.createSecondaryButton();
cancelBtn.setText("Отмена");
cancelBtn.onClick(() => tg.close());

// Secondary Button 2
const infoBtn = tg.createSecondaryButton();
infoBtn.setText("Подробнее");
infoBtn.onClick(() => {
  alert("Здесь можно показать дополнительную информацию.");
});
</script>

<style scoped lang="scss">

</style>