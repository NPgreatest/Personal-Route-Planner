<template>
  <el-card :style="{ margin: '0px 50px' }" class="card-h">
    <el-row>
      <el-col :span="10">
        <div class="container">
          <h2>大模型应用底座平台</h2>
          <h2 style="color: #0c89fd">应用构建</h2><br />
        </div>
        <h5 style="text-align: left; color: #808080">
          欢迎来到大模型对话应用！在这里您可以自由创建、管理您的大模型问答应用。
          您可以根据需要选择已创建的知识库、编写专属提示词、选用最合适的大模型来创建问答应用，期待大模型能够理解并回答您的各类问题！
        </h5>
      </el-col>
      <el-col :span="14">
        <img alt="" src="../assets/images.jpg" style="float: right;">
      </el-col>
    </el-row>
  </el-card>
  <el-card :style="{ margin: '19px 50px' }" class="card-kb">
    <div style="display: flex;justify-content: space-between;align-items: center;">
      <span>
        <h3 style="margin:0;">应用列表</h3>
      </span>
      <span style="display: flex ;justify-content: center;align-items: center;" class="search-container2">
        <el-input style="width: 300px;" v-model="searcAppQuery" placeholder="请搜索应用" class="search-input"></el-input>
        <i class="search-icon2">🔍</i>
        <ul id="searchResults"></ul>
      </span>
    </div>
    <!-- <div style="text-align: left;margin-top: 10px;">
      <el-button type="primary" round @click="currentAllapp = allapp">全部主题</el-button>
      <el-button type="info" text="info" round v-for="theme in themeList" :key="theme.id"
        @click="updateallapp(theme.name)">{{
          theme.name }}</el-button>

      添加下拉框
      <el-select placeholder="更多" style="margin-left: 20px;width: 70px;">
        <el-option @click="this.dialogAddThemeVisible = true">新增</el-option>
        <el-option @click="dialogEditThemeVisible = true">编辑</el-option>
      </el-select>
    </div>  -->
    <n-divider title-placement="left"></n-divider>

    <el-row :gutter="12" style="margin-top: 0px;">
      <el-col :span="8" style="height: 20px;">
        <div class="module" style="border: 1px dashed #aad8e3;cursor: pointer;" @click="$router.push('/create')">
          <br />
          <el-icon style="color: #0c89fd;font-size: 40px;">
            <Management />
          </el-icon>
          <br />
          <el-button text>新建应用</el-button>
          <p>快来创建你的应用吧</p>
        </div>
      </el-col>
      <el-col :span="8" v-for="app in paginatedApps" :key="app._id">
        <div class="module">
          <div class="app-head-row">
            <div class="app-name">
              <div class="app-avatar" :style="generateRandomColorStyle()">{{ app.name[0] }}</div>
              <el-button text @click="$router.push('/detail/' + app._id)">{{ app.name }}</el-button>
              <el-tag>{{ app.state }}</el-tag>
            </div>
            <div class="card-dropdown">
              <el-dropdown trigger="click">
                <span class="el-dropdown-link">
                  ···
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item @click="authorizeapp">授权</el-dropdown-item>
                    <el-dropdown-item @click="editapp(app)">编辑</el-dropdown-item>
                    <el-dropdown-item @click="deleteapp(app)">删除</el-dropdown-item>
                  </el-dropdown-menu>
                  <el-dialog v-model="dialogVisible" title="授权" width="55%">
                    <div>
                      <h3 style="margin-left: 20px;margin-top: 0px;height: 10px;">应用名称</h3>
                      <p style="margin-left: 20px;">{{ app.name }}</p>
                    </div>
                    <h3 style="margin-left: 20px;height: 10px;margin-left: 20px;">添加用户</h3>
                    <div class="search-container">
                      <el-input v-model="searchQuery" placeholder="请搜索用户姓名/账号" class="search-input"></el-input>
                      <i class="search-icon">🔍</i>
                    </div>
                    <div v-if="searchQuery" class="infinite-list" style="max-height: 140px; overflow-y: auto;">
                      <ul class="infinite-list" style="overflow: auto">
                        <li class="infinite-list-item" v-for="account in matchedAccounts" :key="account"
                          @click="selectAccount(account)">
                          <div>
                            {{ account }}
                          </div>
                        </li>
                      </ul>
                    </div>
                    <div>
                      <h3 style="margin-left: 20px;height: 10px;margin-left: 20px;">权限</h3>
                      <p style="margin-left: 20px;">被选择的用户具有该发布应用的对话权限</p>
                    </div>
                    <div style="text-align: right;">
                      <el-button @click="dialogVisible = false">取消</el-button>
                      <el-button @click="inviteUser">搜索并邀请</el-button>
                    </div>
                  </el-dialog>
                </template>
              </el-dropdown>
            </div>
          </div>
          <p class="app-detail">
            {{ app.details }}
          </p>
          <div class="app-description">
            <div class="descriptioncounts">调用次数{{ app.num }}</div>
          </div>
          <router-view />
        </div>
      </el-col>
    </el-row>
    <el-pagination v-if="totalPages > 1" @current-change="handlePageChange" :current-page.sync="currentPage"
      :page-size="pageSize" :total="filterAllapp.length"></el-pagination>
  </el-card>

  <el-dialog v-model="dialogAddThemeVisible" title="新建主题" width="30%">
    <el-form :model="form_create" label-width="120px" style="margin-top: 1px;">
      <el-form-item style="margin-top: 20px;" label="主题名称" required>
        <el-input v-model="form_create.theme" style="width: 600px;" />
      </el-form-item>
      <div style="text-align: left;">
        <el-button type="primary" plain @click="dialogAddThemeVisible = false" style="margin-left: 160px;">取消</el-button>
        <el-button type="primary" @click="createTheme()">保存</el-button>
      </div>
    </el-form>
  </el-dialog>

  <el-dialog v-model="dialogEditThemeVisible" title="编辑主题" width="30%">
    <el-form :model="form_edit" label-width="120px" style="margin-top: 1px;">
      <el-form-item label="主题名称:" v-for="theme in themeList" :key="theme.id" required>
        <el-input v-model="theme.name" style="width: 100px;margin-right: 10px;" />
        <el-button @click="deleteTheme(theme.id)">删除</el-button>
        <el-button type="primary" @click="updateTheme(theme.id, theme.name)">更新</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>
