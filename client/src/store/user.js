import Vue from 'vue';
import Vuex from 'vuex';
import VuexPersist from 'vuex-persist';
import jwt_decode from 'jwt-decode';

Vue.use(Vuex);

const vuexPersist = new VuexPersist({
  key: 'vote-app',
  storage: window.localStorage
});

export default new Vuex.Store({
  plugins: [vuexPersist.plugin],
  state: {
    token: null
  },
  getters: {
    getToken: state => {
      return state.token
    },
    isLogged: state => {
      return state.token ? true : false
    },
    adminPermission: state => {
      return state.accessAdmin = jwt_decode(state.token); 
    }
  },
  mutations: {
    authenticate(state, payload) {
      state.token = payload.jwt;
    },
    logout(state) {
      state.token = null;
    }
  }
});