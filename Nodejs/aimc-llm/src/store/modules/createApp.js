import { createApp } from '@/api/createApp'

const getDefaultState = () => {
  return {
    create: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_CREATE: (state, create) => {
    state.create = create
  }
}

const actions = {
  // get log list
  createApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      createApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_CREATE', list)
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
