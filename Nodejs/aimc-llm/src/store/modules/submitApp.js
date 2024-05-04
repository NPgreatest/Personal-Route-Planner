import { submitApp } from '@/api/submitApp'

const getDefaultState = () => {
  return {
    submit: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_SUBMIT: (state, submit) => {
    state.submit = submit
  }
}

const actions = {
  // get log list
  submitApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      submitApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_SUBMIT', list)
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
