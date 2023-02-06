<template>
  <div>
    <transition appear
                name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__slideInDown"
                leave-active-class="animate__slideOutUp"
    >
      <NavigationBar class="navibar" v-show="showNavi"></NavigationBar>
    </transition>
    <keep-alive include="Home">
      <router-view ref="home"></router-view>
    </keep-alive>
  </div>
</template>

<script>
import NavigationBar from "../components/Navigationbar";
import "animate.css"
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "index",
  components:{NavigationBar},
  data() {
    return {
      showNavi: true,
    }
  },
}
</script>

<style scoped>
.animate__animated{
  animation-duration: 1.5s !important;
}

.navibar {
  position: fixed;
  top: 0;
}
</style>
