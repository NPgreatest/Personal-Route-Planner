<template>
  <div>
    <div class="background">
      <img :src="imgSrc" width="100%" height="100%" alt="" />
    </div>
    <el-form :rules="rules" ref="registerFormRef" :model="registerForm" class="registerContainer">
      <h3 class="registerTitle">
        注册账户
      </h3>
      <el-form-item prop="username">
        <el-input type="text" v-model="registerForm.username" placeholder="请输入昵称" >
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="registerForm.password" placeholder="请输入密码" >
        </el-input>
      </el-form-item>
      <el-form-item prop="password2">
        <el-input type="password" v-model="registerForm.password2" placeholder="请重复密码" >
        </el-input>
      </el-form-item>
      <el-form-item prop="mail">
        <el-input type="mail" v-model="registerForm.mail" placeholder="请输入邮箱" >
        </el-input>
      </el-form-item>
      <div class="ava">头像(必须):</div>
      <el-upload  :on-change="handleelchange" :limit="1" :auto-upload="false" list-type="picture-card" >
        <i class="el-icon-plus"></i>
      </el-upload>
      <el-button type="primary" style="width:100%" @click="submitRegister">注册</el-button>
    </el-form>
  </div>
</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Register",
  data(){
    return{
      checked: true,
      registerForm:{
        username: '',
        password:'',
        password2: '',
        mail:'',
        avatar:'',
      },
      rules:{
        username:[{required:true,message:"请输入昵称",trigger:"blur"},{ min: 1, max: 14, message: '长度在 1 到 14 个字符', trigger: 'blur' }
        ],
        password:[{required:true,message:"请输入密码",trigger:"blur"},{ min: 6,  message: '密码长度要不小于6', trigger: 'blur' }],
        password2:[{required:true,message:"请输入密码",trigger:"blur"},{ min: 6,  message: '密码长度要不小于6', trigger: 'blur' }],
        mail:[{require: true, trigger: "blur"},{min:1,message:'邮箱不能为空',trigger: "blur"}],
        avatar:[{require: true, trigger: "blur"},{min:1,message:'不能为空',trigger: "blur"}],
      },
      imgSrc:require('../static/bg1.jpg'),
      uploadava:'',
    }
  },
  methods:{
    handleelchange(file) {
      var reader = new FileReader();
      reader.readAsDataURL(file.raw);
      var _this=this
      reader.onload = function(){
        console.log(this.result);
        _this.uploadava=this.result

      }
      //axios.post("接口地址", formdata).then(res => { console.log(res) })
    },
    submitRegister(){
      console.log(this.uploadava)
      this.$refs.registerFormRef.validate(async (valid) => {
        if (valid) {
          if (this.registerForm.password!==this.registerForm.password2){
              this.$message.error("两次密码不一致");
              return false;
          }


          const forms={
            name:  this.registerForm.username,
            password:this.registerForm.password,
            email:this.registerForm.mail,
            avatar: this.uploadava,

          }
          const {data: res} = await this.$axios.post("/home/register",forms);
          if (res.status==202 ||res.status==203){
            this.$message.error(res.message);
            return false;
          }
          this.$message.success(res.message);
          await this.$router.push("/home")
        } else {
          this.$message.error('注册出错请重新输入');
          return false;
        }
      });
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


.registerContainer{

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
.registerTitle{
  margin: 0px auto 48px auto;
  text-align: center;
  font-size: 40px;
}
.ava{
  margin: 0px auto ;
  text-align: center;
  font-size: 20px;
}
.registerRemember{
  text-align: left;
  margin: 0px 0px 15px 0px;
}



</style>
