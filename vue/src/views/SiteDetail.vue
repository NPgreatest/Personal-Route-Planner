
<template>

  <div class="site_se">

    <!-- 标题 -->
    <div class="title">
      <h1 style="padding-top: 60px;" > {{ ' '+site.sname+' ' }}</h1>
    </div>
    <div style="display: flex">
      <div style="flex: 1"></div>
      <el-button type="primary" style="width: 10%" @click="last">上一景点</el-button>
      <el-button  type="primary" style="width: 10%" @click="front">下一景点</el-button>
      <div style="flex: 1"></div>
    </div>



      <p top="100px"></p>
      <el-row type="flex" class="row-bg" justify="center"  gutter="50">

        <el-col :span="6"><div class="grid-1" style="left: 10%" />
          <div id="chart" style="width: 600px; height: 400px; left:-30%"></div>
          <div class="grid-content ep-bg-purple" />
        </el-col>

        <el-col :span="6"><div class="grid-2" style="top: 10%; " />
          <div class="background">
            <img :src="imgSrc" style="width:70%; height:70%; top:100%; left: 30%"  />
          </div>
          <el-link  :href="site.website"  type="success" class="link">官方网址</el-link>
          <div class="grid-content ep-bg-purple" />
        </el-col>

      </el-row>


    <!-- 正文 -->
    <div class="content-outer">
      <div class="content-inner" v-html="site.description">

      </div>
    </div>

    <!-- 评论 -->
    <div class="comment-outer">
      <div class="comment-body">
        <div style="font-weight: bold">评论</div>
        <hr>

        <!-- 评论内容 -->
        <div class="comment" :key="item.id" v-for="item in commentList">
          <b-avatar class="comment-avatar" :src="item.avatar" ></b-avatar>
          <!-- 昵称和日期 -->
          <!-- <div class="comment-author"> -->
          <span class="author">
                            {{ item.name }}
                        </span>
          <span class="date">
                            {{ item.time | fromatDate("yyyy-MM-dd hh:mm:ss") }}
          </span>
          <!-- </div> -->
          <!-- 回复内容 -->
          <div class="comment-content">
            {{ item.content }}
          </div>
        </div>

      </div>

      <!-- 回复栏 -->
      <textarea class="replay" v-model="comment.content">
            </textarea>

      <b-container fluid>
        <b-row >
          <b-col sm="3">
            <b-button variant="outline-info" @click="publishComment" >
              发布
            </b-button>
          </b-col>
        </b-row>
      </b-container>
    </div>

    <div style="height: 200px;"></div>
  </div>

</template>



<script>

