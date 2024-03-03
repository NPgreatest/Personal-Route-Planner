<template>

  <div class="chat-window">

    <div class="top">
      <div class="head-pic">
        <HeadPortrait :imgUrl="frinedInfo.pic"></HeadPortrait>
      </div>
      <div class="info-detail">
        <div class="name">{{ frinedInfo.sname }}</div>
        <div class="detail">{{ frinedInfo.description }}</div>
      </div>
    </div>
    <div class="botoom">
      <div class="chat-content" ref="chatContent">
        <div class="chat-wrapper" v-for="item in chatList" :key="item.id">
          <div class="chat-friend" v-if="item.uid !== '1001'">
            <div class="chat-text" v-if="item.chatType == 0">
              {{ item.msg }}
            </div>
            <div class="chat-img" v-if="item.chatType == 1">
              <img
                  :src="item.msg"
                  alt="表情"
                  v-if="item.extend.imgType == 1"
                  style="width: 100px; height: 100px"
              />
              <el-image :src="item.msg" :preview-src-list="srcImgList" v-else>
              </el-image>
            </div>
            <div class="chat-img" v-if="item.chatType == 2">
              <div class="word-file">
                <FileCard
                    :fileType="item.extend.fileType"
                    :file="item.msg"
                ></FileCard>
              </div>
            </div>
            <div class="info-time">
              <img :src="item.headImg" alt="" />
              <span>{{ item.name }}</span>
              <span>{{ item.time }}</span>
            </div>
          </div>
          <div class="chat-me" v-else>
            <div class="chat-text" v-if="item.chatType == 0">
              {{ item.msg }}
            </div>
            <div class="chat-img" v-if="item.chatType == 1">
              <img
                  :src="item.msg"
                  alt="表情"
                  v-if="item.extend.imgType == 1"
                  style="width: 100px; height: 100px"
              />
              <el-image
                  style="max-width: 300px; border-radius: 10px"
                  :src="item.msg"
                  :preview-src-list="srcImgList"
                  v-else
              >
              </el-image>
            </div>
            <div class="chat-img" v-if="item.chatType == 2">
              <div class="word-file">
                <FileCard
                    :fileType="item.extend.fileType"
                    :file="item.msg"
                ></FileCard>
              </div>
            </div>
            <div class="info-time">
              <span>{{ item.name }}</span>
              <span>{{ item.time }}</span>
              <img :src="item.headImg" alt="" />
            </div>
          </div>
        </div>
      </div>
      <div class="chatInputs">
        <div class="emoji boxinput" @click="clickEmoji">
<!--          <img src="@/assets/img/emoji/smiling-face.png" alt="" />-->
        </div>
        <div class="emoji-content">
          <Emoji
              v-show="showEmoji"
              @sendEmoji="sendEmoji"
              @closeEmoji="clickEmoji"
          ></Emoji>
        </div>
        <input class="inputs" v-model="inputMsg" @keyup.enter="sendText" />
        <div class="send boxinput" @click="sendText">
          <img src="../static/rocket.png" alt="" />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import {animation} from"../utils/utils";
import HeadPortrait from "../components/HeadPortrait.vue";
import Emoji from "../components/Emoji.vue";
import FileCard from"../components/FileCard.vue";
export default {
  name: "Chat",
  components: {
    HeadPortrait,
    Emoji,
    FileCard,
  },
  props: {
    frinedInfo: Object,
    default() {
      return {};
    },
  },
  watch: {
    frinedInfo() {
      this.getFriendChatMsg();
    },
  },
  data(){
    return{
      chatList: [],
      inputMsg: "",
      showEmoji: false,
      friendInfo: {},
      srcImgList: [],
      // imgSrc:require('../static/background.jpg')
    }
  },
  mounted() {
    this.getFriendChatMsg();
  },
  methods: {
    //获取聊天记录
    getFriendChatMsg() {

      this.scrollBottom();
    },
    //发送信息
    sendMsg(msgList) {
      this.chatList.push(msgList);
      this.scrollBottom();
    },
    //获取窗口高度并滚动至最底层
    scrollBottom() {
      this.$nextTick(() => {
        const scrollDom = this.$refs.chatContent;
        animation(scrollDom, scrollDom.scrollHeight - scrollDom.offsetHeight);
      });
    },
    //关闭标签框
    clickEmoji() {
      this.showEmoji = !this.showEmoji;
    },
    //发送文字信息
    sendText:async function(){
      if (this.inputMsg) {
        let chatMsg = {
          headImg: require("../static/thinking-face.png"),
          name: "admin",
          time: "09：12 AM",
          msg: this.inputMsg,
          chatType: 0, //信息类型，0文字，1图片
          uid: "1001", //uid
        };
        this.sendMsg(chatMsg);
        const {data: res} = await this.$axios.get("/user/getsitegpt",{params: {sid: this.frinedInfo.sid,detail:this.inputMsg}});
        if(res.status === 1) {
          var str = res.data.length > 0 ? res.data[0] : "";
          let chatMsg2 = {
            headImg: this.frinedInfo.pic,
            name: "AI小助手",
            time: new Date(),
            msg: str,
            chatType: 0, //信息类型，0文字，1图片
            uid: "1002", //uid
          };
          this.sendMsg(chatMsg2);
        }else{
          this.$message.error("获取回答失败")
        }

        this.inputMsg = "";
      } else {
        this.$message({
          message: "消息不能为空哦~",
          type: "warning",
        });
      }
    },
    //发送表情
    sendEmoji(msg) {
      let chatMsg = {
        headImg: require("../static/thinking-face.png"),
        name: "admin",
        time: new Date(),
        msg: msg,
        chatType: 1, //信息类型，0文字，1图片
        extend: {
          imgType: 1, //(1表情，2本地图片)
        },
        uid: "1001",
      };
      this.sendMsg(chatMsg);
      this.clickEmoji();
    },
  },
}
</script>

