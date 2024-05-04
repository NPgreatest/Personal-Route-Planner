import Vue from 'vue'
import Vuex from 'vuex'
import getters from './getters'
import createApp from "@/store/modules/createApp"
import editApp from "@/store/modules/editApp"
import deleteApp from "@/store/modules/deleteApp"
import get_knowledgebase from "@/store/modules/get_knowlegdebase"
import getallApp from "@/store/modules/getallApp"
import getApp from "@/store/modules/getApp"
import submitApp from "@/store/modules/submitApp"
import openApp from "@/store/modules/openApp"
import stopApp from "@/store/modules/stopApp"
import testApp from "@/store/modules/testApp"
import submitPrompt from "@/store/modules/submitPrompt"
import stoptestApp from "@/store/modules/stoptestApp"
import knowledge from "@/store/modules/knowledge";
import file from "@/store/modules/file";

// Vue.use(Vuex)

const store = new Vuex.Store({
  modules: {
    createApp,
    editApp,
    deleteApp,
    get_knowledgebase,
    getallApp,
    getApp,
    knowledge,
    submitApp,
    openApp,
    stopApp,
    testApp,
    submitPrompt,
    stoptestApp,
    file
  },
  getters
})

export default store
