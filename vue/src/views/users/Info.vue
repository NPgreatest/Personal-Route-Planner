
<template>

  <el-container class="home_container">

    <!-- 头部区域 -->
    <el-header>
      <div>
        <span>用户信息管理后台</span>
      </div>
      <div class="grid-content bg-white" style="margin-right: 1300px;"></div>
      <el-button type="info" @click="homepage">返回主界面</el-button>
      <el-button type="info" @click="logout">退出登录</el-button>
    </el-header>

    <el-container>
      <!-- 侧边区域 -->
      <el-aside :width="isCollapse ? '60px' : '230px'"  >
        <div class="toggle-button" @click="toggleCollapse">|||</div>
        <el-menu background-color="#333744" text-color="#fff" active-text-color="#409eff"
                 unique-opened :collapse="isCollapse" :collapse-transition="false" router
                 :default-active="activePath" >

          <el-submenu :index="menu.id" v-for="(menu, i1) in menus" :key="menu.id">
            <!-- 一级菜单的模板区域 -->
            <template v-slot:title>
              <!-- 图标 -->
              <i :class="menu.icon"></i>
              <!-- 文本 -->
              <span>{{ menu.name }}</span>
            </template>

            <!-- 二级菜单 -->
            <el-menu-item :key="subMenu.id" v-for="(subMenu, i2) in menu.submenus" :index="subMenu.index" @click="saveNavState(i1, i2)">
              <template v-slot:title>
                <!-- 图标 -->
                <i :class="subMenu.icon"></i>
                <!-- 文本 -->
                <span>{{ subMenu.name }}</span>
              </template>
            </el-menu-item>
          </el-submenu>
        </el-menu>-->
      </el-aside>

      <!-- 主体区域 -->
      <el-main>
        <el-breadcrumb separator-class="el-icon-arrow-right">
          <el-breadcrumb-item>{{ breadcrumbItem1 }}</el-breadcrumb-item>
          <el-breadcrumb-item>{{ breadcrumbItem2 }}</el-breadcrumb-item>
        </el-breadcrumb>
        <!-- 路由占位符 -->
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>


</template>



<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:"Info",
  data() {
    return {
      isCollapse: false,
      activeMenu: -1,
      activeSubMenu: -1,

      menus: [
        {
          id :"1",
          icon: "el-icon-edit",
          name: "用户信息管理",
          submenus: [
            {id:"11", index: "/userinfo", icon:"el-icon-menu", name:"个人信息"},
            {id:"12", index: "/userrate", icon:"el-icon-plus", name:"评分信息"},
            {id:"13", index: "/usercomment", icon:"el-icon-menu", name:"管理留言"}
          ]
        },
        {
          id :"2",
          icon: "el-icon-notebook-2",
          name: "景点管理",
          submenus: [
            {id:"21", index: "/sites", icon:"el-icon-menu", name:"查看景点"},
          ]
        },
        {
          id :"3",
          icon: "el-icon-paperclip",
          name: "标签管理",
          submenus: [
            {id:"31", index: "/sitestags", icon:"el-icon-menu", name:"查看景点标签"},
          ]
        },
        {
          id :"5",
          icon: "el-icon-notebook-1",
          name: "评论管理",
          submenus: [
            {id:"51", index: "/comments", icon:"el-icon-menu", name:"管理所有评论"},
          ]
        },
        {
          id :"6",
          icon: "el-icon-chat-line-square",
          name: "留言管理",
          submenus: [
            {id:"61", index: "/borad", icon:"el-icon-menu", name:"管理留言"},
          ]
        },
        {
          id :"7",
          icon: "el-icon-star-off",
          name: "数据库管理",
          submenus: [
            {id:"71", index: "/mysql", icon:"el-icon-menu", name:"调用Mysql"},
          ]
        }
      ]
    }
  },
  methods: {
    homepage(){
      this.$router.push("/home");
    },
    logout() {
      // 清空token
      window.sessionStorage.clear();
      // 跳转登录页
      this.$router.push("/login");
    },
    toggleCollapse() {   //单机按钮实现菜单的折叠与展开
      this.isCollapse = !this.isCollapse;
    },
    saveNavState(i1, i2) {     //保存链接的激活状态
      this.activeMenu = i1
      this.activeSubMenu = i2
    }
  },
  computed: {
    activePath() {
      return this.activeMenu === -1 || this.activeSubMenu === -1 ?
          null : this.menus[this.activeMenu].submenus[this.activeSubMenu].index
    },
    breadcrumbItem1() {
      return this.activeMenu === -1 ? null : this.menus[this.activeMenu].name
    },
    breadcrumbItem2() {
      return this.activeMenu === -1 || this.activeSubMenu === -1 ?
          null : this.menus[this.activeMenu].submenus[this.activeSubMenu].name
    }
  }
}
</script>



<style lang="less" scoped>
.home_container {
  height: 100%;
}

.el-header {
  background-color: #373D41;
  display: flex;
  justify-content: space-between;
  align-items: center;
  color: #fff;
  > div {
    display: flex;
    align-items: center;
    span {
      margin-left: 100px;
      font-size: 30px;
      font-weight: bold;
    }
  }
}

.el-aside {
  background-color: #333744;
  ele-menu {
    border-right: none;
  }
}

.el-main {
  background-color: #fff;
}

.el-menu {
  background-color: #1B1C1D;
  opacity: 1;
  height:100vh;
  font-weight: 500;
  font-size: 50px;
}

.toggle-button {
  background-color: #4A5064;
  font-size: 10px;
  line-height: 24px;
  color: #fff;
  text-align: center;
  letter-spacing: 0.2em;
  cursor: pointer;
}
</style>
