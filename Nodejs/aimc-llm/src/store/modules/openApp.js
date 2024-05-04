import { openApp } from '@/api/openApp'

const getDefaultState = () => {
  return {
    openapp: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_OPENAPP: (state, openapp) => {
    state.openapp = openapp
  }
}

const actions = {
  // get log list
  openApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      openApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_OPENAPP', list)
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
