import { stoptestApp } from '@/api/stoptestApp'

const getDefaultState = () => {
  return {
    stoptestapp: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_STOPTESTAPP: (state, stoptestapp) => {
    state.stoptestapp = stoptestapp
  }
}

const actions = {
  // get log list
  stoptestApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      stoptestApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_STOPTESTAPP', list)
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
