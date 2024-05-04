import { getApp } from '@/api/getApp'

const getDefaultState = () => {
  return {
    app: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_APP: (state, app) => {
    state.app = app
  }
}

const actions = {
  // get log list
  getApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      getApp(params).then(response => {
        console.log(response.data)
        const list = response.data.app
        commit('UPDATE_APP', list)
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
