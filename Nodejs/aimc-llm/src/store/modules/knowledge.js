import { getKnowledgeBase } from '@/api/getknowledgelist'
const getDefaultState = () => {
  return {
    knowledgeList: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_KNOWLEDGE: (state, KnowledgeList) => {
    state.knowledgeList = KnowledgeList;
  }
}

const actions = {
  // get log list
  getKnowledgeList ({ commit }) {
    return new Promise((resolve, reject) => {
      getKnowledgeBase().then(response => {
        // console.log(response.data)
        const KnowledgeList = response.data
        console.log({ KnowledgeList })
        // state.knowledgeList = KnowledgeList
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
