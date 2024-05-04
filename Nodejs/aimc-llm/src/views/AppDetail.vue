<template>
    <div style="text-align: left;margin-left: 50px;margin-bottom: 10px;"><el-button plain type="primary"
            @click="this.$router.push('/about')"> 返回应用列表</el-button>
    </div>
    <el-card :style="{ margin: '20px 50px' }" style="height: 173px;text-align: left;">
        <div style="height: 25px;">
            <el-text tag="b">{{ app.name }}</el-text>
            <el-tag style="margin-left: 5px;">{{ app.state }}</el-tag>
            <el-button type="primary" @click="showpublic" style="margin-left: 1100px;" v-if="app.state != '已发布'">
                发布
            </el-button>
        </div>
        <p style="margin-bottom: 3px;">{{ app.details }}</p>
        <div>
            <p style="display: inline-block;margin-top: 15px;">调用次数：{{ app.num }}</p>
            <p style="display: inline-block;margin-left: 10px;margin-right: 10px;">|</p>
            <p style="display: inline-block;">创建时间：{{ app.time }}</p>
        </div>
        <div>
            <el-button text @click="flag1 = true; flag2 = false;">应用详情</el-button>
            <el-button text @click="showprompt">提示词编排</el-button>
        </div>
    </el-card>
    <div v-if="flag1">
        <el-card :style="{ margin: '20px 50px' }" style="height: 130px;text-align: left;">
            <el-text tag="b">应用信息</el-text>
            <div>
                <p style="display: inline-block;margin-right: 50px;margin-bottom: 1px;">应用名称：{{ app.name }}</p>
                <p style="display: inline-block;margin-right: 50px;margin-bottom: 1px;">创建人：张三</p>
                <p style="display: inline-block;margin-right: 10px;margin-bottom: 1px;">应用状态：</p>
                <el-switch v-if="app.state === '已发布'" v-model="isSwitchOn" :disabled="!isSwitchOn"
                    @change="handleSwitchChange" size="large" inline-prompt
                    style="--el-switch-on-color: #13ce66; --el-switch-off-color: #ff4949" active-text="已发布"></el-switch>
                <el-tag v-if="app.state === '未发布'" type="info" effect="dark" round style="margin-bottom: 1px;">未发布</el-tag>
                <el-tag v-if="app.state === '调试中'" type="info" effect="dark" round style="margin-bottom: 1px;">调试中</el-tag>
                <p style="display: inline-block;margin-left: 50px;margin-bottom: 1px;">创建时间：{{ app.time }}</p>
            </div>
            <div>
                <p style="display: inline-block;margin-right: 85px;margin-top: 6px;">应用调用次数：{{ app.num }}</p>
                <p style="display: inline-block;margin-right: 60px;margin-top: 6px;">提示词：{{ app.prompt }}</p>
            </div>
        </el-card>
        <div v-if="app.state === '已发布'">
            <el-card :style="{ margin: '20px 50px' }" style="height: 100px;text-align: left;">
                <el-text tag="b">发布信息</el-text>
                <div>
                    <p style="display: inline-block;margin-right: 85px;margin-top: 6px;">发布大模型：
                        {{ model_list.find(item => item._id === app.model).name }} </p>
                    <p style="display: inline-block; margin-right: 60px; margin-top: 6px;">
                        应用发布URL：
                        <a href="{{ app.link }}" target="_blank">{{ app.link }}</a>
                    </p>
                    <p style="display: inline-block;margin-right: 60px;margin-top: 6px;">提示词：{{ app.prompt }}</p>
                </div>
            </el-card>
        </div>
        <el-card :style="{ margin: '20px 50px' }" style="height: 120px;text-align: left;">
            <el-text tag="b">知识库信息</el-text>
            <div>
                <p v-for="app in selectKnowledgebase" style="margin-right: 85px;margin-top: 18px;display: inline-block;">{{
                    app }}</p>
            </div>
        </el-card>
    </div>
    <div v-if="flag2" style="display: flex;">
        <el-card style="height: 410px;text-align: left;flex: 1;margin-left: 50px;margin-right: 50px;">
            <div style="height: 30px;">
                <el-text tag="b">对话前置提示词</el-text>
                <el-tooltip class="box-item" effect="light"
                    content="（Prompt）是一种用于指导人工智能模型生成输出的小段文本，您可以使用提示词来精确地描述希望模型执行的任务，可以包含一些回答示例，或者提供有关输入数据的一些背景信息、输出要求等。"
                    placement="top-start">
                    <svg @mouseover="showPopup = true" @mouseleave="showPopup = false" version="1.1"
                        xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="20px" y="20px"
                        width="20" height="20" viewBox="0 0 512 512" enable-background="new 0 0 512 512"
                        xml:space="preserve">
                        <g>
                            <g>
                                <g>
                                    <g>
                                        <path
                                            d="M256,76c48.1,0,93.3,18.7,127.3,52.7S436,207.9,436,256s-18.7,93.3-52.7,127.3S304.1,436,256,436
                            c-48.1,0-93.3-18.7-127.3-52.7C94.7,349.3,76,304.1,76,256s18.7-93.3,52.7-127.3C162.7,94.7,207.9,76,256,76 M256,48
                            C141.1,48,48,141.1,48,256s93.1,208,208,208c114.9,0,208-93.1,208-208S370.9,48,256,48L256,48z">
                                        </path>
                                    </g>
                                </g>
                            </g>
                            <g>
                                <path d="M256.7,160c37.5,0,63.3,20.8,63.3,50.7c0,19.8-9.6,33.5-28.1,44.4c-17.4,10.1-23.3,17.5-23.3,30.3v7.9h-34.7l-0.3-8.6
                    c-1.7-20.6,5.5-33.4,23.6-44c16.9-10.1,24-16.5,24-28.9s-12-21.5-26.9-21.5c-15.1,0-26,9.8-26.8,24.6H192
                    C192.7,182.7,216.5,160,256.7,160z M230.4,331.4c0-11.5,9.6-20.6,21.4-20.6c11.9,0,21.5,9,21.5,20.6s-9.6,20.6-21.5,20.6
                    C239.9,352,230.4,343,230.4,331.4z"></path>
                            </g>
                        </g>
                    </svg>
                </el-tooltip>
                <el-button v-if="app.state === '未发布'" type="primary" style="margin-left: 1050px;"
                    @click="testprompt">提示词测试</el-button>
                <el-button v-if="app.state === '调试中'" type="primary" style="margin-left: 1000px;"
                    @click="stoptestprompt">关闭测试</el-button>
                <el-button v-if="app.state === '调试中'" type="primary" style="margin-left: 30px;"
                    @click="this.drawer = true">继续测试</el-button>
            </div>
            <el-form :model="form2" label-width="120px" style="margin-top: 20px;">
                <el-form-item label="提示词输入" required>
                    <el-input v-model="form2.prompt" style="height: 290px;width: 600px;" />
                </el-form-item>
                <div style="margin-top: 10px;margin-left: 118px;">
                    <el-button type="primary" @click="submitprompt">确定</el-button>
                </div>
            </el-form>
            <el-drawer v-model="drawer" title="提示词测试" direction='btt' size="690">
                <iframe :src="app.link" frameborder="0" style="width: 100%; height: 700px"></iframe>
            </el-drawer>

        </el-card>
    </div>
    <el-dialog v-model="flag3" title="发布应用">
        <el-form :model="form" label-width="120px" v-if="flag3_1">
            <el-form-item label="大模型选择" required>
                <el-select v-model="form.model" placeholder="选择大模型" style="width: 300px;">
                    <el-option v-for="app in model_list" :label="app.name" :key="app._id" :value="app._id" />
                </el-select>
            </el-form-item>
            <div style="margin-top: 20px;margin-left: 108px;">
                <el-button type="primary" plain @click="flag3 = 0" style="margin-left: 10px;">取消</el-button>
                <el-button type="primary" @click="submitapp">确定</el-button>
            </div>
        </el-form>
        <n-result status="info" title="发布中" description="应用发布中，请稍候" v-if="flag3_2"></n-result>

    </el-dialog>
    <el-dialog v-model="flag4">
        <img alt="" src="../assets/images.jpg" style="float: center;">
        <h2>应用发布完成！</h2>
        <div style="margin-top: 20px;">
            <el-button type="primary" plain @click="flag4 = false" style="margin-left: 10px;">返回-></el-button>
            <el-button type="primary" @click="alertapp">去试用-></el-button>
        </div>
    </el-dialog>
