import Vue from 'vue';
import Vuex from 'vuex';
import VuexPersist from 'vuex-persist';

Vue.use(Vuex);

const vuexPersist = new VuexPersist({
  key: 'vote-app',
  storage: window.localStorage
});

export default new Vuex.Store({
  plugins: [vuexPersist.plugin],
  state: {
    token: null,
    accessLevel: 0
  },
  mutations: {
    authenticate(state, payload) {
      state.token = payload.JWT;
      state.accessLevel = payload.AccessLevel;
    }
  }
});