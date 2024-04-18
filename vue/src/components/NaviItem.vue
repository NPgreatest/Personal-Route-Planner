<template>
  <router-link :to="rout" class="nav-item" :class="{active: isActive}">
    <i class="iconfont" :class="icoClass"></i>
    {{ indexName }}
  </router-link>
</template>

<script>
import "../assets/icon/iconfont.css"
export default {
  name: "NaviItem",
  props: ["indexName", "icoClass", "rout", "activeRoute"],
  computed: {
    isActive() {
      return this.rout === this.activeRoute
    }
  }
}
</script>

<style lang="less" scoped>
i {
  color: #fda085; /* Keeping your existing icon color, remove !important if not necessary */
}

.nav-item {
  line-height: 40px;
  cursor: pointer;
  padding-bottom: 5px;
  text-decoration: none;
  color: #000000; /* Light blue for the default state */
  font-weight: bold; /* Making the text bold */
}

.nav-item:hover {
  color: #00008b; /* Deep blue on hover */
  border-bottom: 2px solid #00008b; /* Adding deep blue underline on hover */
}

.active {
  color: #00008b; /* Deep blue when active */
  border-bottom: 4px solid #00008b; /* Deeper underline for active state */
}
</style>
