import { testApp } from '@/api/testApp'

const getDefaultState = () => {
  return {
    testapp: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_TESTAPP: (state, testapp) => {
    state.testapp = testapp
  }
}

const actions = {
  // get log list
  testApp({ commit }, params) {
    return new Promise((resolve, reject) => {
      testApp(params).then(response => {
        console.log(response.data)
        const list = response.data
        commit('UPDATE_TESTAPP', list)
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
