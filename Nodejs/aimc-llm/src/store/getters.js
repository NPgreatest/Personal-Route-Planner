const getters = {
  create: state => state.createApp.create,
  submit: state => state.submitApp.submit,
  edit: state => state.createApp.edit,
  delete: state => state.deleteApp.deleteapp,
  allapp: state => state.getallApp.allapp,
  app: state => state.getApp.app,
  openapp: state => state.openApp.openapp,
  stopapp: state => state.stopApp.stopapp,
  testapp: state => state.testApp.testapp,
  prompt: state => state.submitPrompt.prompt,
  stoptestapp: state => state.stoptestApp.stoptestapp,
  knowledge_list: state => state.get_knowledgebase.knowledge,
  test_hello: state => state.test_services.hello,
  knowledgeList: state => state.knowledge.knowledgeList,
  fileList: state => state.file.fileList,
}

export default getters
