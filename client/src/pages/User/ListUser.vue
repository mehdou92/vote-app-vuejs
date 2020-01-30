<template>
  <div class="users">
    <div class="bg-teal-100 border-t-4 border-teal-500 rounded-b text-teal-900 px-4 py-3 shadow-md">
      <div class="flex">
        <svg
          class="fill-current h-6 w-6 text-teal-500 mr-4"
          xmlns="http://www.w3.org/2000/svg"
          viewBox="0 0 20 20"
        >
          <path
            d="M3 13h2v-2H3v2zm0 4h2v-2H3v2zm0-8h2V7H3v2zm4 4h14v-2H7v2zm0 4h14v-2H7v2zM7 7v2h14V7H7z"
          />
        </svg>
        <p class="font-bold">List of all user</p>
      </div>
    </div>
    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">
      <Error :error="error"></Error>
    </div>
    
    <div class="p-6 w-flex -mx-2" v-if="users" :key="componentKey">
      <table class="table" w-full>

        <thead>
          <tr>
            <th class="bg-teal-400 text-white px-4 py-2">First name</th>
            <th class="bg-teal-400 text-white px-4 py-2">Last name</th>
            <th class="bg-teal-400 text-white px-4 py-2">UUID</th>
            <th class="bg-teal-400 text-white px-4 py-2">Status</th>
            <th class="bg-teal-400 text-white px-4 py-2">Action</th>
          </tr>
        </thead>
        <tbody v-for="user in users" :key="user.uuid" class="content">
          <tr>
            <td class="bg-grey-600 border px-4 py-2">{{ user.first_name }}</td>
            <td class="bg-grey-600 border px-4 py-2">{{ user.last_name }}</td>
            <td class="bg-grey-600 border px-4 py-2">{{ (user.access_level === 1) ? "Admin" : "User" }}</td>
            <td class="bg-grey-600 border px-4 py-2">{{ user.uuid }}</td>
            <td class="bg-grey-600 border px-4 py-2">
              <span
                v-on:click="deleteU(user.uuid); forceRerender();"
                class="vote-btn bg-white inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
              >Delete User</span>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
    <div class="flex items-center p-4 justify-between">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
            <router-link to="/create-user" class="inline-block text-sm px-4 py-2 leading-none rounded text-white mt-4 lg:mt-0">
                Create new user
            </router-link>
          </button>
        </div>
  </div>
</template>

<script>
import { getUsers, deleteUser } from "@/services/user.service";
import Error from "@/pages/error/Error.vue";

export default {
  name: "ListUser",
  data() {
    return {
      loading: false,
      users: null,
      error: null,
      componentKey: 0
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    // call again the method if the route changes
    $route: "fetchData"
  },
  components: {
    Error
  },
  methods: {
    forceRerender() {
      this.componentKey += 1;
    },

    deleteU: function(uuid) {
      deleteUser(uuid).then(async response => {
        if (response.status === 200) {
          console.log("deleted user");
        } else {
          console.error("error delete user");
        }
      });
    },

    fetchData() {
      this.error = this.users = null;
      this.loading = true;
      // replace `getUsers` with your data fetching util / API wrapper

      getUsers()
        .then(async response => {
          this.loading = false;
          if (response.status === 200) {
            const userData = await response.json();
            this.users = userData;
          }
        })
        .catch(response => {
          this.error = response.toString();
        });
    }
  }
};
</script>

<style>
</style>