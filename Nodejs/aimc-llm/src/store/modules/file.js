import {uploadFile, getFileListById} from "@/api/serviceFile";
import { getFileListByKBId } from "@/api/uploadFile"; 

const getDefaultState = () => {
  return {
    fileList: []
  }
}

const state = getDefaultState()


const mutations = {
  UPDATE_FILELIST: (state, fileList) => {
    state.fileList = fileList
  }
}

const actions = {
  // get log list
  uploadFile({commit}, data) {
    return new Promise((resolve, reject) => {
      uploadFile(data).then(response => {
        console.log(response.data)
        const fileList = response.data
        commit('UPDATE_FILELIST', fileList)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  getFileListByKBId({commit}, data) {
    console.log(data)
    return new Promise((resolve, reject) => {
      getFileListByKBId(data).then(response => {
        console.log(response.data)
        const fileList = response.data
        commit('UPDATE_FILELIST', fileList)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
