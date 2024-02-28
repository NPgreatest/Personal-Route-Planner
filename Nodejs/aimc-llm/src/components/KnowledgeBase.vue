<template>
  <el-card :style="{ margin: '0px 50px' }" class="card-h">
    <el-row>
      <el-col :span="14">
        <div style="color: #188ffd; text-align: left; font-size:24px;font-weight: bold;">
          知识库构建
        </div>
        <div style="text-align: left; color: #808080;font-size:12px;font-weight: bold;margin-top: 4px;">
          在这里您可以创建和管理自己的知识库，通过关联知识库，大模型能够提供更加专业、精准的回答。<br />
          快来创建你的专属知识库吧，让大模型成为特定场景的业务专家！
        </div>
      </el-col>
      <el-col :span="2">
        <div class="card-icon" style="color: #0c89fd;font-size: 70px">
          <el-icon>
            <FolderChecked />
          </el-icon>
        </div>
      </el-col>
      <el-col :span="2">
        <div style="text-align: left;font-size:18px;font-weight: bold;margin-top: 20px;">知识库信息</div>
        <div style="text-align: left; color: #808080; font-size:12px;font-weight: bold;margin-top: 4px;">
          请输入知识库基本信息
        </div>
      </el-col>
      <!-- <el-icon><Right /></el-icon> -->
      <el-col :span="2">
        <div class="card-icon" style="color: #0c89fd;font-size: 70px">
          <el-icon>
            <School />
          </el-icon>
        </div>
      </el-col>
      <el-col :span="4">
        <div style="text-align: left;font-size:18px;font-weight: bold;margin-top: 20px;">文档配置</div>
        <div style="text-align: left; color: #808080;font-size:12px;font-weight: bold;margin-top: 4px;">
          上传知识文档，设置分段方案
        </div>
      </el-col>
    </el-row>
  </el-card>

  <el-card :style="{ margin: '30px 50px ' }" class="card-kb">
    <div class="knowledge-Create-head">
      <h2 style="text-align: left;">知识库列表</h2>
      <span style="display: flex ;justify-content: center;align-items: center;" class="search-container2">
        <el-input style="width: 300px;" v-model="searchknowledgeQuery" placeholder="请搜索知识库"
          class="search-input"></el-input>
        <i class="search-icon2">🔍</i>
        <ul id="searchkbResults"></ul>
      </span>
    </div>
    <n-divider title-placement="left"></n-divider>
    <el-row :gutter="12">
      <!-- 创建知识库卡片 -->
      <el-col :span="8">
        <div class="create-card">
          <div class="card-icon" style="cursor: pointer;">
            <el-icon text @click="createknowledgecard(this.createKb)" style="color: #0c89fd">
              <DocumentAdd />
            </el-icon>
          </div>
          <div class="card-createknowledgebase" style="cursor: pointer; justify-content: center; width:100% ;">
            <!-- <span>新建知识库</span> -->
            <el-button text @click="createknowledgecard(this.createKb)"
              style="font-weight: bold;font-size: 18  px;">新建知识库</el-button>
          </div>
          <div class="create-card-description" style="text-align: center;">
            快来创建你的知识库吧！
          </div>
        </div>
      </el-col>
      <el-col :span="8" v-for="(item, index) in paginatedkbs" :key="index">
        <div class="knowledge-card">
          <div class="card-head-row">
            <div class="card-name">
              <div class="card-avatar" :style="generateSequentialColorStyle()">
                {{ item.name[0] }}
              </div>
              <span style="margin-right: 10px;">{{ item.name }}</span>
              <!-- <div style="margin-left: 20px;font-size: small;" :style="tagColor()">{{ item.state }}</div> -->
              <!-- <el-tag>{{ item.state }}</el-tag> -->
            </div>
            <div class="card-dropdown">
              <!-- <el-dropdown  trigger="click"> -->
              <el-dropdown trigger="click">
                <span class="el-dropdown-link" style="font-size: 25px;">
                  ···
                </span>
                <template #dropdown>
                  <el-dropdown-menu>
                    <el-dropdown-item text @click="editKnowledgeDialog(item)">编辑</el-dropdown-item>
                    <el-dropdown-item @click="getFileList(item)">文件</el-dropdown-item>
                    <el-dropdown-item @click="shareKnowledge(item)">分享</el-dropdown-item>
                    <el-dropdown-item @click="deleteKnowledge(item)">删除</el-dropdown-item>
                  </el-dropdown-menu>
                </template>
              </el-dropdown>
            </div>
          </div>
          <div class="card-description">
            详情：{{ item.details }}
          </div>
          <div class="crad-hail-row">
            <div class="card-files" style="justify-content: flex-start;">
              <div style="display: flex;flex-direction: row;">
                <div style="margin-right:10px; ">创建时间:{{ kbtime(item.time) }}</div>
                <div style="margin-right:10px; ">文档数:{{ item.file.length }}</div>
                <div style="margin-right:10px; ">应用数:{{ item.app.length }}</div>
              </div>
            </div>
            <div>
              <el-avatar :icon="UserFilled" />
            </div>
            <!-- <div>
              <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png" />
            </div> -->
          </div>
        </div>
      </el-col>
    </el-row>
    <el-pagination class="knowledgepagination" @current-change="handlePageChangekb" :current-page.sync="currentPagekb"
      :page-size="pageSizekb" :total="filterAllkb.length"></el-pagination>
  </el-card>
  <!-- 上传文件对话框 -->
  <el-dialog v-model="uploadDialog" title="知识库详情" width="75%" style="height: 670px;">
    <div class="table" style="height:110px ;margin-bottom:10px">
      <div style="display: flex;flex-direction: row;">
        <div style="width: 48px;height: 48px;margin-right: 10px;border-radius: 24px;line-height: 46px;"
          :style="generateSequentialColorStyle()">
          {{ current_kb.name[0] }}
        </div>
        <div style="display: flex;flex-direction: column;width: 740px;">
          <div style="display: flex;">
            <div style="font-size: 20px;font-weight: bold">{{ current_kb.name }}</div>
            <el-icon style="font-size: large;  margin-top: 5px;margin-left: 8px;cursor: pointer;" :style="buttonColor()"
              @click="editKnowledgeDialog(this.current_kb)">
              <Edit />
            </el-icon>
          </div>

          <div style="display: flex;margin-top: 6px;">知识库详情：{{ current_kb.details }}</div>
          <!-- <div style="display: flex;margin-top: 6px;">{{ current_kb.details }}</div> -->
          <div style="display: flex;flex-direction: row;margin-top: 6px;">
            <div style="margin-right: 20px;">模型：智谱</div>
            <div>更新时间：2023-11-24</div>
          </div>
        </div>
      </div>
    </div>
    <div class="table">
      <div
        style="display: flex; flex-direction: row; margin-bottom: 10px;margin-left: 15px;margin-right: 15px; justify-content: space-between;">
        <div>文件列表</div>
        <el-button @click="uploadDialog2 = true" style="background-color:#0c89fd;color: aliceblue;"> 上传文件</el-button>
      </div>
      <el-table class="filelisttable" stripe :data="paginatedfiles">
        <el-table-column fixed prop="time" label="时间" width="200" align="center" :formatter="filetime" />
        <el-table-column prop="filename" label="文档名称" width="160" align="center" :formatter="formatfilename" />
        <el-table-column label="文本分段规则" width="160" align="center">
          <template #default="scope">
            {{ methstate(scope.row.meth) }}
          </template>
        </el-table-column>
        <el-table-column prop="maxSegmentLength" label="文本分段最大长度" width="160" align="center" />
        <el-table-column align="right" width="160">
          <template #header>
            <el-input v-model="searchfileListQuery" style="margin-top: 12px;" placeholder="Type to search" />
            <ul id="searchfileResults"></ul>
          </template>
          <template #default="scope" style="justify-content:center">
            <el-popconfirm title="确认删除?" @confirm="deleteFile(scope.row)">
              <template #reference>
                <el-icon style="cursor: pointer;">
                  <Delete />
                </el-icon>
              </template>
            </el-popconfirm>
            <el-icon style="margin-left: 10px;margin-right: 50px;cursor: pointer;" @click="DownloadFile(scope.row)">
              <Download />
            </el-icon>
          </template>
        </el-table-column>
      </el-table>
    </div>
    <el-pagination class="filepagination" v-if="totalPagesfile > 1" @current-change="handlePageChangefile"
      :current-page.sync="currentPagefile" :page-size="pageSizefile" :total="filterAllfile.length"></el-pagination>

  </el-dialog>
  <!--    以下为嵌入上传对话框-->
  <el-dialog v-model="uploadDialog2" title="上传文件" width="75%" style="height: 670px;" after-close="refresh()">
    <el-container>
      <el-aside style="width:360px; margin-top: 15px;">
        <div
          style="color:white; background: linear-gradient(rgb(79, 171, 255),rgb(255,255,255)); border-radius: 20px;height: 320px; padding-top:1px; margin:10px 30px 10px 30px">
          <h3 style="text-align: left;margin: 15px 15px 20px 15px;height: 20px;">帮助小助手</h3>
          <div v-if="uploadfilelabel == 0"
            style="background-color: #e8f1fd; border-radius: 20px; padding: 15px; margin: 0px;height: calc(100% - 85px);text-align: left;">
            <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">文档上传</div>
            <p style="color: black;">单个文档最大支持100M。注意关于文档内容，请尊重和维护公共道德和社会公序，文件中禁止含有违反工序良俗的内容。</p>
            <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">为什么进行文本分段？</div>
            <p style="color: black;">当用户提问时，通过计算问题与文本块的相似度,能够快速找到与问题最匹配的文本段落，让大模型给到更准确且专业的回答。</p>
          </div>
          <div v-if="uploadfilelabel == 1"
            style="background-color: #e8f1fd; border-radius: 20px; padding: 15px; margin: 0px;height: calc(100% - 85px);text-align: left;">
            <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">文本分段规则：</div>
            <p style="color: black;">您可为上传的文档选择合适的分段规则及分段最大长度。注意:不同的分段策略也会对最终知识匹配结果产生影响。</p>
            <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">注意事项：</div>
            <p style="color: black;">当前平台仅提供通用的分段规则，后续平台升级后会提供更多分段策略供您配置,敬请期待!</p>
          </div>
        </div>
      </el-aside>
      <el-main class="createkbcard-upload"
        style=" box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);border-radius: 8px; width: 85%;margin-right: 100px;padding-left: 75px;margin-top: 15px;">
        <n-upload ref="upload" :default-upload="false" drag accept=".txt" :max="1" @change="handleChange"
          style="text-align: left;">
          <div class="upload-file">
            <div style="display: flex; flex-direction: row; margin-bottom: 10px;">
              <div style="color: rgb(116, 176, 249);font-weight: bold;margin-right: 5px;">|</div>
              <div style="color: red; margin-right: 3px; margin-top: 3px;">*</div>
              <div style="text-align: left ; font-weight: bold;">上传文件</div>
            </div>
            <div class="upload-file-icon"
              style="display: flex;flex-direction: column;background-color: rgb(250, 251, 255);width: 550px;">
              <el-icon style="font-size:50px;color: rgb(125, 181, 249);">
                <UploadFilled />
              </el-icon>
              <em style="color: rgb(123, 180, 249);">click to upload</em>
              <div class="el-upload__text" style="font-size: small;">
                支持文件(WORD/PDF/PPT)上传，单个文档最大支持100M
              </div>
            </div>
          </div>
        </n-upload>

        <div class="upload-lcut">
          <div class="upload-lcut-head" style="display: flex; flex-direction: row; margin-bottom: 10px;margin-top: 20px;">
            <div style="color: rgb(116, 176, 249);font-weight: bold;margin-right: 5px;">|</div>
            <div style="color: red; margin-right: 3px; margin-top: 3px;">*</div>
            <div style="text-align: left ; font-weight: bold; ">文档分段设置</div>
            <div style="text-align: left ; margin-left: 10px;">自定义分段规则、分段长度</div>
          </div>
          <el-form :model="filestate">
            <el-form-item class="el-form-item-text" label="标签标题前缀" style="width: 550px;">
              <el-select v-model="filestate.addTitlePrefix" placeholder="是否添加标签标题" style="width: 550px;">
                <el-option label="是" value="0" />
                <el-option label="否" value="1" />
              </el-select>
            </el-form-item>
            <el-form-item class="el-form-item-text" label="文本分段规则" style="width: 550px;">
              <el-select v-model="filestate.meth" placeholder="选择文本分段规则" style="width: 550px;">
                <el-option label="/n" value="0" />
                <el-option label="。" value="1" />
              </el-select>
            </el-form-item>
            <el-form-item class="el-form-item-text" label="分段最大长度" style="width: 550px;">
              <el-input v-model="filestate.maxSegmentLength" placeholder="请输入分段最大长度,在1-2000之间" type="number"
                oninput="value=value.replace(/[^\d.]/g,'');if(value>2000) value=2000;if(value<1)value=1" min="1"
                max="1000" />
            </el-form-item>
          </el-form>
        </div>
      </el-main>
    </el-container>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="uploadDialog2 = false">取消</el-button>
        <el-button type="primary" @click="uploadfileDialog2">完成</el-button>
      </span>
    </template>
  </el-dialog>

  <!-- 创建知识库对话框+步骤条 -->
  <el-dialog v-model="dialogVisible" title="新建知识库" width="75%" :before-close="handleDialogClose"
    style="height: 700px;display: flex;flex-direction: column;justify-content: space-between;">
    <!-- 步骤条 -->
    <el-steps :active="currentStep" align-center>
      <el-step title="知识库信息" description="知识库基本信息" />
      <el-step title="文档配置" description="上传文档,配置文档信息" />
      <el-step title="完成" description="知识库创建完成" />
    </el-steps>
    <!-- 表单或其他内容 -->
    <div v-show="currentStep === 0" class="currentStep0">
      <!-- 第一步内容 -->
      <el-container>
        <el-aside style="width:360px; margin-top: 15px;">
          <div
            style="color:white; background: linear-gradient(rgb(79, 171, 255),rgb(255,255,255)); border-radius: 20px;height: 320px; padding-top:1px; margin:10px 30px 10px 30px">
            <h3 style="text-align: left;margin: 15px 15px 20px 15px;height: 20px;">帮助小助手</h3>
            <div
              style="background-color: #e8f1fd; border-radius: 20px; padding: 15px; margin: 0px;height: calc(100% - 85px);text-align: left;">
              <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">知识库基本信息</div>
              <p style="color: black;">请输入知识库名称，并在描述里面添加知识库的知识范围、应用场景等</p>
              <p style="color: black;">注意：知识库名称不得重复！</p>
            </div>
          </div>
        </el-aside>

        <el-main class="createkbcard">
          <el-form :model="form" style="width: 85%;">
            <div style="font-size: 18px; font-weight: bold; color: #0c89fd;margin-bottom: 30px; ">知识库信息</div>
            <el-form-item label="知识名称：">
              <el-input maxlength="20" v-model="form.name" show-word-limit placeholder="请输入知识库名称,不超过20字" type="text" />
            </el-form-item>
            <el-form-item label="知识描述：">
              <el-input v-model="form.description" maxlength="100" placeholder="请输入知识库描述,不超过100字" show-word-limit
                type="textarea" rows="3" style="height: 75px; " />
            </el-form-item>
            <el-form-item label="向量化模型" style="width: 598px;">
              <el-select v-model="form.model" placeholder="选择模型" style="width: 598px;">
                <el-option label="智谱" value="zhipu" />
              </el-select>
            </el-form-item>
          </el-form>
        </el-main>
      </el-container>
      <!-- <span class="dialog-footer">
          <el-button @click="dialogVisible = false">Cancel</el-button>
          <el-button type="primary" @click="createKnowledge">
            创建
          </el-button>
        </span> -->
    </div>
    <div v-show="currentStep === 1">
      <!-- 第二步内容 -->
      <el-container>
        <el-aside style="width:360px; margin-top: 15px;">
          <div
            style="color:white; background: linear-gradient(rgb(79, 171, 255),rgb(255,255,255)); border-radius: 20px;height: 320px; padding-top:1px; margin:10px 30px 10px 30px">
            <h3 style="text-align: left;margin: 15px 15px 20px 15px;height: 20px;">帮助小助手</h3>
            <div v-if="uploadfilelabel == 0"
              style="background-color: #e8f1fd; border-radius: 20px; padding: 15px; margin: 0px;height: calc(100% - 85px);text-align: left;">
              <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">文档上传</div>
              <p style="color: black;">单个文档最大支持100M。注意关于文档内容，请尊重和维护公共道德和社会公序，文件中禁止含有违反工序良俗的内容。</p>
              <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">为什么进行文本分段？</div>
              <p style="color: black;">当用户提问时，通过计算问题与文本块的相似度,能够快速找到与问题最匹配的文本段落，让大模型给到更准确且专业的回答。</p>
            </div>
            <div v-if="uploadfilelabel == 1"
              style="background-color: #e8f1fd; border-radius: 20px; padding: 15px; margin: 0px;height: calc(100% - 85px);text-align: left;">
              <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">文本分段规则：</div>
              <p style="color: black;">您可为上传的文档选择合适的分段规则及分段最大长度。注意:不同的分段策略也会对最终知识匹配结果产生影响。</p>
              <div style="margin-bottom: 1px;font-size: 18px; font-weight: bold; color: black;">注意事项：</div>
              <p style="color: black;">当前平台仅提供通用的分段规则，后续平台升级后会提供更多分段策略供您配置,敬请期待!</p>
            </div>
          </div>
        </el-aside>
        <el-main class="createkbcard-upload"
          style=" box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);border-radius: 8px; width: 85%;margin-right: 100px;padding-left: 75px;margin-top: 15px;">
          <n-upload ref="upload" :default-upload="false" drag accept=".txt" :max="1" @change="handleChange"
            style="text-align: left;">
            <div class="upload-file">
              <div style="display: flex; flex-direction: row; margin-bottom: 10px;">
                <div style="color: rgb(116, 176, 249);font-weight: bold;margin-right: 5px;">|</div>
                <div style="color: red; margin-right: 3px; margin-top: 3px;">*</div>
                <div style="text-align: left ; font-weight: bold;">上传文件</div>
              </div>
              <div class="upload-file-icon"
                style="display: flex;flex-direction: column;background-color: rgb(250, 251, 255);width: 550px;">
                <el-icon style="font-size:50px;color: rgb(125, 181, 249);">
                  <UploadFilled />
                </el-icon>
                <em style="color: rgb(123, 180, 249);">click to upload</em>
                <div class="el-upload__text" style="font-size: small;">
                  支持文件(WORD/PDF/PPT)上传，单个文档最大支持100M
                </div>
              </div>
            </div>
          </n-upload>

          <div class="upload-lcut">
            <div class="upload-lcut-head"
              style="display: flex; flex-direction: row; margin-bottom: 10px;margin-top: 20px;">
              <div style="color: rgb(116, 176, 249);font-weight: bold;margin-right: 5px;">|</div>
              <div style="color: red; margin-right: 3px; margin-top: 3px;">*</div>
              <div style="text-align: left ; font-weight: bold; ">文档分段设置</div>
              <div style="text-align: left ; margin-left: 10px;">自定义分段规则、分段长度</div>
            </div>
            <el-form :model="filestate">
              <el-form-item class="el-form-item-text" label="标签标题前缀" style="width: 550px;">
                <el-select v-model="filestate.addTitlePrefix" placeholder="是否添加标签标题" style="width: 550px;">
                  <el-option label="是" value="0" />
                  <el-option label="否" value="1" />
                </el-select>
              </el-form-item>
              <el-form-item class="el-form-item-text" label="文本分段规则" style="width: 550px;">
                <el-select v-model="filestate.meth" placeholder="选择文本分段规则" style="width: 550px;">
                  <el-option label="/n" value="0" />
                  <el-option label="。" value="1" />
                </el-select>
              </el-form-item>
              <el-form-item class="el-form-item-text" label="分段最大长度" style="width: 550px;">
                <el-input v-model="filestate.maxSegmentLength" placeholder="请输入分段最大长度,在1-2000之间" type="number"
                  oninput="value=value.replace(/[^\d.]/g,'');if(value>2000) value=2000;if(value<1)value=1" min="1"
                  max="1000" />
              </el-form-item>
            </el-form>
          </div>
        </el-main>
      </el-container>
    </div>

    <div v-show="currentStep === 2">
      <!-- 第三步内容 -->
      <el-card class="create-success-card" shadow="never">
        <template #header>
          <div class="create-success-card-head">
            <el-result icon="success" title="知识库创建完成！">
            </el-result>
          </div>
        </template>
        <div class="table" style="height:160px ;margin-bottom:10px">
          <el-descriptions :columns="1" :border="true">
            <el-descriptions-item label="知识库名称" :span="20" label-class-name="result-label">
              {{ createKb.knowledgename }}
            </el-descriptions-item>
            <el-descriptions-item label="上传文档" :span="20" label-class-name="result-label">
              {{ createKb.filename }}
            </el-descriptions-item>
            <el-descriptions-item label="文档分段设置" :span="20" label-class-name="result-label">
              {{ methstate(createKb.meth) }}
            </el-descriptions-item>
            <el-descriptions-item label="分段最大长度" :span="20" label-class-name="result-label">
              {{ createKb.maxSegmentLength }}
            </el-descriptions-item>
          </el-descriptions>
          <!-- <div style="display: flex;flex-direction: column">
            <div style="display: flex;flex-direction: row;margin-top: 10px;">
              <div>知识库名称：</div>
              <div>{{ createKb.knowledgename }}</div>
            </div>
            <div style="display: flex; margin-top: 6px;margin-top: 10px;">
              <div>文档 </div>
              <div>{{ createKb.filename }}</div>
            </div>
            <div style="display: flex;flex-direction: row;margin-top: 10px;">
              <div style="margin-right: 5px;">文档分段设置 </div>
              <div style="margin-right: 10px;">{{ createKb.meth }}</div>
              <div style="margin-right: 5px;">分段最大长度</div>
              <div >{{ createKb.maxSegmentLength }}</div>
            </div>
          </div> -->
        </div>
      </el-card>
    </div>
    <!-- 对话框底部按钮 -->
    <template #footer>
      <span class="dialog-footer">
        <el-button v-if="currentStep != 2" @click="quitcreate">退出</el-button>
        <el-button type="primary" v-if="currentStep == 1" @click="handlePrevStep">上一步</el-button>
        <el-button type="primary" v-if="currentStep != 2" @click="handleNextStep">{{
          createknowledgecardbutton }}</el-button>
        <div class="create-success-button">
          <el-button type="primary" v-if="currentStep == 2" @click="createKnowledgesuccess">恭喜你，知识库创建成功！</el-button>
        </div>
      </span>
    </template>
  </el-dialog>


  <!-- 编辑知识库对话框 -->
  <el-dialog v-model="dialogVisibleedit" title="编辑知识库" width="50%">
    <!-- <div style="display: flex; flex-direction: row; margin-bottom: 10px;">
      <div style="color: rgb(116, 176, 249);font-weight: bold;margin-right: 5px;">|</div>
      <div style="color: red; margin-right: 3px; margin-top: 3px;">*</div>
      <div style="text-align: left ; font-weight: bold;">更新知识库信息</div>
    </div> -->
    <el-form :model="form">
      <el-form-item label="知识名称：">
        <el-input v-model="form.name" maxlength="20" type="text" show-word-limit></el-input>
      </el-form-item>
      <el-form-item label="知识描述：">
        <el-input v-model="form.description" maxlength="100" type="textarea" rows="3" show-word-limit
          style="height: 75px;" />
      </el-form-item>
      <el-form-item label="向量化模型" style="width: 790px;">
        <el-select v-model="form.model" placeholder="选择模型" style="width: 790px;">
          <el-option label="智谱" value="zhipu" />
        </el-select>
      </el-form-item>
    </el-form>
    <template #footer>
      <span class="dialog-footer">
        <el-button @click="dialogVisibleedit = false">Cancel</el-button>
        <!-- 问题：editKnowledge -->
        <el-button type="primary" @click="editknowledge">
          保存
        </el-button>
      </span>
    </template>
  </el-dialog>
