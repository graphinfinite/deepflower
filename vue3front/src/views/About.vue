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
    <p>Можно создавать локации</p>
    <ul> 
        <li>Локация должна иметь уникальное название</li>
        <li>К локации можно указать координаты.В противном случае она привязывается к нулевой координате</li>
        <li>Также можно добавить описание, высоту, радиус</li>
    </ul>
    <p>Можно создавать мечты и публиковать их.</p>
    <ul> 
        <li>Мечта должна иметь уникальное название</li>
        <li>К мечте нужно указать имя локации. (может быть без координат)</li>
        <li>Публикация расходует определенное количество энергии. После публикации мечту нельзя будет изменять. </li>
        <li>Можно добавлять энергию к своим и чужим опубликованым мечтам, расходуя свою энергию</li>

    </ul>
    </div>
  </div>


</template>

<style scoped lang="scss">


@use '@/assets/scss/_colors' as clr;


.maindata{
  display: flex;
  flex-direction: row;
  color: blueviolet;
  padding: 20px;
  background-color: clr.$bg-route-header;
}
.maindata .el {
  padding: 10px;
}


.info {
  color: rgb(56, 10, 10);
  padding:20px;
  border: 1px solid whitesmoke;
}


.info p {
  margin-top: 20px;
}


.info ul {


}

.info li{
    padding:6px;
}
.info li:before {
    padding-right:10px;
    font-weight: bold;
    color: #d1ceff;
    content: ".";
    transition-duration: 0.5s;
}
.info li:hover:before {
    color: #024C6F;
    content: ".";
} 


.wrapper{
 background-color: white;

}
</style>