<template>
  <div class="user">
    <h1>Get user</h1>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

      <div v-if="user" class="content">
        <h2>{{ user.uuid }}</h2>
        <p>{{ user.first_name }}</p>
        <p>{{ user.last_name }}</p>
        <p>{{ user.accessLevel }}</p>
      </div>
    </div>
    
</template>

<script>
import { getUser } from "@/services/user.service";

export default {

  name: "GetUser",
  data() {
    return {
      loading: false,
      user: null,
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
      this.error = this.user = null;
      this.loading = true;
      // replace `getPost` with your data fetching util / API wrapper
      getUser(this.$route.params.id)
        .then(async response => {
          this.loading = false;
          if (response.status === 200) {
            const userData = await response.json();
            this.user = userData;
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