export default {
  name:"SiteDeatil",
  data() {
    return {
      //imgSrc: require('../static/logo.png'),
      site: {
        sid:"1",
        sname:" ",
        webset: "test",
        discription: `esfsfesfesfsefsefs`,
        updateTime: "2021-03-30",
        createTime: "2021-03-30",
        views: 1024,
        appreciation: false,
      },
      comment: {         //评论数据
        nickname: "",
        email: "",
        content: "",
        //avatar: "http://localhost:8080/images/1.jpg",
      },
      commentList: [
        {
          id:1,
          name: "小兔叽",
          content: "very good!",
          //avatar: "http://localhost:8080/images/1.jpg",
          createTime: "2021-03-31 22:35",
        },
        {
          id:2,
          name: "小脑斧",
          content: "excellent!",
          //avatar: "http://localhost:8080/images/1.jpg",
          createTime: "2021-03-31 22:35",
        }
      ]
    }
  },
  created() {
    window.scrollTo(0, 0);
    this.getSiteDetails();
  },
  filters: {
    fromatDate: function(val, fmt) {
      var date = new Date(val);
      let ret;
      const opt = {
        "y+": date.getFullYear().toString(),        // 年
        "M+": (date.getMonth() + 1).toString(),     // 月
        "d+": date.getDate().toString(),            // 日
        "h+": date.getHours().toString(),           // 时
        "m+": date.getMinutes().toString(),         // 分
        "s+": date.getSeconds().toString()          // 秒
        // 有其他格式化字符需求可以继续添加，必须转化成字符串
      };
      for (let k in opt) {
        ret = new RegExp("(" + k + ")").exec(fmt);
        if (ret) {
          fmt = fmt.replace(ret[1], (ret[1].length == 1) ? (opt[k]) : (opt[k].padStart(ret[1].length, "0")))
        }
      }
      return fmt;
    }
  },
  mounted(){
    this.drawLine();
  },
  methods: {
    getSiteDetails: async function() {
      let siteId = this.$route.query.id;
      siteId = parseInt(siteId);
      const {data: res} = await this.$axios.get("/home/sitedetails", {params: {siteid: siteId}});
      if(res.status === 1) {
        this.site = res.data.length > 0 ? res.data[0] : this.site;
        this.imgSrc=this.site.pic;

        var hljs = require('highlight.js');
        var md = require('markdown-it')({
          html: true,
          linkify: true,
          typographer: true,
          highlight: function (str, lang) {
            if (lang && hljs.getLanguage(lang)) {
              try {
                return '<pre class="hljs"><code>' +
                    hljs.highlight(lang, str, true).value +
                    '</code></pre>';
              } catch (__) { }
            }

            return '<pre class="hljs"><code>' + md.utils.escapeHtml(str) + '</code></pre>';
          }

        });

        // 让图片居中
        md.renderer.rules.image = (token, idx, options, env, self) => {
          return "<div align='center' class='blog_image'>" + md.renderer.renderToken(token, idx, options, env, self) + "</div>"
        }
        this.site.description = md.render(this.site.description);

        this.getCommentList();
      }
      this.drawChart();
    },
    last(){
      const next = this.site.sid - 1;
      if (next===0){
        this.$message.info('已经到了首个景点');
          return;
      }
      this.$router.push("/sitedetails?id="+next.toString())
      location.reload()
    },
    front: async function(){
      const next = this.site.sid + 1;
      const {data: res} = await this.$axios.get("/home/sitedetails", {params: {siteid: next}});
      if (res.status !=1){
        this.$message.info('到达最后一个景点');
        return;
      }
      this.$router.push("/sitedetails?id="+next.toString())
      location.reload()
    },
    publishComment: async function () {
      if(!this.comment.content) {
        this.$message.warning("请输入评论内容")
        return
      }
      this.comment.name=window.sessionStorage.getItem("userId");
      this.comment.Sid=this.site.sid;
      //alert(this.comment.name+this.comment.Sid);
      await this.$axios.post("/user/comment", this.comment);//sgrgdrgergesfsegsgdrgdrgdrgrdgdgdrgdrgdvrd
      if (window.sessionStorage.getItem("userId")==null){
        this.$message.error('评论失败，请登录。')
        this.$router.push("/login")
        return;
      }
      this.$message.success('评论成功！')
      location.reload()
    },
    getCommentList: async function() {

      let siteId = this.$route.query.id;
      const {data: res} = await this.$axios.get("/home/getcomments", {params: {siteid: siteId}});
      if(res.status === 1) {
        this.commentList = res.data.length > 0 ? res.data[0] : this.commentList;
      }
    },
    drawChart() {
      let myChart = this.$echarts.init(document.getElementById("chart"));

      var option = {
        title: {
          text: "景点价格",
        },
        tooltip: {},
        legend: {
          data: ["价格"],
        },
        xAxis: {
          data: ["成人票","儿童票"],
        },
        yAxis: {},
        series: [
          {
            name: "销量",
            type: "bar",
            data: [0,0],
          },
        ],
      };
      myChart.setOption(option);
      for (var i=0;i<this.site.prices.length;i++){
        //alert(this.site.prices[i].aimed);
        option.xAxis.data[i]=this.site.prices[i].aimed;
        option.series[0].data[i]=this.site.prices[i].price;
      }
      myChart.setOption(option);

    },
  }
}
</script>

<style>

/*改变博客中的图片的最大宽度*/
.site_image img {
  max-width: 1000px !important;
  cursor: pointer;
}

</style>

<style lang="less" scoped>

.common-layout{

  border-radius: 15px;
  background-clip: padding-box;
  margin: 20px auto;
  width: 1500px;
  padding: 15px 35px 15px 35px;
  background: aliceblue;
  border:1px solid blueviolet;
  box-shadow: 0 0 25px #643965;
  z-index:1;
}

