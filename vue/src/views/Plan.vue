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

    <div v-for="(item, index) in selectedSites" :key="index" >
      <el-form-item>
        <span slot="label"><span style="color: #00B5AD;font-size: 24px;"> 选择第{{ index + 1 }}个景点： </span></span>
        <el-select v-model="item.id" @change="changeSite(index,item.id)" placeholder="选择感兴趣的景点" value-key="id">
          <el-option v-for="site in availableSites[index]" :label="site.sname" :value="site.sid" :key="site.sid"></el-option>
        </el-select>
        <el-button v-if="index < 4" @click="addSite(index)">添加景点</el-button>
        <el-button v-if="index > 0" @click="removeSite(index)">删除景点</el-button>
      </el-form-item>

      <el-form-item style="display: inline-block; margin-right: 10px;">
        <span slot="label" style="color: #00B5AD;font-size: 24px;"> 选择该景点活动： </span>
        <el-select v-model="item.aid" @change="changeActivity(index,item.id)" placeholder="选择感兴趣的景点" value-key="name" >
          <el-option v-for="site in activities[index]" :label="site" :value="site" :key="site"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item style="display: inline-block;width:1000px; margin-left: -200px; ">
        <div class="container">
          <div ref="textContent" class="textarea-content"></div>
          <textarea ref="textarea" v-model="activityDescription[index]" style="color: #1B1C1D;"   placeholder="活动介绍" :disabled="true"  @input="resizeTextarea"></textarea>
        </div>
      </el-form-item>
    </div>


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
import axios from "axios";
 //
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
      availableSites: [
        [],
          [],
          [],
          [],//
          []
      ],
      activityDescription: [],
      activities:[
          [''],[''],[''],[''],['']
      ],
      selectedSites: [{ id: null, aid:null }],
      maxSiteCount: 5,
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
        }
      ],
      activity:['','','','',''],
      array2:[],
      array3:[],
      key:'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',
    }
  },
  created() {
    this.getFirstSites()
  },
  methods:{
    resizeTextarea() {
      const container = this.$refs.textarea.parentNode;
      const textarea = this.$refs.textarea;

      const computedStyle = window.getComputedStyle(textarea);
      const borderTopWidth = parseFloat(computedStyle.getPropertyValue('border-top-width'));
      const borderBottomWidth = parseFloat(computedStyle.getPropertyValue('border-bottom-width'));
      const paddingTop = parseFloat(computedStyle.getPropertyValue('padding-top'));
      const paddingBottom = parseFloat(computedStyle.getPropertyValue('padding-bottom'));

      const lineHeight = parseFloat(computedStyle.getPropertyValue('line-height'));

      const contentHeight = this.$refs.textContent.offsetHeight;
      const rows = Math.ceil(contentHeight / lineHeight);

      const height = rows * lineHeight + borderTopWidth + borderBottomWidth + paddingTop + paddingBottom;
      container.style.height = `${height}px`;
      container.style.minHeight = `${height}px`;
    },
    addSite(index) {
      this.selectedSites.splice(index + 1, 0, { id: null });
    },
    removeSite(index) {
      this.selectedSites.splice(index, 1);
    },
    changeActivity:async function(index,id){
      const item = this.selectedSites[index];
      this.activity[index]=item.aid
      //console.log(this.select[index].sname,item.aid);
      axios.post('https://api.openai.com/v1/completions',
          {
            prompt: `请介绍在景点${this.select[index].sname}举办的${item.aid}活动。`,
            max_tokens: 1024,
            model: "text-davinci-003",},{
            headers: {
              'Content-Type': 'application/json',
              'Authorization': 'Bearer '+this.key
            },
          })
          .then(response => {
            const activityDescription = response.data.choices[0].text.trim()
            this.$set(this.activityDescription, index, activityDescription)
          })
          .catch(error => {
            console.log(error)
          })
    },
    changeSite:async function(index,value) {
      this.select.push()
      this.select[index]=({sid:value,pic: this.sites[value-1].pic, sname:this.sites[value-1].sname})
      // console.log(index,value)
      const {data: res} = await this.$axios.get("/user/getrecommand",{params: {sid: value}});
      const {data: res2} = await this.$axios.get("/user/getactivity",{params: {sid: value}});
      if (res.status === 1) {
        this.availableSites[index+1] = res.data.length > 0 ? res.data[0] : this.availableSites[index+1];
        //this.availableSites[index+1] =this.availableSites[index].filter(item =>item.sid!=this.select[index].sid)
      }else{
        this.$message.error("获取推荐失败，请登录")
        return
      }
      if (res2.status === 1) {
        //this.activities[index] = res2.data.length > 0 ? res2.data[0] : this.activities[index];
        this.$set(this.activities,index,res2.data[0])
        // console.log(this.activities)
        // console.log(this.availableSites)

      }else{
        this.$message.error("获取活动失败")
        return
      }
      // const selectedSiteIds = this.selectedSites.map(item => item.id);
      // const duplicateIndex = selectedSiteIds.findIndex(
      //     (id, i) => id === this.selectedSites[index].id && i !== index
      // );
      // if (duplicateIndex !== -1) {
      //   this.selectedSites[duplicateIndex].id = null;
      //   this.$message({
      //     message: "不能选择重复的景点",
      //     type: "warning"
      //   });
      // }
    },
    selected(id){
      this.selectid=id;
      localStorage.setItem('tags',id);
    },
    finish(){
      var temp=[];
      temp.push(this.input)
      for(var i=0;i<this.select.length;i++){
        temp.push(this.select[i].sname)
      }
      localStorage.setItem('routes', JSON.stringify(temp))
      localStorage.setItem('activites',JSON.stringify(this.activity))
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
          this.availableSites[0].push(this.sites[item])
        }
      } else {
      this.$message.error("获取景点库失败")
      return
      }
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
.background {
  width: 100%;
  height: 100%;
  z-index: -1;
  position: fixed;
  top: 0;
  left: 0;
  background-size: 100% 100%;
}
.elForm{
  top: 10%;
  left: 25%;
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
.container {
  width: 100%;
  height: auto;
  min-height: 120px;
  padding: 3px;
  font-size: 16px;
  border-radius: 4px;
  border: 1px solid #ccc;
  box-sizing: border-box;
  line-height: 1; /* 将行高设置为字体大小的倍数，例如 1 或 1.2 */
}

.textarea-content {
  position: absolute;
  visibility: hidden;
  white-space: pre-wrap;
  word-wrap: break-word;
}

textarea {
  width: 100%;
  height: auto;
  min-height: 120px;
  border: none;
  resize: none;
  outline: none;
  font-size: inherit;
  color: #1B1C1D;
  font-weight: bold;
}
textarea::placeholder {
  color: #1B1C1D;
  font-weight: bold;
}
</style>