<style lang="scss" scoped>
.chat-window {
  width: 100%;
  height: 100%;
  margin-left: 20px;
  position: relative;

.top {
  margin-bottom: 50px;
&::after {
   content: "";
   display: block;
   clear: both;
 }
.head-pic {
  float: left;
}
.info-detail {
  float: left;
  margin: 5px 20px 0;
.name {
  font-size: 20px;
  font-weight: 600;
  color: #fff;
}
.detail {
  color: #9e9e9e;
  font-size: 12px;
  margin-top: 2px;
  position: absolute;
}
}
.other-fun {
  float: right;
  margin-top: 20px;
span {
  margin-left: 30px;
  cursor: pointer;
}
// .icon-tupian {

   // }
input {
  display: none;
}
}
}
.botoom {
  width: 100%;
  height: 70vh;
  background-color: rgb(50, 54, 68);
  border-radius: 20px;
  padding: 20px;
  box-sizing: border-box;
  position: relative;
.chat-content {
  width: 100%;
  height: 85%;
  overflow-y: scroll;
  padding: 20px;
  box-sizing: border-box;
&::-webkit-scrollbar {
   width: 0; /* Safari,Chrome 隐藏滚动条 */
   height: 0; /* Safari,Chrome 隐藏滚动条 */
   display: none; /* 移动端、pad 上Safari，Chrome，隐藏滚动条 */
 }
.chat-wrapper {
  position: relative;
  word-break: break-all;
.chat-friend {
  width: 100%;
  float: left;
  margin-bottom: 20px;
  display: flex;
  flex-direction: column;
  justify-content: flex-start;
  align-items: flex-start;
.chat-text {
  max-width: 90%;
  padding: 20px;
  border-radius: 20px 20px 20px 5px;
  background-color: rgb(56, 60, 75);
  color: #fff;
&:hover {
   background-color: rgb(39, 42, 55);
 }
}
.chat-img {
img {
  width: 100px;
  height: 100px;
}
}
.info-time {
  margin: 10px 0;
  color: #fff;
  font-size: 14px;
img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  vertical-align: middle;
  margin-right: 10px;
}
span:last-child {
  color: rgb(101, 104, 115);
  margin-left: 10px;
  vertical-align: middle;
}
}
}
.chat-me {
  width: 100%;
  float: right;
  margin-bottom: 20px;
  position: relative;
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  align-items: flex-end;
.chat-text {
  float: right;
  max-width: 90%;
  padding: 20px;
  border-radius: 20px 20px 5px 20px;
  background-color: rgb(29, 144, 245);
  color: #fff;
&:hover {
   background-color: rgb(26, 129, 219);
 }
}
.chat-img {
img {
  max-width: 300px;
  max-height: 200px;
  border-radius: 10px;
}
}
.info-time {
  margin: 10px 0;
  color: #fff;
  font-size: 14px;
  display: flex;
  justify-content: flex-end;

img {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  vertical-align: middle;
  margin-left: 10px;
}
span {
  line-height: 30px;
}
span:first-child {
  color: rgb(101, 104, 115);
  margin-right: 10px;
  vertical-align: middle;
}
}
}
}
}
.chatInputs {
  width: 90%;
  position: absolute;
  bottom: 0;
  margin: 3%;
  display: flex;
.boxinput {
  width: 50px;
  height: 50px;
  background-color: rgb(66, 70, 86);
  border-radius: 15px;
  border: 1px solid rgb(80, 85, 103);
  position: relative;
  cursor: pointer;
img {
  width: 30px;
  height: 30px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);
}
}
.emoji {
  transition: 0.3s;
&:hover {
   background-color: rgb(46, 49, 61);
   border: 1px solid rgb(71, 73, 82);
 }
}

.inputs {
  width: 90%;
  height: 50px;
  background-color: rgb(66, 70, 86);
  border-radius: 15px;
  border: 2px solid rgb(34, 135, 225);
  padding: 10px;
  box-sizing: border-box;
  transition: 0.2s;
  font-size: 20px;
  color: #fff;
  font-weight: 100;
  margin: 0 20px;
&:focus {
   outline: none;
 }
}
.send {
  background-color: rgb(29, 144, 245);
  border: 0;
  transition: 0.3s;
  box-shadow: 0px 0px 5px 0px rgba(0, 136, 255);
&:hover {
   box-shadow: 0px 0px 10px 0px rgba(0, 136, 255);
 }
}
}
}
}
</style>