</template>

<script>
import { ref } from 'vue';
import { mapGetters } from 'vuex'
import { ElNotification } from 'element-plus'

export default {
    data() {
        return {
            model_list: [
                { '_id': '654dd9c4ea13fc974c0504c9', 'name': 'openai_gpt_3.5_turbo' },
            ],
            showPopup: "（Prompt）是一种用于指导人工智能模型生成输出的小段文本，您可以使用提示词来精确地描述希望模型执行的任务，可以包含一些回答示例，或者提供有关输入数据的一些背景信息、输出要求等。",
            flag1: true,
            flag2: false,
            flag3: false,
            flag3_1: false,
            flag3_2: false,
            flag4: false,
            drawer: false,
            form: {
                _id: this.$route.params._id,
                model: '',
            },
            form2: {
                _id: this.$route.params._id,
                prompt: ""
            },
            isSwitchOn: true
        }
    },
    computed: {
        ...mapGetters(['knowledge_list', 'app', 'submit', 'openapp', 'testapp', 'prompt']),
        selectKnowledgebase() {
            if (!this.app.knowledgebase || this.app.knowledgebase.length === 0) {
                return ['无知识库'];
            };
            return this.app.knowledgebase.map(id => {
                const item = this.knowledge_list.find(item => item._id === id);
                return item ? item.name : '未找到';
            })
        },
    },
    async mounted() {
        const id = this.$route.params._id;

        try {
            this.$store.dispatch('get_knowledgebase/getKnowledgeList');
        } catch (err) {
            console.log('[HomeView]: 获取知识库失败！');
        }

        try {
            await this.$store.dispatch('getApp/getApp', { _id: id });
            this.form.model = this.app.model
            this.form2.prompt = this.app.prompt;

        } catch (err) {
            console.log('[HomeView]: 获取应用失败！');
        }
    },
    methods: {
        async alertapp() {
            this.flag4 = false;
            window.open(this.app.link, '_blank');

        },
        showpublic() {
            if (this.app.state === '调试中') { alert('请先关闭测试再进行发布'); }
            else { this.flag3 = true; this.flag3_1 = true }
        },
        showprompt() {
            if (this.app.state === '已发布') { alert('已发布的应用不支持提示词编排功能'); }
            else { this.flag1 = false; this.flag2 = true; }
        },
        stoptestprompt() {
            const id = this.$route.params._id;
            this.$store.dispatch('stoptestApp/stoptestApp', id).then((res) => {
                this.$store.dispatch('getApp/getApp', { _id: id });
            }).catch(() => {
                console.log('[HomeView]: 获取hello失败！')
            });
        },
        testprompt() {
            ElNotification({
                title: '环境创建中',
                message: '测试环境创建中，请稍候',
                type: 'success',
            });
            const id = this.$route.params._id;
            this.$store.dispatch('testApp/testApp', id).then((res) => {
                this.drawer = true
                this.$store.dispatch('getApp/getApp', { _id: id });
            }).catch(() => {
                console.log('[HomeView]: 获取hello失败！')
            });

        },
        async submitapp() {
            this.flag3_1 = false
            this.flag3_2 = true
            const id = this.$route.params._id;
            console.log(this.form)
            await this.$store.dispatch('submitPrompt/submitPrompt', this.form).then((res) => {
            }).catch(() => {
                console.log('[HomeView]: 获取hello失败！')
            });
            console.log(this.form2)
            await this.$store.dispatch('openApp/openApp', this.form2).then((res) => {
                console.log('发布成功')
                this.flag3 = false
                this.flag3_2 = false
                this.flag4 = true;
            }).catch(() => {
                console.log('[HomeView]: 获取hello失败！')
            });
            this.$store.dispatch('getApp/getApp', { _id: id });
        },
        submitprompt() {
            const id = this.$route.params._id;
            this.$store.dispatch('submitPrompt/submitPrompt', this.form2).then((res) => {
                ElNotification({
                    title: '已提交',
                    message: '提示词已提交',
                    type: 'success',
                });
                this.$store.dispatch('getApp/getApp', { _id: id });
            }).catch(() => {
                console.log('[HomeView]: 获取hello失败！')
            });
        },
        handleSwitchChange(newVal) {
            if (!newVal) {
                // 开关尝试关闭时，弹出确认弹框
                this.$confirm('确认要关闭吗？', '关闭确认', {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: 'warning'
                }).then(() => {
                    // 确认关闭
                    this.$store.dispatch('stopApp/stopApp', this.$route.params._id).then((res) => {
                        this.$router.go(0)
                    }).catch(() => {
                        console.log('[HomeView]: 获取hello失败！')
                    });
                }).catch(() => {
                    // 取消关闭，保持原状态
                    this.isSwitchOn = true
                });
            }
        }
    }

}

</script>