</template>
<style>
.result-label {
  max-width: 150px;
}

.module {
  margin-bottom: 25px;
  /* 可选：底部间距，以分隔行 */
  height: 165px;
  margin-right: 15px;
  margin-left: 10px;
  border: 1px solid #f0f0eb;
  /* 添加样式以定义模块的外观 */
  border-radius: 10px;
  /* 圆角 */
  padding: 10px;
  /* 内边距 */
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  /* 阴影效果 */
  background-color: #f6fafae7;
  /* 背景色 */
}

.table {
  margin-left: 165px;
  margin-right: 165px;
  border: 1px solid #f0f0eb;
  border-radius: 10px;
  padding: 10px;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  background-color: rgb(244, 249, 254);
}

.card-kb {
  height: 600px;
  position: relative;
}

.card-h {
  height: 130px;
}
</style>
<script >

import { mapGetters } from 'vuex'
import { ref } from 'vue'
import { uploadKnowledgeBase } from '@/api/getknowledgelist.js'
import { deleteFile } from "@/api/deleteFile";
import { uploadFile, downloadFile } from "@/api/uploadFile";
import { deletekb } from "@/api/deletekb";
import { editkb } from "@/api/editkb";
import { uploadFilestate } from "@/api/uploadFilestate";
// import { downloadFile } from "@api/downloadFile";
import { Document, DocumentAdd, FolderChecked, School, Search, UserFilled, UploadFilled, Delete, Download, Edit, Right } from '@element-plus/icons-vue';
import { end } from '@popperjs/core';
import { thisTypeAnnotation } from '@babel/types';