</template>

<style>
.el-dialog__header {
  justify-content: flex-start;
}

.el-dialog__header {
  padding-bottom: 0% !important;
}

.el-dropdown-link {
  position: relative;
  width: 100%;
}
</style>


<style scoped>
.infinite-list {
  height: 300px;
  padding: 0;
  margin: 0;
  list-style: none;
}

.infinite-list .infinite-list-item {
  display: flex;
  align-items: center;
  justify-content: center;
  height: 35px;
  background: var(--el-color-primary-light-9);
  margin: 10px;
  color: var(--el-color-primary);
}

.infinite-list .infinite-list-item+.list-item {
  margin-top: 5px;
}

.search-input {
  width: 100%;
  /* 占据整个容器宽度 */
  height: 30px;
  /* 设置搜索框的高度 */
  padding-right: 30px;
  /* 右内边距，留出空间放置图标 */
  border: 1px solid #ccc;
  /* 添加边框 */
  border-radius: 5px;
  /* 添加圆角 */
  outline: none;
  /* 移除默认的输入框轮廓 */
}

.search-icon {
  position: absolute;
  /* 绝对定位图标 */
  top: 50%;
  /* 垂直居中 */
  right: 10px;
  /* 设置图标距离输入框右边的距离 */
  transform: translateY(-50%);
  /* 垂直居中 */
  font-size: 1.2em;
  /* 设置图标大小 */
  cursor: pointer;
  /* 设置鼠标悬停时的指针样式 */
  pointer-events: none;
  /* 防止图标遮挡搜索框的输入 */
}

