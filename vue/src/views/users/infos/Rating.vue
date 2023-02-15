<template>
  <div>
    <!-- 卡片视图区域 -->
    <el-card class="box-card">
      <!-- 搜索栏 -->
      <el-row :gutter="15">
        <el-col :span="6">
          <el-select  v-model="selectedSites"   @change="changeSites" clearable value-key="id"  placeholder="景点">
            <el-option v-for="item in sites" :label="item.sname" :key="item.id" :value="item.id"></el-option>
          </el-select>
        </el-col>

        <el-col :span="4">

          <el-select v-model="test"  @change="changeType"   placeholder="分类" value-key="id" >
            <el-option v-for="item in types" :label="item.name" :value="item.id" :key="item.id" ></el-option>
          </el-select>
        </el-col>

        <el-col :span="2">
          <!-- 推荐选项 -->
          <el-checkbox v-model="queryInfo.recommended" label="未评分" border></el-checkbox>
        </el-col>

        <el-col :span="4">
          <el-button @click="add" type="success">添加</el-button>
          <el-button type="primary" @click="getBlogList" icon="el-icon-search">搜索</el-button>
        </el-col>
      </el-row>

      <!-- 博客列表区域 -->
      <el-table :data="blogs" border stripe>
        <el-table-column label="id" prop="sid" width="150"></el-table-column>
        <el-table-column label="名称" prop="name"></el-table-column>
        <el-table-column label="评分" prop="rate" width="550">
          <template v-slot="scope">
            <div class="block">
              <span class="demonstration"></span>
              <el-slider v-model="blogs[scope.$index].rate"></el-slider>
            </div>
          </template>
        </el-table-column>
        <el-table-column label="操作"  width="150">
          <template v-slot="scope">
            <el-button size="mini" @click="handleEdit(blogs[scope.$index].sid,scope.$index)">保存</el-button>
            <el-button size="mini" type="danger" @click="handleDelete(blogs[scope.$index].sid)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>

      <!-- 分页区域 -->
      <el-pagination @size-change="handleSizeChange" @current-change="handleCurrentChange"
                     :current-page="queryInfo.pageNum" :page-sizes="[8, 10, 12, 15]" :page-size="queryInfo.pageSize"
                     layout="total, sizes, prev, pager, next, jumper" :total="total">
      </el-pagination>
    </el-card>

  </div>

</template>

<script>
export default {
  name: "Rating",
  data() {
    return {
      types: [],
      test: "",
      sid:"",
      sname:"",
      sites:[],
      selectedID:"",
      selectedSites: "",
      blogs: [],
      total: 2,
      queryInfo: {          //获取博客列表的数据
        pageNum: 1,
        pageSize: 8,
        blogTitle: "",
        typeId: 2,
        recommended: false,
      }
    }
  },
  created() {
    this.getBlogList();
    this.getAllTags();
  },
  methods: {
    clear: function() {
      this.queryInfo.blogTitle = "123";
      this.queryInfo.typeId = 0;
      this.queryInfo.recommended = false;
      this.selectedType = "";
    },
    changeSites(value){
      this.sid=value
      let obj = {};
      obj = this.options.find((item)=>{
        return item.value === value;
      });
      this.sname=obj.label
      alert(this.sname)
    },
    getBlogList: async function() {
      const{data: res} = await this.$axios.get("/user/getrate",{params: this.queryInfo});
        this.blogs=res.data[0].rates
        this.getNames();
    },
    getNames: async function(){
      for(var key in this.blogs){
        const {data: tres} = await this.$axios.get("/home/sitedetails", {params: {siteid: this.blogs[key].sid}});
        if(tres.status === 1) {
          this.blogs[key]['name']=tres.data[0].sname;
        }else{
          this.$message.error("查询失败,请重新查询");
        }
      }
      let dataArr=JSON.stringify(this.blogs)
      this.blogs=JSON.parse(dataArr)
    },
    getAllTags: async function(){
      const{data: res} = await this.$axios.get("/home/alltags");
      if(res.status === 1) {
        this.types = res.data.length > 0 ? res.data[0] : [];
      } else {
        this.$message.error(res.message);
      }
      for(var i=0;i<this.types.length;i++){
        this.types[i].id=this.types[i].tagid
      }
    },
    changeType:async function(value) {
      const{data: res} = await this.$axios.get("/home/sitesbytags",{params: {tagid:value}});
      if(res.status === 1) {
        this.sites=res.data.length > 0 ? res.data[0] : [];
        for(var i=0;i<this.sites.length;i++){
          this.sites[i].id=this.sites[i].sid
        }
      }else{
        this.$message.error(res.message);
      }
    },
    add: async function(){
        for(var i=0;i<this.blogs.length;i++){
          if(this.blogs[i].sid==this.sid){
            this.$message.error("已经添加该景点");
            return
          }
        }
        this.blogs.push({sid:this.sid,name:""});
        this.getNames();
    },
    handleEdit: function(id,index) {
      console.log(id,index)
      this.uploadrate()
    },
    handleDelete: async function(sid){
      for(var i=0;i<this.blogs.length;i++){
        if(this.blogs[i].sid==sid){
          console.log(this.blogs)
          this.blogs = this.blogs.filter(item => item.sid !=sid)
              console.log(this.blogs)
        }
      }
      this.uploadrate()
    },
    uploadrate: async function(){
      var temp={name:'admin',rates:[]}
      temp.rates=this.blogs//12312
      const {data: res} = await this.$axios.post("/user/rate",temp);
      if (res.status==200 ){
        this.$message.error(res.message);
        return false;
      }
      this.$message.success(res.message);
    },
    handleSizeChange: function(pagesize) {     // 监听pagesize 改变的事件
      this.queryInfo.pageSize = pagesize;
      this.getBlogList();
    },
    handleCurrentChange: function(newPage) {  // 页码值发送变化
      this.queryInfo.pageNum = newPage;
      this.getBlogList();
    },
  }
}
</script>

<style scoped>

</style>
