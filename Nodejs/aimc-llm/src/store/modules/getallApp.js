import { getallApp } from '@/api/getallApp'

const getDefaultState = () => {
  return {
    allapp:[]
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_ALLAPP: (state, allapp) => {
    state.allapp = allapp
  }
}

const actions = {
  // get log list
  getallApp({ commit }) {
    return new Promise((resolve, reject) => {
      getallApp().then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_ALLAPP', list)
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