.search-icon2 {
  position: absolute;
  /* 取消注释来绝对定位图标 */
  top: 6.5%;
  right: 5%;
  transform: translateY(-50%);
  /* 这行代码现在可以正常工作 */
  font-size: 1.2em;
  cursor: pointer;
  pointer-events: none;
  /* 可考虑添加 left 或 right 属性来调整图标的水平位置 */
}

.search-container2 {
  display: flex;
  align-items: center;
  /* 添加这一行来保证垂直居中 */
  /* ...其他样式... */
}

.search-container {
  margin-top: 15px;
  margin-left: 20px;
  /* 与上方的应用名称间距 */
  display: flex;
  /* 使用 flex 布局 */
  align-items: center;
  position: relative;
  /* 定位搜索框的容器 */
  /* 垂直居中 */
}

.search-container el-input {
  flex: 1;
  /* 占据剩余空间 */
  margin-right: 10px;
  /* 与搜索按钮的间距 */
}

.module {
  /* 可选：底部间距，以分隔行 */
  margin: 10px;
  border: 1px solid #f0f0eb;
  /* 添加样式以定义模块的外观 */
  border-radius: 10px;
  /* 圆角 */
  padding: 20px;
  /* 内边距 */
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  /* 阴影效果 */
  background-color: #f6fafae7;
  /* 背景色 */
  width: 80%;
  height: 130px;
}

.app-head-row {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  font-size: large;
  font-weight: bold;
}

.app-name {
  display: flex;
  flex-direction: row;
  justify-content: left;
  align-items: center;
  font-size: large;
  font-weight: bold;
}

.app-head-row .app-avatar {
  width: 48px;
  height: 48px;
  background: RGB(214, 234, 253);
  color: RGB(96, 178, 254);
  margin-right: 10px;
  border-radius: 24px;
  line-height: 46px;
}

.card-dropdown {
  cursor: pointer;
}

.card-dropdown:hover {
  background: #dddddd;
}

.descriptioncounts {
  text-align: left;
  padding: 0 10px
}

.app-detail {
  text-align: left;
  color: gray;
  font-size: 16px;
  margin-top: 25px;
  padding: 0 10px;
}

.app-use {
  display: inline-block;
  width: 33.3%;
}

.card-kb {
  height: 553px;
}

.app-buttons {
  display: flex;
  flex-direction: row;
  justify-content: space-around;
  align-items: center;
}

.container {
  display: flex;
  margin-bottom: -20px;
}

.card-kb {
  height: 553px;
}

.card-h {
  height: 130px;
}
</style>

<script>
// @ is an alias to /src
import { mapGetters } from 'vuex'
import { Management, Rank } from "@element-plus/icons-vue";
import { createTheme } from 'naive-ui';

