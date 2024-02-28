import { editApp } from '@/api/editApp'

const getDefaultState = () => {
  return {
    edit: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_EDIT: (state, edit) => {
    state.edit = edit
  }
}

const actions = {
  // get log list
  editApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      editApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_EDIT', list)
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
