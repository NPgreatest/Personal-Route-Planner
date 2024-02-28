import { deleteApp } from '@/api/deleteApp'

const getDefaultState = () => {
  return {
    deleteapp:[]
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_DELETEAPP: (state, deleteapp) => {
    state.deleteapp = deleteapp
  }
}

const actions = {
  // get log list
  deleteApp({ commit },params) {
    return new Promise((resolve, reject) => {
      deleteApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_DELETEAPP', list)
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