
<template>

  <div class="tag-bg">
    <div class="title" align="center">景点分类</div>


    <!-- 标签栏 -->
    <transition appear
                name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__fadeInDown"
                leave-active-class="animate__fadeOutUp"
    >
      <el-row  class="tag-area" :gutter="20" style="margin: auto; display: flex;flex-wrap: wrap; padding-top: 20px; padding-bottom: 20px "    >
          <li :key="item.tagid" v-for="item in tags">
            <el-col :span="4">
            <el-button type="success" @click="jumpPage(item.tagid)"  class="buttons" >{{item.name}}</el-button>
            </el-col>
          </li>
      </el-row>
    </transition>

    <transition appear
                name="animate__animated animate__bounce animate__slow"
                enter-active-class="animate__fadeInUp"
    >
      <div>
        <!-- 博客列表栏 -->
        <div class="blog-area">
          <SiteCard class="blog-item" :key="item.id" v-for="(item, index) in siteDetails"
                    :item="item" :imgRight="index % 2 === 0"
                    @click.native="siteDetail(item.sid)">
          </SiteCard>
        </div>
      </div>
    </transition>

  </div>
</template>



<script>

//import img1 from "../assets/1.png"

//import BlogCard from "../components/BlogCard";
//import Pagination from "../components/Pagination";

import SiteCard from "../components/SiteCard";
export default {
  components:{SiteCard},
  // eslint-disable-next-line vue/multi-word-component-names
  name:"tags",
  //components: { BlogCard, Pagination, BlogTagItem },
  data() {
    return {
      currentTagId: 0,
      pages: 1,              // 页面数量
      queryInfo: {
        pageNum: 1,
        pageSize: 5,
        tagId: 0
      },
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
      siteDetails: [{
        id: 1,
        title: "博客系统开发",
        description: "本章主要介绍博客系统的开发工作...",
        firstPicture: "",
        nickname: "Jack Ma",
        avatar: "https://placekitten.com/300/300",
        typename: "框架底层原理",
        createTime: "2021-03-29",
        views: 1024
      },
        {
          id: 2,
          title: "博客系统开发",
          description: "本章主要介绍博客系统的开发工作...",
          firstPicture: "",
          nickname: "Jack Ma",
          avatar: "https://placekitten.com/300/300",
          typename: "框架底层原理",
          createTime: "2021-03-29",
          views: 1024
        },
        {
          id: 3,
          title: "博客系统开发",
          description: "本章主要介绍博客系统的开发工作...",
          firstPicture: "",
          nickname: "Jack Ma",
          avatar: "https://placekitten.com/300/300",
          typename: "框架底层原理",
          createTime: "2021-03-29",
          views: 1024
        }],
    }
  },
  created() {
    window.scrollTo(0, 0)
    this.getTagList()
  },
  methods: {
    async getTagList() {
      const {data: res} = await this.$axios.get("/home/alltags");
      if(res.status === 1) {
        this.tags = res.data.length > 0 ? res.data[0] : this.tags;
      }
      await this.getBlogList(1);
    },
    async getBlogList(id) {
      const {data: res} = await this.$axios.get("/home/sitesbytags", {params: {tagid: id}});
      if(res.status === 1) {

        this.siteDetails = res.data.length > 0 ? res.data[0] : this.siteDetails;
      } else {
        this.$message.error("获取景点失败，请重试")
        return
      }
    },
    jumpPage(id) {
      window.scrollTo(0, 0)
      this.getBlogList(id);
    },
    siteDetail(id) {
      this.$router.push({
        path: "/sitedetails?id="+id
      });
    }
  }
}

</script>



<style lang="less" scoped>

ul, li {
  margin: 0;
  padding: 0;
}

.tag-bg {
  background: url('../static/background.jpg') 0 0 / cover no-repeat;
  background-attachment: fixed;
  min-height: 1000px;
}

.title {
  font-size: 450%;
  color: #ffffff;
  margin-bottom: 50px;
  bottom: 0 !important;
  right: 0 !important;
  font-family:'STXingkai';
  opacity: 0.5;
  padding-top: 6%;
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

.blog-area {
  width: 850px;
  margin: 40px auto 50px;
  padding-bottom: 100px;
}

.clearfix::after {
  content: '';
  display: block;
  height: 0;
  clear: both;
  visibility: hidden;
}

.pagebar {
  padding-bottom: 50px;
}

.buttons:hover {
  transform:  scale(1.22);
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
}

</style>
