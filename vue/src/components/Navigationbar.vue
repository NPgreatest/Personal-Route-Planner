<template>
  <div class="navi-bar">
    <p>{{ '欢迎!' }}</p>
    <div class="navi-bar-center">
      <div class="navi-bar-item-group">
        <i @click="fullscreen" class="iconfont icon-quanpingmu fullscreen"></i>
        <!--导航按键-->
        <ul>
          <li><NaviItem :activeRoute="activeRoute" index-name="首页"  rout="/home"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="问答"  rout="/types"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="景点展示" rout="/sitedetails?id=1"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="景点推荐" rout="/timeLine"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="路线规划"  rout="/essay"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="留言板" rout="/msgBoard"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="关于"  rout="/about"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="登录"  rout="/login"></NaviItem></li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
import NaviItem from "./NaviItem";
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Navigationbar",
  components: { NaviItem },
  data(){
      return{
      activeRoute:"/home"
      }
  },
  methods:{
      fullscreen(){
        const de = document.documentElement;
        if(de.requestFullscreen) {
          de.requestFullscreen();
        }
        window.scrollTo(0, 1);
      }
  },
  created() {
    this.activeRoute = this.$router.currentRoute.path
  },
  watch: {
    // 监听路由改变，路由改变时给activeRoute赋值，从而让导航栏组件高亮
    $route(to) {
      this.activeRoute = to.path
    }
  }
}
</script>


<style lang="less" scoped >

ul {
  margin: 0;
  display: inline-block;
}

li {
  display: inline-block;
  list-style: none;
  margin: 0 15px;
}

.search-warp {
  position: relative;
  display: inline-block;
}

.search-result {
  width: 240px;
  position: absolute;
  background-color: #fff;
  top: 45px;
  left: -20px;
  border-radius: 8px;

  .search-content {
    overflow-y:auto;
    max-height: 300px;
    margin: 8px 0;
  }
}

.search-result::before {
  content: '';
  position: absolute;
  display: inline-block;
  width: 16px;
  height: 16px;
  border-left: 1px solid #fff;
  border-top: 1px solid #fff;
  transform: rotate(45deg);
  top: -8px;
  left: 50px;
  background-color: #fff;
  border-top-left-radius: 3px;
}

.search-result ul  {
  padding-left: 16px;
  font-size: 14px;

  li {
    display: block;
    cursor: pointer;
    margin: 2px 0;

    a {
      outline: none;
      text-decoration: none;
      color: #61666d;
    }
  }

  li:first-child {
    cursor: default;
  }

}

p {
  float: left;
  line-height: 1.5;
  text-align: center;
  font-size: 24px;
  align-items: center;
  margin: 10px 0 0 50px;
  height: 40px;
  background: -webkit-linear-gradient(45deg,#70f7fe,#fbd7c6,#fdefac,#bfb5dd,#bed5f5);
  -webkit-background-clip: text;
  font-weight: bolder;
  color: transparent;
  user-select: none;
}

.navi-bar {
  z-index: 999;
  width: 100%;
  height: 60px;
  position: fixed;
  background: rgba(210,249,253,0.3);
}

.navi-bar-center {
  margin: 10px 0;
  height: 40px;
  width: 100%;
}

.navi-bar-item-group {
  display: inline-block;
  align-items: center;
  height: 100%;
  float: right;
  margin-right: 100px;
}

.fullscreen {
  color: #fda085;
  font-size: 20px;
  cursor: pointer;
  margin-right: 20px;
}

</style>
