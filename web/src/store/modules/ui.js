const state = {
  progress: false
};

const mutations = {
  setProgress: (state, progress) => state.progress = progress
};

export default {
  namespaced: true,
  state,
  mutations
}