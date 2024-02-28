import { getKnowledgeBase } from '@/api/getknowledgelist'
const getDefaultState = () => {
  return {
    knowledgeList: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_KNOWLEDGE: (state, KnowledgeList) => {
    state.knowledgeList = KnowledgeList
  }
}

const actions = {
  // get log list
  getKnowledgeList({ commit }) {
    return new Promise((resolve, reject) => {
      getKnowledgeBase().then(response => {
        console.log(response.data)
        const KnowledgeList = response.data
        for(let i=0;i<KnowledgeList.length;i++) {
          KnowledgeList[i].label=KnowledgeList[i].name[0] + KnowledgeList[i].name[1]
        }
        console.log(KnowledgeList)
        commit('UPDATE_KNOWLEDGE', KnowledgeList)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
