<template>
  <div class="full">
    <nav>
      <el-row>
        <el-col :span="10" />
        <el-col :span="2">
          <router-link to="/">
            <h3>知识库构建</h3>
          </router-link>
        </el-col>
        <el-col :span="2">
          <router-link to="/about">
            <h3>应用构建</h3>
          </router-link>
        </el-col>
        <el-col :span="10">

        </el-col>
      </el-row>
    </nav>
    <router-view/>
  </div>

</template>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

nav {
  padding: 30px;
}

nav a {
  font-weight: bold;
  color: #2c3e50;
}

nav a.router-link-exact-active {
  color: #0c89fd;
}

.full{
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  /* 设置 div 的背景颜色 */
  background-color: #e8f1fd;
}
</style>
<script setup>
</script>
