<template>
  <div style="text-align: left;margin-left: 50px;margin-bottom: 10px;"><el-button plain type="primary"
      @click="$router.push('/about')"> 返回应用列表</el-button></div>
  <el-row>
    <el-col span="8" v-if="showCard1">
      <el-card
        style="height: 600px; width: 300px; margin-left: 50px; border-radius: 8px; display: flex; flex-direction: column; justify-content: center;background-color:">
        <el-card style="color:white;background-color: #e8f1fd;height: 280px;">
          <h3 style="line-height: 0px;color: #0c89fd;text-align: left;">帮助小助手</h3>
          <el-card style="background-color: rgb(251, 252, 252);">
            <h3 style="margin-bottom: 1px;">知识库选择</h3>
            <p>如有调整，您可以为此应用关联一个知识库，以便大模型能够提供更加专业、精准的回答</p>
          </el-card>
        </el-card>
      </el-card>
    </el-col>
    <el-col span="8" v-if="showCard3">
      <el-card
        style="height: 600px; width: 300px; margin-left: 50px; border-radius: 8px; display: flex; flex-direction: column; justify-content: center;background-color:">
        <el-card style="color:white;background-color: #e8f1fd;height: 280px;">
          <h3 style="line-height: 0px;color: #0c89fd;text-align: left;">帮助小助手</h3>
          <el-card style="background-color: rgb(251, 252, 252);height: 200px;">
            <h3 style="margin-bottom: 1px;">下一步做什么</h3>
            <p>您的应用已经创建完成，快去调试你的应用吧！</p>
          </el-card>
        </el-card>
      </el-card>
    </el-col>
    <el-col span="8" v-if="showCard2">
      <el-card style="height: 600px;width: 1000px;margin-left: 30px;border-radius: 8px;">
        <el-form :model="form" label-width="120px">
          <el-form-item style="margin-top: 30px;" label="应用名称" required>
            <el-input v-model="form.name" style="width: 600px;" />
          </el-form-item>

          <el-form-item style="margin-top: 30px;" label="应用描述">
            <el-input v-model="form.details" style="height: 100px;width: 600px;" />
          </el-form-item>
          <el-form-item style="margin-top: 30px;" label="知识库">
            <el-select v-model="form.knowledgebase" placeholder="选择知识库" multiple style="width: 300px;">
              <el-option v-for="app in knowledge_list" :label="app.name" :key="app._id" :value="app._id" />
            </el-select>
          </el-form-item>
          <el-form-item style="margin-top: 30px;" label="绑定到景点(id)">
            <el-input v-model="form.bind" type="number" style="height: 50px;width: 100px;" />
          </el-form-item>
          <!-- <el-form-item label="应用主题" required>
            <el-select v-model="form.appTheme" placeholder="选择主题" style="width: 300px;" multiple>
              <el-option v-for="app in [
                { '_id': '654dd9c4ea13fc974c0504c9', 'name': 'chatgml_turbo' },
                { '_id': '654dd9feea13fc974c0504cd', 'name': 'skylark_plus' },
                { '_id': '654dda09ea13fc974c0504d0', 'name': 'skylark_pro' }
              ]" :label="app.name" :key="app._id" :value="app._id" />
            </el-select>
          </el-form-item> -->
          <div style="text-align: left;">
            <el-button type="primary" plain @click="$router.push('/about')" style="margin-left: 120px;">取消</el-button>
            <el-button type="primary" @click="submitForm()">保存</el-button>
          </div>
        </el-form>
      </el-card>
    </el-col>
    <el-col span="8" v-if="showCard4">
      <el-card style="height: 600px;width: 1000px;margin-left: 30px;border-radius: 8px;">
        <img alt="" src="../assets/images.jpg" style="float: center;">
        <h2>应用创建完成！</h2>
        <el-card style="width: 800px;height: 175px;background-color: #e8f1fd;margin-left: 84px;">
          <div style="text-align: left;">应用名称：{{ form.name }}</div>
          <div>
            <div style="text-align: left;margin-top: 15px;">知识库:</div>
            <el-card style="height: 52px;margin-top: 10px;text-align: left;">
              <p style="margin-top: 0px;display: inline-block;margin-right: 10px;" v-for="app in selectKnowledgebase">{{
                app }}</p>
            </el-card>
          </div>
        </el-card>
        <div style="margin-top: 20px;">
          <el-button type="primary" plain @click="$router.push('/about')" style="margin-left: 10px;">应用列表-></el-button>
          <el-button type="primary" @click="$router.push('/detail/' + create.data._id)">去调试-></el-button>
        </div>
      </el-card>
    </el-col>
  </el-row>
</template>
<style></style>
<script>
import { mapGetters } from 'vuex'

export default {
  data() {
    return {
      form: {
        state: '未发布',
        prompt: "",
        num: 1,
        bind: 0,
      },
      showCard1: true,
      showCard2: true,
      showCard3: false,
      showCard4: false,
      app_id: ''
    };
  },
  computed: {
    ...mapGetters(['knowledge_list', 'allapp', 'create']),
    selectKnowledgebase() {
      if (!this.form.knowledgebase || this.form.knowledgebase.length === 0) {
        return ['无知识库'];
      };
      return this.form.knowledgebase.map(id => {
        const item = this.knowledge_list.find(item => item._id === id);
        return item ? item.name : '未找到';
      })
    },
  },
  methods: {
    submitForm() {
      console.log(this.form)
      console.log(this.form)
      try {
        this.$store.dispatch('createApp/createApp', this.form).then((res) => {
          this.showCard1 = false
          this.showCard2 = false
          this.showCard3 = true
          this.showCard4 = true;
          this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
          }).catch(() => {
            console.log('[KnowledgeBase]: 获取知识库失败！')
          })
        }).catch(() => {
          console.log('[HomeView]: 获取hello失败！')
        });
        console.log(this.create.data)
      } catch (error) {
        // 处理请求失败的情况
        console.error('提交失败:', error);
      }
    }
  },
  mounted() {
    this.$store.dispatch('get_knowledgebase/getKnowledgeList').then((res) => {

    }).catch(() => {
      console.log('[HomeView]: 获取hello失败！')
    })

  }
};
</script>
