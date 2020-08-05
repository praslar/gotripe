import { userService } from '../_services';

const state = {
    all: {
    },
};

const actions = {
    getAll({ commit }) {
        commit('getAllRequest');
        userService.getAll()
            .then(
                users => commit('getAllSuccess', users),
                error => commit('getAllFailure', error),
            );
    },

    getPlans({ commit }) {
        userService.getPlans()
            .then(
                plan => commit('getPlanSuccess', plan),
                error => commit('getAllFailure', error),
            );
    },

    getPubKeys({ commit }) {
        userService.getPubKey()
            .then(
                key => commit('getKeySuccess', key),
                error => commit('getAllFailure', error),
            );
    },
};

const mutations = {
    getAllRequest(state) {
        state.all = { loading: true };
    },
    getAllSuccess(state, users) {
        state.all = { items: users };
    },
    getAllFailure(state, error) {
        state.all = { error };
    },
    getKeySuccess(state, key) {
        state.all = { key: key };
    },
    getPlanSuccess(state, plan) {
        state.all = { plan: plan}
    }
};

export const users = {
    namespaced: true,
    state,
    actions,
    mutations
};
