<template>
  <div>
    <div class="background">
      <img :src="imgSrc" width="100%" height="100%" alt="" />
    </div>
    <el-form :rules="rules" ref="loginFormRef" :model="loginForm" class="loginContainer">
      <h3 class="loginTitle">
        系统登录
      </h3>
      <el-form-item prop="username">
        <el-input type="text" v-model="loginForm.username" placeholder="请输入用户名">
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="loginForm.password" placeholder="请输入密码">
        </el-input>
      </el-form-item>
      <el-form-item prop="primary">
      <el-button type="primary" style="width:100%" @click="submitLogin">登录</el-button>
      </el-form-item>
      <el-form-item prop="primary">
      <el-button type="primary" style="width:100%" @click="submitRegister">注册</el-button>
      </el-form-item>
    </el-form>
    </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Login",
  data(){
    return{
      captchaUrl: "",
      loginForm:{
        username:"admin",
        password:"123456",
      },
      checked: true,
      rules:{
        username:[{required:true,message:"请输入用户名",trigger:"blur"},{ min: 1, max: 20, message: '长度在 1 到 20 个字符', trigger: 'blur' }
        ],
        password:[{required:true,message:"请输入密码",trigger:"blur"},{ min: 6,  message: '密码长度要不小于6', trigger: 'blur' }],
      },
      imgSrc:require('../static/bg2.jpg')

    }
  },
  methods:{
    submitLogin(){
      this.$refs.loginFormRef.validate(async (valid) => {
        if (valid) {
          const forms={
              name:  this.loginForm.username,
              password:this.loginForm.password
          }
          const {data: res} = await this.$axios.post("/login",forms);
          if (res.status==200 ){
            this.$message.error(res.message);
            return false;
          }
          const token = res.data[0]
          const name = res.data[1]
          this.$message.success(res.message);
          window.sessionStorage.setItem("token", token);
          window.sessionStorage.setItem("userId", name);

          await this.$router.push("/home")
        } else {
          this.$message.error('登录出错请重新输入');
          return false;
        }
      });
    },
    submitRegister(){
      this.$router.push("/register")
    }
  }
};
</script>

<style lang="less" scoped>

.background{
  top:0;
  width:100%;
  height:100%;
  z-index:-1;
  position: absolute;
}


.loginContainer{

  left:30px;
  border-radius: 15px;
  background-clip: padding-box;
  margin: 180px auto;
  width: 450px;
  padding: 15px 35px 15px 35px;
  background: aliceblue;
  border:1px solid blueviolet;
  box-shadow: 0 0 25px #643965;
  z-index:1;
}
.loginTitle{
  margin: 0px auto 48px auto;
  text-align: center;
  font-size: 40px;
}
.loginRemember{
  text-align: left;
  margin: 0px 0px 15px 0px;
}



</style>

