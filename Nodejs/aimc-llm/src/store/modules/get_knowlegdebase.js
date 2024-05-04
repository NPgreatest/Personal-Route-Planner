import { getKnowledgeList } from '@/api/get_konwledgebase'

const getDefaultState = () => {
  return {
    knowledge:[]
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_KNOWLEDGE: (state, knowledge) => {
    state.knowledge = knowledge
  }
}

const actions = {
  // get log list
  getKnowledgeList({ commit }) {
    return new Promise((resolve, reject) => {
      getKnowledgeList().then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_KNOWLEDGE', list)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
