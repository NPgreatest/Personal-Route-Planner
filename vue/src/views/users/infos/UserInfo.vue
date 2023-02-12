<template>
  <div>
    <el-row>
      <el-col :span="2">
        头像：
      </el-col>
      <el-col :span="4">
        <b-avatar class="avatar" :src="info.avatar" size="9rem" ></b-avatar>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="2">
        用户名：
      </el-col>
      <el-col :span="4">
        <el-input
            placeholder="UserName"
            v-model="info.name"
            :disabled="true">
        </el-input>
      </el-col>
    </el-row>
    <el-row>


      <el-col :span="2">
        修改密码：
      </el-col>
      <el-form :rules="rules" ref="registerFormRef" :model="info" class="registerContainer">
      <el-col :span="4">
        <el-form-item prop="password">
        <el-input
            el-input placeholder="输入密码" v-model="info.password" show-password>
        </el-input>
        </el-form-item>
      </el-col>
      <el-col :span="1">
        <div class="grid-content bg-white"></div>
      </el-col>
      <el-col :span="4">
        <el-form-item prop="password">
        <el-input
            placeholder="再次输入密码"
            v-model="info.password2" show-password>
        </el-input>
        </el-form-item>
      </el-col>
      </el-form>
    </el-row>
    <el-row>
      <el-col :span="2">修改邮箱：</el-col>
      <el-col :span="4">
        <el-input
            placeholder="example@gmail.com"
            v-model="info.email">
        </el-input>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="2">修改头像：</el-col>
      <el-col :span="5">
        <el-upload  :on-change="handleelchange" :limit="1" :auto-upload="false" list-type="picture-card" ></el-upload>
      </el-col>
    </el-row>
    <el-row>
      <el-col :span="2"><div class="grid-content bg-white"></div></el-col>
      <el-col :span="3">
        <el-button type="success" round @click="upload">保存</el-button>
      </el-col>
    </el-row>
    <el-row></el-row>
    <el-row></el-row>
    <el-row></el-row>
    <el-row></el-row>
    <el-row></el-row>
    <el-row></el-row>
  </div>
</template>

<script>

import dayjs from "dayjs";

export default {
  name: "UserInfo",
  data() {
    return {
      info:{
        name:'',
        password:'',
        password2:'',
        email:'',
        avatar:'',
      },
      imgSrc:require('../../../static/bg3.jpg'),
      rules:{
        password:[{required:true,message:"请输入密码",trigger:"blur"},{ min: 6,  message: '密码长度要不小于6', trigger: 'blur' }],
        password2:[{required:true,message:"请输入密码",trigger:"blur"},{ min: 6,  message: '密码长度要不小于6', trigger: 'blur' }],
      },
    }
  },
  methods: {
    getDetails: async function() {
      const {data: res} = await this.$axios.get("/user/getinfo");
      if(res.status === 1) {
        this.info = res.data.length > 0 ? res.data[0] : this.info;
        this.info.password='';
      }
    },
    handleelchange(file) {
      var reader = new FileReader();
      reader.readAsDataURL(file.raw);
      var _this=this
      reader.onload = function(){
        console.log(this.result);
        _this.uploadava=this.result
      }
    },
    upload:async function(){
      if (this.info.password!=this.info.password2){
        this.$message.error("两次密码不一致")
        return;
      }
      const forms={
        name:  this.info.name,
        password:this.info.password,
        email:this.info.email,
        avatar: this.uploadava,
      }
      const {data: res} = await this.$axios.post("/user/update",forms);
      if (res.status==202 ||res.status==203 || res.status==0){
        this.$message.error(res.message);
        return false;
      }
      this.$message.success(res.message);
      location.reload()
    }
  },
  created() {
    this.getDetails();
    this.getMessageList()
  },

}
</script>

<style scoped>

.el-row {
  margin-top: 50px;
  margin-bottom: 20px;
 /* &:last-child {*/
 /*  margin-bottom: 0;*/
 /*}*/
}
.el-col {
  border-radius: 4px;
}
.bg-purple-dark {
  background: #99a9bf;
}
.bg-purple {
  background: #d3dce6;
}
.bg-purple-light {
  background: #e5e9f2;
}
.grid-content {
  border-radius: 4px;
  min-height: 36px;
}
.row-bg {
  padding: 10px 0;
  background-color: #f9fafc;
}
.avatar{
  size: 40rem;
}
.background{
  top:0;
  width:100%;
  height:100%;
  z-index:-1;
  position: absolute;
}
</style>
