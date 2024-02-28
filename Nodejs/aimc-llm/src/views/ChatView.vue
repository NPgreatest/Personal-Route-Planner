<template>
  <div>
    <el-container>
      <el-aside width="20%">
      </el-aside>
      <el-main style="width: 60%;margin-right: 100px;">
        <div class="chat-container" style="width: 100%;">
          <div class="chat-messages">
            <div v-for="(message, index) in messages" :key="index"
              :class="{ 'message-right': message.fromMe, 'message-left': !message.fromMe }">
              <span class="username">{{ message.username }}：</span>
              <span>{{ message.content }}</span>
            </div>
          </div>
          <div class="chat-input">
            <input v-model="data.question" @keyup.enter="sendMessage" placeholder="输入消息..." />
            <button @click="sendMessage">发送</button>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<script>
import { chat, chatmessages, getChatStream } from '@/api/chat'
import axios from 'axios';


export default {
  data() {
    return {
      // 消息列表
      messages: [],
      eventSource: null,
      data: {
        "user_id": 123,
        "user_name": "我",
        "question": "你好",
        "model": "chatgpt"
      },
      // 发送消息
      // newMessage: '',
      outputMessage: '好的'
    }
  },
  methods: {
    sendMessage() {
      this.messages.push({ username: this.data.user_name, content: this.data.question, fromMe: true });
      // console.log(this.data.question)
      if (this.data.question) {
        getChatStream(this.data).then(response => {
          const reader = response.body.getReader();
          console.log(reader)
          const textDecoder = new TextDecoder();
          this.messages.push({ username: this.data.user_name, content: '', fromMe: false });
          this.outputMessage = ''
          const getStream = (reader) => {
            return reader.read().then((result) => {
              if (result.done) {
                console.log('Stream ended')
                return
              }
              const chunkText = textDecoder.decode(result.value);
              console.log('Received chunk:', chunkText);
              // this.outputMessage += chunkText;
              // this.outputMessage += JSON.parse(chunkText).answer;
              this.outputMessage +=chunkText;
              this.messages[this.messages.length-1].content = this.outputMessage
              return getStream(reader);
            })
          }
          return getStream(reader);
        })
      }
    }
  }

}
</script>

<style scoped>
.chat-container {
  width: 400px;
  margin: 0 auto;
}

.chat-messages {
  height: 300px;
  border: 1px solid #ccc;
  padding: 10px;
  overflow-y: scroll;
}

.chat-input {
  margin-top: 10px;
}

.message {
  margin-bottom: 10px;
  padding: 5px;
}

.message-right {
  text-align: right;
}

.message-left {
  text-align: left;
}

.username {
  font-weight: bold;
}
</style>
