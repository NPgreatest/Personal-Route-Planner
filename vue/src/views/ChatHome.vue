<template>
  <div class="chatHome">
    <div class="chatLeft">
      <div class="title">
        <h1>景点咨询</h1>
      </div>
      <div class="online-person">
        <span class="onlin-text">可选景点</span>
        <div class="person-cards-wrapper">
          <div
              class="personList"
              v-for="personInfo in personList"
              :key="personInfo.sid"
              @click="clickPerson(personInfo)"
          >
            <PersonCard
                :personInfo="personInfo"
                :pcCurrent="pcCurrent"
            ></PersonCard>
          </div>
        </div>
      </div>
    </div>
    <div class="chatRight">
      <div v-if="showChatWindow">
        <ChatWindow
            :frinedInfo="chatWindowInfo"
            @personCardSort="personCardSort"
        ></ChatWindow>
      </div>
      <div class="showIcon" v-else>
        <span class="iconfont icon-snapchat"></span>
      </div>
    </div>
    <!-- <el-col :span="4"><div class="grid-content bg-purple"></div></el-col> -->
  </div>
</template>

<script>
import PersonCard from "../components/PersonCard.vue";
import ChatWindow from "../views/Chat.vue";

 // import { getFriend } from "../utils/utils.js"
export default {
  name: "ChatHome",
  props:['tags'],
  components: {
    PersonCard,
    ChatWindow,
  },
  data() {
    return {
      pcCurrent: "",
      personList: [
        {sid:1,
          pic:require("../static/logo.png")},
        {sid:2,
          pic:require("../static/logo.png")},
        {sid:3,
          pic:require("../static/logo.png")}
      ],
      showChatWindow: false,
      chatWindowInfo: {},
    };
  },
  watch:{
    tags:{
      immediate:true,
      handler(tags) {
        console.log({tags})
        this.updateSites(tags);
      }
    }
  },
  mounted() {
    this.getSites();
  },
  methods: {
    updateSites:async function(v){
      const {data: res} = await this.$axios.get("/home/sitesbytags", {params: {tagid: v}});
      console.log(res)
      if(res.status === 1) {
        this.personList = res.data.length > 0 ? res.data[0] : this.personList;
        // for(var i=0;i<this.personList.length;i++){
        //   this.personList[i].short=this.personList[i].description.substr(0,11)+"..."
        // }
        console.log(this.personList)
      } else {
        this.$message.error("获取景点失败，请重试")
      }
    },
    getSites:async function(){
        const {data: res} = await this.$axios.get("/home/allsites");
        if(res.status === 1) {
          this.personList = res.data.length > 0 ? res.data[0] : this.personList;
          for(var i=0;i<this.personList.length;i++){
            this.personList[i].short=this.personList[i].description.substr(0,11)+"..."
          }
        }else{
          this.$message.error("获取景点失败")
        }

    },
    clickPerson(info) {
      this.showChatWindow = true;
      this.chatWindowInfo = info;
      this.personInfo = info;
      this.pcCurrent = info.sid;
    },
    personCardSort(id) {
      if (id !== this.personList[0].id) {
        console.log(id);
        let nowPersonInfo;
        for (let i = 0; i < this.personList.length; i++) {
          if (this.personList[i].id == id) {
            nowPersonInfo = this.personList[i];
            this.personList.splice(i, 1);
            break;
          }
        }
        this.personList.unshift(nowPersonInfo);
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.chatHome {
  // margin-top: 20px;
  display: flex;
  .chatLeft {
    width: 320px;
    .title {
      color: #fff;
      padding-left: 10px;
    }
    .online-person {
      margin-top: 50px;
      .onlin-text {
        padding-left: 10px;
        color: rgb(176, 178, 189);
      }
      .person-cards-wrapper {
        padding-left: 10px;
        height: 65vh;
        margin-top: 20px;
        overflow: hidden;
        overflow-y: scroll;
        box-sizing: border-box;
        &::-webkit-scrollbar {
          width: 0; /* Safari,Chrome 隐藏滚动条 */
          height: 0; /* Safari,Chrome 隐藏滚动条 */
          display: none; /* 移动端、pad 上Safari，Chrome，隐藏滚动条 */
        }
      }
    }
  }

  .chatRight {
    flex: 1;
    padding-right: 30px;
    .showIcon {
      position: absolute;
      top: calc(50% - 150px); /*垂直居中 */
      left: calc(50% - 50px); /*水平居中 */
      .icon-snapchat {
        width: 300px;
        height: 300px;
        font-size: 300px;
        // color: rgb(28, 30, 44);
      }
    }
  }
}
</style>
