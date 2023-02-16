<template>
  <div class="nav">
    <div class="nav-menu-wrapper">
      <ul class="menu-list">
        <li
          v-for="(item, index) in tags"
          :key="index"
          :class="{ activeNav: index == current }"
          @click="changeMenu(index)"
        >
          <div class="block"></div>
          <el-button style="left: -33px; " type="text" @click="change(item.tagid)">{{item.name}}</el-button>
<!--          <b-avatar class="comment-avatar" :src="item.pic" size="2rem" style="left: -33px; "></b-avatar>-->
        </li>
      </ul>
    </div>
    <div class="own-pic">
        <HeadPortrait :imgUrl="userInfo.avatar" ></HeadPortrait>
    </div>
  </div>
</template>

<script>
import HeadPortrait from "./HeadPortrait.vue";

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name:"Nav",
  props:['final_tag'],
  components: {
    HeadPortrait,
  },
  data() {
    return {
      userInfo:{name:"",avatar:""},
      tags:[
        {tagid:1,name:"红色团日"}
      ],
      menuList: [
        {
          pic:require("../static/logo.png"),
        },
        {
          pic:require("../static/logo.png"),
        }
      ],
      current: 1,
      imgUrl: require('../static/bg1.jpg')
    };
  },
  created(){
    this.ava();
  },
  methods: {
    change:async function(value){
      this.final_tag(value);
    },
    ava:async function(){
      const {data: res} = await this.$axios.get("/user/getinfo");
      if(res.status === 1) {
        this.userInfo= res.data.length > 0 ? res.data[0] : this.userInfo;
        console.log(this.userInfo)
      }else{
        this.$message.error("获取头像失败")
      }
      const {data: res2} = await this.$axios.get("/home/alltags");
      if(res2.status === 1) {
        this.tags = res.data.length > 0 ? res2.data[0] : this.tags;
      }
    },
    changeMenu(index) {
      this.current = index;
    },
  },
};
</script>

<style lang="scss" scoped>
.nav {
  width: 100%;
  height: 90vh;
  position: relative;
  border-radius: 20px 0 0 20px;
  .nav-menu-wrapper {
    position: absolute;
    top: 40%;
    transform: translate(0, -50%);
    .menu-list {
      margin-left: 10px;

      li {
        margin: 40px 0 0 0;
        list-style: none;
        cursor: pointer;
        position: relative;
        .block {
          background-color: rgb(29, 144, 245);
          position: absolute;
          left: -50px;
          width: 10px;
          height: 35px;
          transition: 0.5s;
          border-top-right-radius: 5px;
          border-bottom-right-radius: 4px;
          opacity: 0;
        }
        &:hover {
          span {
            color: rgb(29, 144, 245);
          }
          .block {
            opacity: 1;
          }
        }
      }
    }
  }
  .own-pic {
    position: absolute;
    bottom: 10%;
    margin-left: 25px;
  }
}
.activeNav {
  span {
    color: rgb(29, 144, 245);
  }
  .block {
    opacity: 1 !important;
  }
}
</style>
