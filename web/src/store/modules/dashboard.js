import $http from '@/http'

const state = {
  info: null,
};

const mutations = {
  setInfo: (state, info) => state.info = info
};

const actions = {
  fetchInfo: ({commit}) => {
    $http.get('/').then(({data}) => {
      commit('setInfo', data);
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
}