import $http from '@/http'

const state = {
  items: []
};

const mutations = {
  setItems: (state, items) => state.items = items
};

const actions = {
  fetchItems: ({ commit }) => {
    $http.get('/container').then(({ data }) => {
      commit('setItems', data);
    });
  }
};

export default {
  namespaced: true,
  state,
  mutations,
  actions
}