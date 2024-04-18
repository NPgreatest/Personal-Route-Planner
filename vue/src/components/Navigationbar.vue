<template xmlns:v-slot="http://www.w3.org/1999/XSL/Transform">
  <div class="navi-bar">

    <div class="navi-bar-center">
      <p class="title">{{ '欢迎 ' }}  {{userName}} {{' !'}}</p>
<!--      <el-button v-if="stat=='已登录'" @click="info" class="title3" type="text" >(个人信息)</el-button>-->
<!--      <el-button v-if="stat=='已登录'" @click="last" class="title2" type="text" >(退出)</el-button>-->
      <div class="navi-bar-item-group">
<!--        <i @click="fullscreen" class="iconfont icon-quanpingmu fullscreen"></i>-->
        <!--导航按键-->
        <ul>
          <li><NaviItem :activeRoute="activeRoute" index-name="首页"  rout="/home"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="问答"  rout="/chat"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="景点展示" rout="/sitedetails?id=1"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="景点推荐" rout="/tags"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="路线规划"  rout="/plan"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="留言板" rout="/msgBoard"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="关于"  rout="/about"></NaviItem></li>
          <li><NaviItem :activeRoute="activeRoute" index-name="登录" rout="/login"   ></NaviItem></li>
        </ul>
        <el-dropdown v-if="stat === '已登录'">
          <template v-slot:dropdown>
            <el-dropdown-menu class="vertical-dropdown-menu">
              <el-dropdown-item class="vertical-dropdown-item" @click.native="info">个人信息</el-dropdown-item>
              <el-dropdown-item class="vertical-dropdown-item" @click.native="last">退出登录</el-dropdown-item>
            </el-dropdown-menu>

          </template>
          <span class="avatar-dropdown">
            <b-avatar class="avatar" :src="avatar" size="3rem" ></b-avatar>
          </span>
        </el-dropdown>
      </div>
    </div>
  </div>
</template>

<script>
import NaviItem from "./NaviItem";
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Navigationbar",
  userName:'',


  components: { NaviItem },
  data(){
      return{
      activeRoute:"/home",
        stat:'登录',
        avatar:'',
      }
  },
  methods:{
    handleCommand(command) {
      if (command === 'info') {
        this.info();
      } else if (command === 'last') {
        this.last();
      }
    },
    fullscreen(){
        const de = document.documentElement;
        if(de.requestFullscreen) {
          de.requestFullscreen();
        }
        window.scrollTo(0, 1);
      },
    last(){
      window.sessionStorage.clear()
      location.reload()
    },
    info(){
      this.$router.push("/userinfo")
    },
    getDetails: async function() {
      const {data: res} = await this.$axios.get("/user/getinfo");
      console.log(res)
      if(res.status === 1) {
        this.avatar = res.data.length > 0 ? res.data[0].avatar : this.avatar;
        console.log(this.avatar)
      }
    },
  },

  created() {
    this.activeRoute = this.$router.currentRoute.path
    this.userName=window.sessionStorage.getItem("userId");
    this.avatar=window.sessionStorage.getItem("avatar")

    if(this.userName==null){
      this.stat='未登录'
    }else{
      this.stat='已登录'
    }
    this.getDetails().then(() => {
      console.log('done')
    })
  },

  watch: {
    // 监听路由改变，路由改变时给activeRoute赋值，从而让导航栏组件高亮
    $route(to) {
      this.activeRoute = to.path
    }
  }
}
</script>

<style lang="less" scoped>
ul {
  margin: 0;
  display: inline-block;
  color: #add8e6; /* Light blue */
}

li {
  display: inline-block;
  list-style: none;
  margin: 0 15px;
  font-weight: bold; /* Making text bold */
}

li a, .el-dropdown-item {
  color: inherit; /* Ensure anchor tags inside list items inherit the color from li */
  font-weight: inherit; /* Bold text for dropdown items */
}

li a:hover, .el-dropdown-item:hover {
  color: #00008b; /* Dark blue on hover */
}

/* Additional styles for dropdown interaction */
.el-dropdown-menu {
  .el-dropdown-item {
    &:hover {
      background-color: #add8e6; /* Light blue background on hover */
    }
    &:active {
      background-color: #00008b; /* Dark blue background when active/clicked */
      color: #fff; /* White text when active/clicked */
    }
  }
}

.title {
  float: left;
  line-height: 1.5;
  text-align: center;
  font-size: 24px;
  align-items: center;
  margin: 10px 0 0 50px;
  height: 40px;
  background: -webkit-linear-gradient(45deg, #000000, #0f259f, #0f259f, #000000, #bed5f5);
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
  background: rgba(210, 249, 253, 0.3);
}
.vertical-dropdown-menu {
  display: block; /* 确保使用块状布局 */
  flex-direction: column; /* 确保子元素垂直排列，如果父元素使用了flex布局 */
}

.vertical-dropdown-item {
  display: block; /* 单个项也确保为块状布局 */
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
  color: #1ec3ff;
  font-size: 20px;
  cursor: pointer;
  margin-right: 20px;
}
</style>
