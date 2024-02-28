import { submitPrompt } from '@/api/submitPrompt'

const getDefaultState = () => {
  return {
    prompt: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_PROMPT: (state, prompt) => {
    state.prompt = prompt
  }
}

const actions = {
  // get log list
  submitPrompt({ commit }, params) {
    return new Promise((resolve, reject) => {
      submitPrompt(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_PROMPT', list)
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
