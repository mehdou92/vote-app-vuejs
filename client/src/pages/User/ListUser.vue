<template>
  <div class="users">
    <h1>List of all user</h1>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

    <div v-if="users">
      <div v-for="user in users" :key="user.uuid" class="content">
        <h2>{{ user.uuid }}</h2>
        <p>{{ user.first_name }}</p>
        <p>{{ user.last_name }}</p>
        <p>{{ user.accessLevel }}</p>
      </div>
    </div>
  </div>
</template>

<script>
import { getUsers } from "@/services/user.service";

export default {
  name: "ListUser",
  data() {
    return {
      loading: false,
      users: null,
      error: null
    };
  },
  created() {
    this.fetchData();
  },
  watch: {
    // call again the method if the route changes
    $route: "fetchData"
  },
  methods: {
    fetchData() {
      this.error = this.users = null;
      this.loading = true;
      // replace `getPost` with your data fetching util / API wrapper

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