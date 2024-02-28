import { stopApp } from '@/api/stopApp'

const getDefaultState = () => {
  return {
    stopapp: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_STOPAPP: (state, stopapp) => {
    state.stopapp = stopapp
  }
}

const actions = {
  // get log list
  stopApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      stopApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_STOPAPP', list)
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
