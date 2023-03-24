<script setup lang="ts">
</script>

<template>
  <div class="wrapper">


    <div class="maindata">
      <div class="el">DEEPFLOWER</div><div class="el"> 
      <a href="https://t.me/graphinfinit">graphinfinit</a></div> 
    </div>

    <div class="info">
    <p>Вы рождаетесь с некоторым запасом энергии.</p>
    <p>Можно создавать мечты и публиковать их.</p>
    <p>Публикация расходует минимальное количество энергии</p>
    <p>Можно добавлять энергию к своим и чужим опубликованым мечтам, расходуя свою энергию</p>
    <p>К мечте можно указать местоположение 3мя способами:</p>
    <ul> 
        <li>-Tекстовое описание</li>
        <li>-Kоординаты</li>
        <li>-Прикрепить созданную ранее локацию</li>
      </ul>
    </div>






    




  </div>


</template>

<style scoped lang="css">

.maindata{
  display: flex;
  flex-direction: row;
  color: azure;
  padding: 20px;
  background-color: #172025;
}
.maindata .el {
  padding: 10px;
}


li{
    padding:6px;
}
 li:before {
    padding-right:10px;
    font-weight: bold;
    color: #d1ceff;
    content: ".";
    transition-duration: 0.5s;
}
 li:hover:before {
    color: #024C6F;
    content: ".";
} 


.wrapper{
  padding: 20px;
  margin-top: 30px;
}
</style>