export default {
  data() {
    return {
      currentPage: 1,
      pageSize: 5, // 每页显示的数量
      dialogVisible: false,
      dialogAddThemeVisible: false,
      dialogEditThemeVisible: false,
      searcAppQuery: '',
      searchQuery: '',
      searchResults: '',
      platformAccounts: ['account1', 'account2', 'account3', 'exampleAccount', 'testAccount'],
      colorList: [
        {
          background: '#87CEFA',
          color: '#0000FF'
        },
        {
          background: '#7FFFAA',
          color: '#228B22'
        },
        {
          background: '#FDF5E6',
          color: '#FFA500'
        },
        {
          background: '#FFA07A',
          color: '#FF4500'
        },
        {
          background: '#FFB6C1',
          color: '#FF1493'
        },
      ],
      themeList: [
        { id: 1, name: '宝马' },
        { id: 2, name: '奔驰' },
        { id: 3, name: '奇骏' },
      ],
      form_create: {},
      form_edit: {},
      currentAllapp: [],
      filterAllapp: [],
    }
  },
  computed: {
    ...mapGetters(['allapp', 'app']),
    totalPages() {
      return Math.ceil(this.filterAllapp.length / this.pageSize); // 计算总页数
    },
    matchedAccounts() {
      // 使用模糊匹配筛选匹配的账号  
      return this.platformAccounts.filter(account => account.includes(this.searchQuery));
    },
    paginatedApps() {
      this.filterAllapp = this.filter_allapp()
      const startIndex = (this.currentPage - 1) * this.pageSize;
      const endIndex = startIndex + this.pageSize;
      return this.filterAllapp.slice(startIndex, endIndex);
    },
  },
  name: 'AppView',
  methods: {
    editapp(app) {
      if (app.state === '已发布') {
        alert("当前应用已发布，请在应用停用后再进行编辑/删除");
      }
      else {
        this.$router.push('/edit/' + app._id)
      }
    },
    updateallapp(name) {
      this.currentAllapp = this.allapp.filter(app => app.theme.includes(name))
    },
    filter_allapp() {
      return this.currentAllapp.filter(app => app.name.includes(this.searcAppQuery));
    },
    createTheme() {
      this.dialogAddThemeVisible = false
      console.log(this.form_create)
    },
    editTheme(theme) {
      if (theme) {
        this.selectedTheme = theme; // 保存当前选中的主题  
        this.editForm.name = theme.name; // 将主题名称填充到编辑框中  
        this.dialogEditThemeVisible = true; // 打开编辑对话框  
      } else {
        // 处理点击"编辑"选项的情况，可以对所有主题进行编辑操作  
        // 这里可以根据你的需求来编写逻辑，比如遍历主题列表进行批量编辑操作  
      }
    },
    // 更新主题的方法，用于在编辑对话框中点击确定按钮时更新主题名称  
    updateTheme() {
      if (this.selectedTheme) {
        this.selectedTheme.name = this.editForm.name; // 更新选中主题的名称  
        this.dialogEditThemeVisible = false; // 关闭编辑对话框  
      }
    },
    handlePageChange(newPage) {
      console.log(this.currentPage)
      this.currentPage = newPage;
    },
    authorizeapp() {
      this.dialogVisible = true;
      console.log(this.dialogVisible)
    },
    selectAccount(account) {
      this.searchQuery = account; // 将点击的账号填充到输入框中 
      this.searchResult = account
    },
    inviteUser() {
      this.dialogVisible = false
      console.log(this.searchResult)
    },
    generateRandomColorStyle() {
      let random_index = Math.floor((Math.random() * this.colorList.length));
      console.log(this.colorList[random_index].background, this.colorList[random_index].color)
      return {
        'background-color': this.colorList[random_index].background,
        'color': this.colorList[random_index].color
      }
    },
    deleteapp(app) {
      if (app.state === '已发布') {
        alert("当前应用已发布，请在应用停用后再进行编辑/删除");
      }
      else{
        this.$confirm('确认删除该应用吗？', '提示', {
          confirmButtonText: '确定',
          cancelButtonText: '取消',
          type: 'warning'
        }).then(async () => {
          // 用户点击了确定按钮    
          try {
            await this.$store.dispatch('deleteApp/deleteApp', { _id: app._id });
            // 在deleteApp执行成功后刷新页面    
            this.$router.go(0);
          } catch (error) {
            console.log('[App]: 删除失败！');
          }
        }).catch(() => {
          // 用户点击了取消按钮，返回about页面  
          this.$router.push('/about');
        });
      }
    }

  },
  components: {
    Management,
    Rank
  },
  async mounted() {
    await this.$store.dispatch('getallApp/getallApp').then((res) => {
    }).catch(() => {
      console.log('[HomeView]: 获取hello失败！')
    })
    this.currentAllapp = this.allapp
  }
}
// .card-dropdown: hover {
//   background: #dddddd;
// }
</script>
