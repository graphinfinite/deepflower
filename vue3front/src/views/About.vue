<script setup lang="ts">
</script>

<template>
  <div class="wrapper">


    <div class="maindata">
      <div class="el">DEEPFLOWER</div><div class="el"> 
      <a href="https://t.me/graphinfinit">graphinfinit</a></div> 
    </div>

    <div class="info">
    <p>Вы рождаетесь с запасом Энергии.</p>
    <p>Можно создавать Локации</p>
    <ul> 
        <li>Локация должна иметь уникальное название</li>
        <li>К Локации можно указать координаты.В противном случае она привязывается к нулевой координате</li>
        <li>Также можно добавить описание, высоту, радиус</li>
    </ul>
    <p>Можно создавать Мечты и публиковать их.</p>
    <ul> 
        <li>Мечта должна иметь уникальное название</li>
        <li>К Мечте нужно указать имя Локации.</li>
        <li>Публикация расходует n Энергии. После публикации Мечту нельзя будет изменять. </li>
        <li>Можно добавлять Энергию к своим и чужим опубликованым Мечтам, расходуя свою Энергию</li>

    </ul>

    <p>Можно создавать проекты и публиковать их</p>
    <ul> 
        <li>У Проекта должно быть уникальное имя</li>
        <li>Публикация расходует n Энергии. После публикации Проект нельзя будет изменять.</li>
        <li>У Проекта есть Граф задач с входной (START) и выходной (END) Нодой. 
          Когда достигается выходная Нода - Мечта реализована
        </li>
    </ul>

    <p>Граф задач</p>
    <ul> 
        <li>В графе задач имеются 7 типов Нод: STAR, END, SLOW, FAST, CHOICE, MULTI, FAIL</li>
        <li>STAR, END - начало и конец. Начало одно. Завершений много.
           В результате разных альтернативных путей можно прийти к разным конечным Нодам
        </li>
        <li>SLOW, FAST - ноды задач. Условно делятся на быстрые и медленные для удобства восприятия графа</li>
        <li>CHOICE - Нода выбора.</li>
        <li>MULTI - синтаксический сахар для схлопывания путей множественного выбора после CHOICE. 
          На самом деле граф после мультиплексора (MULTI) клонируется столько раз сколько входов множественного выбора после CHOICE.
        </li>
        <li>FAIL - завершающая нода (Невозможно успешно завершить проект в данных обстоятельствах)</li> 

        <li>В графе не должно быть циклических зависимостей, не содержащих в себе Ноду выбора</li>
        <li>Выполнение задач (SLOW, FAST) происходит последовательно или параллельно.
           !Задачу нельзя выполнить если не выполнены все входные ноды задач.
        </li>

    </ul>

    <p>Задача</p>
    <ul>
      <li>К задаче из опубликованного проекта можно добавлять энергию</li>
      <li>У задачи есть максимальное время выполнения</li>
      <li>Задачу можно брать в работу. Нужно успеть выполнить ее за выделенное время</li>
      <li>После выполнения нужно закрыть задачу. После закрытия будет запущен процесс Консенсуса</li>
    </ul>

    <p>Консенсус</p>
    <li>Простой консенсус: выполнение задачи должны подтвердить все участники, энергнтически учавствующие в задаче</li>
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
  font-size: clr.$route-header-fontsize;
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
  margin-bottom: 10px;
  color:#024C6F;
}


.info ul {
  color: grey;


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