.pageContainer{
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

.site_se {
  background: url('~@/assets/bg1.jpg')  fixed 0px 0px;
  //min-height: 1000px;
  //padding-top: 8%;
  width:100%;
  height:100%;
}

.title {
  color: #FFFFFF;
  font-size: 40px !important;
  font-family: STSong;
  text-align: center;
  margin-top: 20px;
}

.siteInfo {
  color: #FFFFFF;
  text-align: center;
  margin-top: 30px;
  font-size: 18px;
  padding-top: 5px;
  padding-bottom: 5px;

  span {
    margin-left: 12px;
  }

  .flag {
    background-color: #FFF;
    color: #F2711C;
    border-color: #F2711C;
    border: 1px solid;
    box-shadow: none;
    padding: 2px;
    border-radius: .28571429rem;
    font-size: 14px;
  }

}


.content-outer {
  width: 76%;
  min-height: 800px;
  margin-top: 60px;
  margin-left: 12%;
  margin-right: 12%;

  top: 0;
  bottom: 0;
  border-radius: .28571429rem;
  box-shadow: none;
  border: 1px solid #D4D4D5;
  padding: 1.5em;


  background: #FFF;
  opacity: 0.9;

}

.content-inner {
  padding: 4em 40px 2em;
  box-sizing: border-box;
  font: 1em/1.5 Tahoma,Helvetica,Arial,'宋体',sans-serif !important;
  color: rgba(0,0,0,.87);
  height: auto;

}
.link{
  left: 50px;
  font: 2em/1.5 Tahoma,Helvetica,Arial,'Times New Roman',sans-serif !important;
}


.rights {
  width: 76%;
  margin-left: 12%;
  margin-right: 12%;
  border-radius: .28571429rem;
  box-shadow: 0 0 0 1px #a3c293 inset;
  border: 1px solid #a3c293;
  padding: 1.5em;


  background: #FCFFF5;
  color: #2C662D;
  opacity: 0.9;

  .list {
    text-align: left;
    padding: 0;
    opacity: .85;
    list-style-position: inside;
    margin-bottom: 0.8em;

  }
}

.comment-outer {
  width: 76%;
  margin-left: 12%;
  margin-right: 12%;
  border-radius: .28571429rem;
  box-shadow: none;
  border: 1px solid #D4D4D5;
  padding: 1.5em;

  min-height: 300px;
  background: #FFF;
  opacity: 0.9;


  //margin-bottom: 60px;
}

.comment-body {
  position: relative;
  background: #FFF;
  box-shadow: 0 1px 2px 0 rgb(34 36 38 / 15%);
  margin: 0.5rem 0;
  padding: 1em;
  border-radius: .28571429rem;
  border: 1px solid rgba(34,36,38,.15);
  border-top: 2px solid #00B5AD;
  min-height: 260px;

}

.comment {
  width: 90%;
  padding: 6px 18px;
  margin-left: 10px;
  margin-right: 20px;
  margin-top: 12px;
  display: block;

  .comment-avatar {
    float: left;
    margin-top: 6px;
    margin-right: 12px;
  }

  .author {
    font-size: 1.1em;
    color: rgba(0,0,0,.87);
    font-weight: 700;
    margin-left: 8px;
  }

  .date {
    display: inline-block;
    color: rgba(0,0,0,.4);
    font-size: .875em;
    margin-left: 8px;
  }


  .comment-content {
    margin-top: 4px;
    font-size: 15px;
    word-wrap: break-word;
    color: rgba(0,0,0,.87);
    line-height: 1.3;
    padding-left: 60px;
  }
}

.replay {
  width: 100%;
  min-height: 180px;
  background: #FFF;
  margin: 0.5rem 0;
  padding: 1em;
  border-radius: .28571429rem;

  padding: .78571429em 1em;
  border: 1px solid rgba(34,36,38,.15);
  outline: 0;
  color: rgba(0,0,0,.87);
  box-shadow: 0 0 0 0 transparent inset;
  transition: color .1s ease,border-color .1s ease;
  font-size: 1em;
  line-height: 1.2857;
  resize: vertical;
}

//// 多行文本溢出显示省略号
//.text-ellipsis {
//    //弹性伸缩盒子模型
//    display: -webkit-box;
//    //限制在一个块元素显示的文本行数
//    -webkit-line-clamp: 3;
//    //设置或检索伸缩盒子对象的子元素的排列方式
//    -webkit-box-orient: vertical;
//}

.el-row {
  margin-bottom: 20px;
}
.el-row:last-child {
  margin-bottom: 0;
}
.el-col {
  border-radius: 4px;
}

.grid-content {
  border-radius: 4px;
  min-height: 36px;
}

</style>




