<template>
<div class="all">
  <div class="home">
    <el-container height="100%">
      <el-aside width="120px">
        <Nav></Nav>
      </el-aside>
      <el-main>
        <chat-home></chat-home>
      </el-main>
    </el-container>
  </div>
</div>
</template>

<script>
import Nav from "../components/Nav.vue";
import ChatHome from "@/views/ChatHome";
export default {
  name: "ChatBase",
  components: {
    ChatHome,
    Nav,
  },
  data(){
    return{
      imgSrc:require('../static/background.jpg')
    }
  }
};
</script>

<style lang="scss" scoped>
.home {
  width: 90vw;
  height: 87vh;
  background-color: rgb(39, 42, 55);
  border-radius: 15px;
  position: absolute;
  left: 50%;
  top: 53%;
  transform: translate(-50%, -50%);
  z-index: 1;
}
.all {
  width: 100vw;
  height: 100vh;
  background-image: url('../static/background.jpg');
}


</style>
