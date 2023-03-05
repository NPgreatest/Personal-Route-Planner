<template>
<div style="width: 240px">
  <div class="background">
    <img :src="imgSrc" width="100%" height="100%" alt="" />
  </div>




<div class="elForm">



  <el-form ref="form" :model="form" label-width="220px" style="margin: auto">
    <el-row  class="tag-area" :gutter="20" style="margin: auto; display: flex;flex-wrap: wrap; padding-top: 20px; padding-bottom: 20px "    >
    <li :key="item.tagid" v-for="item in tags">
      <el-col :span="4">
        <el-button type="success" v-if="selectid==item.tagid" @click="selected(item.tagid)"  class="buttons" >{{item.name}}</el-button>
        <el-button type="primary" v-if="selectid!=item.tagid" @click="selected(item.tagid)"  class="buttons" >{{item.name}}</el-button>
      </el-col>
    </li>
    </el-row>

    <el-form-item  >
      <span slot="label"><span style="color: #00B5AD;font-size: 24px;"> 请输入起点： </span></span>
      <el-input placeholder="起点" v-model="input" style="width: 200px"></el-input>
    </el-form-item>


    <el-form-item  >
      <span slot="label"><span style="color: #00B5AD;font-size: 24px;"> 选择第一个景点： </span></span>
      <el-select v-model="a1" @change="changeSite1" placeholder="选择感兴趣的景点" value-key="id"  :disabled="dis1">
        <el-option v-for="item in array1" :label="item.sname" :value="item.sid" :key="item.sid" ></el-option>
      </el-select>

    </el-form-item>


    <el-form-item  >
      <span slot="label"><span style="color: #00B5AD;font-size: 24px;"> 选择第二个景点： </span></span>
      <el-select v-model="a2" @change="changeSite2" placeholder="选择感兴趣的景点" value-key="id" :disabled="dis2">
        <el-option v-for="item in array2" :label="item.sname" :value="item.sid" :key="item.sid" ></el-option>
      </el-select>

    </el-form-item>

    <el-form-item  >
      <span slot="label"><span style="color: #00B5AD;font-size: 24px;"> 选择第三个景点： </span></span>
      <el-select v-model="a3" @change="changeSite3" placeholder="选择感兴趣的景点" value-key="id" :disabled="dis3">
        <el-option v-for="item in array3" :label="item.sname" :value="item.sid" :key="item.sid" ></el-option>
      </el-select>

    </el-form-item>


    <b-avatar  v-for="item in select"   class="comment-avatar" :src="item.pic" size="15rem" ></b-avatar>
    <el-form-item>
      <el-button type="primary" @click="repick">重置</el-button>
      <el-button type="success" @click="finish">完成</el-button>
    </el-form-item>



  </el-form>

</div>

</div>








</template>

<script>
export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "Plan",
  data(){
    return{
      imgSrc:require('../static/background.jpg'),
      avatar:"http://localhost:8080/images/3.jpg",
      a1:"",
      a2:"",
      a3:"",
      input:"",
      selectid:0,
      dis1:false,
      dis2:false,
      dis3:false,
      tags: [{
        id: 1,
        name: "前端",
        count: 5
      },
        {
          id: 2,
          name: "后端",
          count: 8
        },
        {
          id: 3,
          name: "SpringBoot",
          count: 5
        },
        {
          id:4,
          name:"123",
          count:5
        }
      ],
      array1:[
        {sid:1,
          sname:"景点",
          pic:"http://localhost:8080/images/1.jpg",
          rate:999,
          tag:"标签"
        },
        {sid:2,
          sname:"景点",
          pic:"http://localhost:8080/images/2.jpg",
          rate:999,
          tag:"标签"
        },
        {sid:3,
          sname:"景点",
          pic:"http://localhost:8080/images/3.jpg",
          rate:999,
          tag:"标签"
        },
      ],
      select:[
        {sid:1,
          pic:"http://localhost:8080/images/1.jpg"
        },
      ],
      array2:[],
      array3:[],
    }
  },
  created() {
    this.getFirstSites()
  },
  methods:{
    selected(id){
      this.selectid=id;
    },
    finish(){
      var temp=[];
      temp.push(this.input)
      temp.push(this.array1[0].sname)
      temp.push(this.array2[0].sname)
      temp.push(this.array3[0].sname)
      localStorage.setItem('routes', JSON.stringify(temp))
      this.$router.push("/route")
    },
    getFirstSites:async function() {
      const {data: res2} = await this.$axios.get("/home/alltags");
      if(res2.status === 1) {
        this.tags = res2.data.length > 0 ? res2.data[0] : this.tags;
      }
      const {data: res} = await this.$axios.get("/home/allsites");
      if (res.status === 1) {
        this.sites = res.data.length > 0 ? res.data[0] : this.site;
        var s1=this.randomids()
        this.array1=[]
        for(var item of s1){
          this.array1.push(this.sites[item])
        }
      } else {
      this.$message.error("获取景点库失败")
      return
      }
    },
    changeSite1:async function(value){
      this.select=[]
      this.select.push({sid:value,pic: this.sites[value-1].pic})
      this.dis1=true
      const {data: res} = await this.$axios.get("/user/getrecommand",{params: {sid: this.select[0].sid}});
      if (res.status === 1) {
        this.array2 = res.data.length > 0 ? res.data[0] : this.array2;
        this.array2 =this.array2.filter(item =>item.sid!=this.select[0].sid)
      }else{
        this.$message.error("获取推荐失败，请登录")
        return
      }
    },
    changeSite2:async function(value){
      this.dis2=true
      this.select.push({sid:value,pic: this.sites[value-1].pic})
      const {data: res} = await this.$axios.get("/user/getrecommand",{params: {sid: this.select[1].sid}});
      if (res.status === 1) {
        this.array3 = res.data.length > 0 ? res.data[0] : this.array3;
      }else{
        this.$message.error("获取推荐失败，请登录")
        return
      }
    },
    changeSite3:async function(value){
      this.dis3=true
      this.select.push({sid:value,pic: this.sites[value-1].pic})
    },
    repick(){
      this.dis1=false
      this.dis2=false
      this.dis3=false
      this.select=[]
      this.a1=""
      this.a2=""
      this.a3=""
      this.getFirstSites();
    },
    randomids(){
      var t=[];
      var flag=false;
      for(var i=0;i<5;i++){
        var idx=Math.floor(Math.random()*this.sites.length)
        flag=false;
        for(var k=0;k<t.length;k++){
          if(t[k]==idx){
            flag=true;
          }
        }
        switch (flag){
          case false:t.push(idx);
            break;
          case true:i-=1;
            break;
        }
      }
      return t;
    }
  }
}
</script>

<style scoped>
.background{

  width:100%;
  height:100%;
  z-index:-1;
  position: absolute;
  background-size: 100% 100%;
}
.elForm{
  top: 20%;
  left: 35%;
  margin: auto;
  display: flex;
  justify-content: center;
  position: absolute;
  z-index: 1;
}
.tag-area {
  width: 840px;
  margin: 0 auto;
  overflow: hidden;
}

.tag-area li {
  list-style: none;
  float: left;
}
.buttons:hover {
  transform:  scale(1.22);
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
}
</style>