export default {

  name: 'KnowledgeBase',
  components: {
    School, Search, UploadFilled, Delete, Download, Right,
    Document, DocumentAdd, FolderChecked, Edit,
  },
  data() {
    return {
      // createkbdetails:{
      //   name:'',
      //   description:''
      // },
      buttoncolor: '',
      tagcolor: '',
      currentPagekb: 1,
      pageSizekb: 5, // 每页显示的数量
      currentPagefile: 1,
      pageSizefile: 6, // 每页显示的数量
      UserFilled: UserFilled,
      createknowledgecardbutton: "下一步",
      inputValue: '',
      colorListIndex: 0,
      currentStep: 0,
      uploadfilelabel: 0,
      searchknowledgeQuery: '',
      searchkbQuery: '',
      searchkbResults: '',
      searchfileListQuery: '',
      searchfileQuery: '',
      searchfileResults: '',
      kbcard: {
        filenumber: '1',
        time: '',
        appnumber: '0'
      },
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
      test: 'test',
      file: [],
      filestate: {
        addTitlePrefix: '',
        maxSegmentLength: '',
        meth: '1'
      },
      dialogVisible: false,
      dialogVisibleedit: false,
      uploadDialog: false,
      uploadDialog2: false,
      form: {
        name: '',
        description: '',
        model: '',
        user: 'admin',
        files: []
      },
      knowledge_upload: {
        files: []
      },
      createKb: {
        knowledgename: '',
        filename: '',
        model: '',
        description: '',
        maxSegmentLength: '',
        meth: '',
        addTitlePrefix: ''
      },
      current_kl: '',
      current_kb: '',
      currentAllkb: [],
      filterAllkb: [],
      currentAllfile: [],
      filterAllfile: [],
    }
  },
  // watch: {  
  //   knowledgeList: {  
  //     handler(newVal, oldVal) {  
  //       this.$store.dispatch('knowledge/getKnowledgeList')  
  //     },  
  //     // deep: true  
  //   }  
  // },  
  computed: {
    ...mapGetters(['test_hello', "knowledgeList", 'fileList']),
    totalPageskb() {
      return Math.ceil(this.filterAllkb.length / this.pageSizekb); // 计算总页数
    },
    totalPagesfile() {
      return Math.ceil(this.filterAllfile.length / this.pageSizefile); // 计算总页数
    },
    paginatedkbs() {
      this.filterAllkb = this.filter_allkb()
      const startIndexkb = (this.currentPagekb - 1) * this.pageSizekb;
      const endIndexkb = startIndexkb + this.pageSizekb;
      return this.filterAllkb.slice(startIndexkb, endIndexkb);
    },
    paginatedfiles() {
      this.filterAllfile = this.filter_allfile()
      const startIndexfile = (this.currentPagefile - 1) * this.pageSizefile;
      const endIndexfile = startIndexfile + this.pageSizefile;
      console.log(this.filterAllfile.slice(startIndexfile, endIndexfile));
      return this.filterAllfile.slice(startIndexfile, endIndexfile);

    },

  },
  mounted() {

    this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
    }).catch(() => {
      console.log('[KnowledgeBase]: 获取知识库失败！')
    })
  },
  methods: {
    refresh() {
      this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
      }).catch(() => {
        console.log('[KnowledgeBase]: 获取知识库失败！')
      })
    },
    openCreateDialog() {
      this.dialogVisible = true
      this.form.name = ''
      this.form.description = ''
      this.form.model = ''
    },
    shareKnowledge() {
      console.log("分享")
    },
    handleDialogClose() {
      this.$confirm('是否退出创建？信息将不会保留！如果你想暂存输入信息，请点击退出按钮！', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        this.currentStep = 0;
        this.dialogVisible = false;
      }).catch(() => {
        console.log('用户取消退出');
      });
    },
    createKnowledgesuccess() {
      this.dialogVisible = false
      this.currentStep = 0
      this.createKb = {};
      // this.createKb.filename = ''
      // this.createKb.knowledgename = ''
      // this.createKb.maxSegmentLength = ''
      // this.createKb.meth = ''
      // this.createKb.addTitlePrefix=''
      // this.createKb.model=''
    },
    handlePageChangekb(newPagekb) {
      console.log(this.currentPagekb)
      this.currentPagekb = newPagekb;
    },
    handlePageChangefile(newPagefile) {
      console.log(this.currentPagefile)
      this.currentPagefile = newPagefile;
    },
    methstate(e) {
      if (e == '0') {
        return "/n"
      }
      if (e == '1') {
        return "。"
      }
      else {
        return "未定义！"
      }
    },
    DownloadFile(e) {
      downloadFile(e._id).then(response => {
        console.log(response)
        const fileContent = response.data;
        const blob = new Blob([fileContent], { type: 'text/plain' });
        const url = URL.createObjectURL(blob);
        const link = document.createElement('a');
        link.href = url;
        link.download = e.filename;
        document.body.appendChild(link);
        link.click();
      }
      )
    },

    createknowledgecard(data) {
      this.dialogVisible = true
      // console.log("1111")
      this.form.name = data.knowledgename
      this.form.description = data.description
      this.form.model = data.model
      this.filestate.addTitlePrefix = data.addTitlePrefix
      this.filestate.maxSegmentLength = data.maxSegmentLength
      this.filestate.meth = data.meth
    },
    handleNextStep() {
      this.uploadfilelabel = 0
      if (this.currentStep == 0) {
        // this.createkbdetails.name = this.form.name
        // this.createkbdetails.description = this.form.description
        if (this.form.name == '' || this.form.description == '') {
          this.$message.warning("知识库名称或知识库描述不能为空！")
          return
        }
        this.createknowledgecardbutton = "创建";
      }
      if (this.currentStep == 1) {
        if (this.file.length == 0) {
          this.$message.warning("请上传知识库文件！")
          //暂存信息
          // this.form.name = this.createkbdetails.name
          // this.form.description = this.createkbdetails.description
          return

        }
        const textForm = {
          name: this.form.name,
          details: this.form.description,
          // model:this.form.model
          // 模型还未上传
        }
        this.createKb.knowledgename = textForm.name
        this.createKb.description = textForm.details
        this.createKb.model = 'zhipu'
        // this.createKb.model = textForm.model
        uploadKnowledgeBase(textForm).then(res => {
          if (res.data) {
            this.createknowledgecardbutton = "下一步";
            this.currentStep++;
            const formData = new FormData();
            formData.append("file", this.file[0].file);
            this.createKb.filename = this.file[0].file.name
            console.log(this.file[0].file)
            uploadFile(formData, res.data.data._id).then(fileresponse => {
              console.log(fileresponse)
              const filestate = {
                addTitlePrefix: this.filestate.addTitlePrefix,
                meth: this.filestate.meth,
                maxSegmentLength: this.filestate.maxSegmentLength
              }
              console.log(filestate)
              this.createKb.maxSegmentLength = this.filestate.maxSegmentLength
              this.createKb.meth = this.filestate.meth
              this.createKb.addTitlePrefix = this.filestate.addTitlePrefix
              uploadFilestate(filestate, fileresponse.data.data._id).then(response => {
                console.log(response)
                if (response.data) {
                  this.$refs.upload.clear()
                  this.file = []
                  this.filestate = {
                    addTitlePrefix: '',
                    maxSegmentLength: '',
                    meth: '1'
                  }
                  this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
                    // console.log(res)
                  })
                }
              }
              ).catch((err) => {
                this.$message.error("上传失败！")
              })
            })

          }
        }).catch((err) => {
        })
      }

      if (this.currentStep < 1) {
        this.currentStep++;
      }
    },

    uploadfileDialog2() {
      if (this.file.length == 0) {
        this.$message.warning("请上传知识库文件！")
        return
      }
      this.uploadfilelabel = 0
      const formData = new FormData();
      formData.append("file", this.file[0].file);
      uploadFile(formData, this.current_kl).then(fileresponse => {
        this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
        }).catch(() => {
          console.log('[KnowledgeBase]: 获取知识库失败！')
        })
        console.log(fileresponse)
        const filestate = {
          addTitlePrefix: this.filestate.addTitlePrefix,
          meth: this.filestate.meth,
          maxSegmentLength: this.filestate.maxSegmentLength
        }
        console.log(filestate)
        uploadFilestate(filestate, fileresponse.data.data._id).then(response => {
          console.log(response)
          if (response.data) {
            this.$refs.upload.clear()
            this.file = []
            this.filestate = {
              addTitlePrefix: '',
              maxSegmentLength: '',
              meth: '1'
            }
            this.$store.dispatch('file/getFileListByKBId', this.current_kl).then((res) => {
              // console.log(res)
            }).catch(() => {
              console.log('[KnowledgeBase]: 获取知识库失败！')
            })
          }
        }
        ).catch((err) => {
          console.log(err)
          this.$message.error("删除失败！")
        })
      })
      this.uploadDialog2 = false
    },
    handlePrevStep() {
      if (this.currentStep == 1) {
        this.createknowledgecardbutton = "下一步";
      }
      if (0 < this.currentStep && this.currentStep <= 2) {
        this.currentStep--;
      }
    },

    filetime(row, column, e) {
      return e.substring(0, 10) + " " + e.substring(12, 19);
    },
    kbtime(e) {
      return e.substring(0, 10) + " " + e.substring(12, 19);
    },
    formatfilename(row, column, e) {
      // let filenameformat = e.split('.'); 
      // return filenameformat[0]; 
      return e
    },
    generateSequentialColorStyle() {
      // this.colorListIndex = 0
      const color = this.colorList[this.colorListIndex];
      this.colorListIndex = (this.colorListIndex + 1) % this.colorList.length;
      this.buttoncolor = color.background;
      this.tagcolor = color.color
      // 返回颜色样式
      return {
        'background-color': color.background,
        'color': color.color
      };
    },
    buttonColor() {
      return {
        'color': this.buttoncolor
      }
    },
    tagColor() {
      return {
        'color': this.tagcolor
      }
    },
    createKnowledge() {
      // this.dialogVisible = false
      const textForm = {
        name: this.form.name,
        details: this.form.description
      }
      uploadKnowledgeBase(textForm).then(res => {
        // console.log(res)
        if (res.data) {
          this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
            // console.log(res)
          }).catch(() => {
            console.log('[KnowledgeBase]: 获取知识库失败！')
          })
        }
      })
    },
    filter_allkb() {
      return this.knowledgeList.filter(filterAllkb => filterAllkb.name.includes(this.searchknowledgeQuery));
    },
    filter_allfile() {
      return this.fileList.filter(filterAllfile => filterAllfile.filename.includes(this.searchfileListQuery));
    },
    // 上传文件
    // submitUpload() {
    //   this.uploadDialog2 = false
    //   // this.file = ref([])
    //   console.log(this.current_kl)
    //   const formData = new FormData();
    //   formData.append("file", this.file[0].file);
    //   uploadFile(formData, this.current_kl).then(response => {
    //     if (response.data) {
    //       // console.log('111')
    //       this.$refs.upload.clear()
    //       this.file = []
    //       this.$store.dispatch('file/getFileListByKBId', this.current_kl).then((res) => {
    //         // console.log(res)
    //       }).catch(() => {
    //         console.log('[KnowledgeBase]: 获取知识库知识失败！')
    //       })
    //       // this.$store.dispatch('file/getFileListByKBId').then(res=>{
    //       // })
    //     }
    //   })
    // },
    handleChange(data) {
      this.file = data.fileList
      this.uploadfilelabel = 1
    },

    // 问题：editKnowledge
    editknowledge() {
      this.dialogVisibleedit = false
      const textForm = {
        _id: this.form._id,
        name: this.form.name,
        details: this.form.description
      }
      this.current_kb = textForm
      // console.log(textForm)
      editkb(textForm).then(res => {
        if (res.data) {
          this.$store.dispatch('knowledge/getKnowledgeList').then((res) => {
          }).catch(() => {
            console.log('[KnowledgeBase]: 获取知识库失败！')
          })
        }
        this.$message.success('更新知识库成功！')
        console.log(res)
      })
    },

    editKnowledgeDialog(e) {
      this.dialogVisibleedit = true
      console.log(e)
      this.form._id = e._id
      this.form.name = e.name
      this.form.description = e.details
      this.form.model = 'zhipu'
    },


    // 获取文件
    getFileList(e) {
      console.log(e)
      this.current_kb = e
      this.uploadDialog = true
      this.current_kl = e._id
      this.$store.dispatch('file/getFileListByKBId', this.current_kl).then((res) => {

        console.log(this.file.length)
      }).catch(() => {
        console.log('获取文件列表失败！')
      })
    },
    quitcreate() {
      this.$confirm('是否需要保存该条记录吗？', '提示', {
        distinguishCancelAndClose: true,
        confirmButtonText: '保存',
        cancelButtonText: '不保存',
        type: 'warning'
      }).then(() => {
        this.createKb.knowledgename = this.form.name
        this.createKb.description = this.form.description
        this.createKb.model = this.form.model
        this.createKb.addTitlePrefix = this.filestate.addTitlePrefix
        this.createKb.maxSegmentLength = this.filestate.maxSegmentLength
        this.createKb.meth = this.filestate.meth
        console.log(this.createKb)
        this.dialogVisible = false
      }).catch(action => {
        if (action == 'cancel') {
          this.currentStep = 0
          this.createknowledgecardbutton = "下一步"
          this.createKb = {}
          this.dialogVisible = false;
        }
        if (action == 'close') {
          console.log('取消退出！')
        }
      })
    },
    //bug 取消删除，报错 已解决
    deleteKnowledge(e) {
      if(e.app.length=='0'){
        this.$confirm('确认删除该条记录吗？', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      }).then(() => {
        // 用户点击了确定按钮
        // 调用 deletekb 方法删除知识记录
        deletekb(e._id).then(res => {
          console.log(res);
          if (res.data) {
            // 如果删除成功，更新知识列表
            this.$store.dispatch('knowledge/getKnowledgeList').then(() => {
              // 处理删除成功后的逻辑
            }).catch(() => {
              console.log('[KnowledgeBase]: 删除失败！');
            });
          }
        });
      }).catch(() => {
        console.log('用户取消删除');
      });
      }
      else {
        this.$message.success("当前知识库已绑定应用，暂不能删除！")
      }
      // 使用 Element UI 的 $confirm 方法显示确认对话框
    },
    deleteFile(e) {
      deleteFile(e._id).then((res) => {
        this.$message.success("删除成功！")
        this.$store.dispatch('file/getFileListByKBId', this.current_kl).then((res) => {
          console.log(res)
        }).catch(() => {
          console.log('[KnowledgeBase]: 获取知识库失败！')
        })
      }).catch((err) => {
        this.$message.error("删除失败！")
      })
      // console.log("删除文件",e._id)
    }

  },
  props: {
    msg: String
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.create-success-card {
  border: 0;
}

h3 {
  margin: 40px 0 0;
}

ul {
  list-style-type: none;
  padding: 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

a {
  color: rgb(66, 137, 185);
}

/* 步骤条card样式 */


.knowledge-Create-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

/* 创建知识库css */
.create-success-button {
  display: flex;
  justify-content: center;
  align-items: center;
}

.createkbcard {
  display: flex;
  justify-content: center;
  align-items: center;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  border-radius: 8px;
  margin-top: 15px;
  margin-right: 100px;
}

.create-card {
  width: 80%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: start;
  margin: 10px;
  padding: 20px;
  border: 1px solid #f0f0eb;
  /* 添加样式以定义模块的外观 */
  border-radius: 10px;
  /* 圆角 */
  background-color: #f6fafae7;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  height: 150px;
}

.card-icon {
  width: 100%;
  margin-bottom: 10px;
  align-items: center;
  font-size: 40px;
}

.card-createknowledgebase {
  margin-bottom: 10px;
  cursor: pointer;
}

.card-createbutton {
  cursor: pointer;
  background-color: #c1dfff;
  border: none;
  color: black;
}

.card-createbutton:hover {
  background-color: #dddddd;
}

.create-card-description {
  font-size: 14px;
  width: 100%;
  align-items: center;
}

.crad-hail-row {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.filelisttable {
  width: 100%;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: auto;
}

/* 展示知识库css */
.knowledge-card {
  width: 80%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: start;
  margin: 10px;
  padding: 20px;
  border: 1px solid #f0f0eb;
  /* 添加样式以定义模块的外观 */
  border-radius: 10px;
  /* 圆角 */
  background-color: #f6fafae7;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
  /* 背景色 */
  height: 150px;
}

.card-head-row {
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
}

.card-head-row .card-name {
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  font-size: large;
  font-weight: bold;
  font-family: sans-serif;
}

.card-name .card-avatar {
  width: 48px;
  height: 48px;
  background: RGB(214, 234, 253);
  color: RGB(96, 178, 254);
  margin-right: 10px;
  border-radius: 24px;
  line-height: 46px;
}

.filepagination {
  position: absolute;
  bottom: 15px;
}

.knowledgepagination {
  position: absolute;
  bottom: 10px;
}

.card-description {
  margin: 10px 0;
  color: gray;
  font-size: 16px;
  max-width: 350px;
  /* 或任何你需要的长度 */
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
}

.card-dropdown {
  cursor: pointer;
}

.card-dropdown:hover {
  background: #dddddd;
}

.card-files {
  display: flex;
  width: 100%;
  flex-direction: row;
  font-size: 14px;
  display: flex;
  justify-content: flex-start;
  /* align-items: center; */
}

.card-files .card-file {

  text-align: left;
  cursor: pointer;
  padding: 3px;
  margin-right: 5px;
  border-radius: 5px;

}

.upload-file-icon {
  cursor: pointer;
  width: 450px;
  height: 100px;
  display: flex;
  /* 使用flexbox布局 */
  justify-content: center;
  /* 水平居中 */
  align-items: center;
  /* 垂直居中 */
  border: 1px dashed #000;
  /* 虚线边框 */
  border-radius: 20px;
  /* 增加20度圆角 */
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1) when (hover) scale(1.1);
  /* 鼠标悬停时的阴影效果 */
}

.card-files .card-file:hover {
  background: #dddddd;
}
</style>
