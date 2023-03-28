<script setup lang="ts">
</script>

<template>
  <div class="wrapper">


    <div class="maindata">
      <div class="el">DEEPFLOWER</div><div class="el"> 
      <a href="https://t.me/graphinfinit">graphinfinit</a></div> 
    </div>

    <div class="info">
    <p>Цель проекта проверить на небольших сообществах возможность и эффективность !Many-to-Many управление проектами.</p>
    <p>Доп.цель проекта - практика работы с vue/typescript, golang, графовыми БД, архитектурой приложений</p>


    <p>Вы рождаетесь с минимальным запасом Энергии.</p>
    <p>Вы можете опубликовать Мечту о том чтобы пользователь с вашим ID стал полноправным участником сообщества</p>
    <p>Когда в результате Консенсуса ваша Мечта исполнится вам открываются возможности.</p>
    <p>Можно создавать Локации</p>
    <ul> 
        <li>Локация должна иметь уникальное название</li>
        <li>К Локации можно указать координаты.В противном случае она привязывается к нулевой координате</li>
        <li>Также можно добавить описание, высоту, радиус</li>
    </ul>
    <p>Можно создавать Мечты и публиковать их.</p>
    <ul> 
        <li>Мечта должна иметь уникальное название</li>
        <li>К Мечте нужно указать имя Локации. (может быть без координат)</li>
        <li>Публикация расходует определенное количество Энергии. После публикации Мечту нельзя будет изменять. </li>
        <li>Можно добавлять Энергию к своим и чужим опубликованым Мечтам, расходуя свою Энергию</li>

    </ul>

    <p>К Мечте можно создавать проекты</p>
    <ul> 
        <li>У Проекта есть уникальное в рамках мечты имя</li>
        <li>У Проекта есть Граф задач с входной и выходной Нодой. Внутри графа нет циклических зависимостей</li>
        <li>Проекту можно добавить свою Энергию</li>
        <li>У каждой задачи-Ноды есть вес. Сумма весов равна 1.</li>
        <li> Исходя из Энергии проекта и веса рассчитывается Энергия задачи</li>
        <li>Когда достигается выходная Нода - мечта реализована </li> 
    </ul>

    <p>?Источники Энергии</p>
    <ul> 
        <li>?Начальная энергия</li>
        <li>?Выполнение задач</li>
        <li>?Создание Проектов</li>
        <li>?Создание Мечт</li>
        <li>...</li>
        <li>...</li> 
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