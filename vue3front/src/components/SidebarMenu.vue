<script setup lang="ts">
import { shallowRef } from 'vue';
import { RouterLink } from 'vue-router';
 import IconAngle from './icons/IconAngle.vue';
const isSidebarOpen = shallowRef<Boolean>(false)
function toggleSidebar() {
  isSidebarOpen.value = !isSidebarOpen.value;
}
</script>

<template>
  <div class="wrapper">
    <aside :vue:is-open="isSidebarOpen">

      <h4 :transparent="!isSidebarOpen">*</h4>
      <ul class="sidebar-head">
        <li>
          <img src="@/assets/triangle-32.ico" alt="logo" width="32" height="32">
        </li>
        <li>
          <button class="sidebar-toggle" :class="isSidebarOpen ? 'toggle-button' : ''">
            <IconAngle @click="toggleSidebar" />
          </button>
        </li>
      </ul>

      <ul>
        <li>
          <router-link to="/">
            <img src="@/assets/white-home-icon-png-21.png" alt="logo" width="32" height="32">
            <Transition name="fade">
              <span v-show="isSidebarOpen">Home</span>
            </Transition>
          </router-link>
        </li>
        <li>
          <router-link to="/locations">
            <img src="@/assets/planet-32.png" alt="logo" width="32" height="32">
            <Transition name="fade">
              <span v-show="isSidebarOpen">Locations</span>
            </Transition>
          </router-link>
        </li>
        <li>
          <router-link to="/dreams">
            <img src="@/assets/icons8-огонь-48.png" alt="logo" width="32" height="32">
            
            <Transition name="fade">
              <span v-show="isSidebarOpen">Dreams</span>
            </Transition>
          </router-link>
        </li>
        <li>
          <router-link to="/tasks">
            <img src="@/assets/icongraph.png" alt="logo" width="32" height="32">
            
            <Transition name="fade">
              <span v-show="isSidebarOpen">Tasks</span>
            </Transition>
          </router-link>
        </li>

        <li>
          <router-link to="/settings">
            <img src="@/assets/settings-5-32.png" alt="logo" width="32" height="32">
            
            <Transition name="fade">
              <span v-show="isSidebarOpen">Settings</span>
            </Transition>
          </router-link>
        </li>
        <li>
          <router-link to="/about">
            <img src="@/assets/info-6-32.png" alt="logo" width="32" height="32">
            <Transition name="fade">
              <span v-show="isSidebarOpen">About</span>
            </Transition>
          </router-link>
        </li>
      </ul>
    </aside>
  </div>
</template>

<style scoped lang="scss">
@use '@/assets/scss/_colors' as clr;
$sidebar-width: 4rem;
$toggle-duration: 300ms;
$sidebar-padding-inline-start: 1rem;


aside {
  color: clr.$primary;
  background: clr.$bg-dark;
  display: flex;
  flex-direction: column;
  min-height: 100vh;
  height: 100%;
  padding-block: 1rem;
  transition: all $toggle-duration;
  width: $sidebar-width;
}
aside[vue\:is-open=true] {
  width: 3 * $sidebar-width;
}
ul {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
  padding-block-end: 1rem;
}
img {
  object-fit: contain;
}
li {
  min-width: fit-content;
  cursor: pointer;
  padding-inline-start: $sidebar-padding-inline-start;
  &:hover {
    color: clr.$secondary;
    background-color: lighten($color: clr.$bg-dark, $amount: 5);
  }
  & a {
    // border-right: 0.25rem solid white;
    display: flex;
    align-items: center;
    column-gap: 0.75rem;
    position: relative;
    padding-block: 0.5rem;
  }
  a.router-link-exact-active::after {
    content: '';
    position: absolute;
    right: 0;
    width: 0.25rem;
    height: 100%;
    background-color: clr.$secondary;
  }
}
.sidebar-head {
  position: relative;
  padding-block-end: 4rem;
}
.sidebar-toggle {
  padding-inline-start: $sidebar-padding-inline-start;
}
h4 {
  padding-block-end: 1rem;
  padding-inline-start: $sidebar-padding-inline-start;
  user-select: none;
  pointer-events: none;
  opacity: 0.5;
  text-transform: uppercase;
  font-size: 0.75rem;
  letter-spacing: 0.125ch;
  transition: opacity $toggle-duration;
}
h4[transparent=true] {
  opacity: 0;
}
button {
  cursor: pointer;
  position: absolute;
  transition-duration: $toggle-duration;
  transition-property: transform, left, top;
  left: 0;
  top: 1rem;
  transform: translateX(0%) translateY(2rem) rotateZ(0deg);
  &.toggle-button {
    left: 100%;
    top: 0;
    transform: translateX(-100%) translateY(0rem) rotateZ(180deg);
  }
}
.fade-enter-active,
.fade-leave-active {
  transition-property: opacity, transform;
  transition-duration: $toggle-duration;
}
.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform: translateX(-100%);
}
</style>