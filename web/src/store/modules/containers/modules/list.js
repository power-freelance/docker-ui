import { getField, updateField } from 'vuex-map-fields'

import $http from '@/http'

import statusItems from '@/assets/data/container-status'

const state = {
  filter: {
    all: true,
    status: statusItems.map(item => item.id)
  },
  items: []
};

const getters = {
  getField
};

const mutations = {
  updateField,
  setItems: (state, items) => state.items = items
};

const actions = {
  fetchItems: ({ state, commit }) => {
    $http.get('/container', {params: state.filter}).then(({ data }) => commit('setItems', data));
  }
};

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}