<template>
  <div class="law">
    <h1>Get law</h1>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

      <div v-if="law" class="content">
        <h2>{{ law.uuid }}</h2>
        <p>{{ law.title }}</p>
        <p>{{ law.description }}</p>
        <p>{{ law.start_date }}</p>
        <p>{{ law.end_date }}</p>
      </div>
    </div>
    
</template>

<script>
import { getLaw } from "@/services/vote.service";

export default {

  name: "GetLaw",
  data() {
    return {
      loading: false,
      law: null,
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
      this.error = this.law = null;
      this.loading = true;
      // replace `getPost` with your data fetching util / API wrapper
      getLaw(this.$route.params.id)
        .then(async response => {
          this.loading = false;
          if (response.status === 200) {
            const lawData = await response.json();
            this.law = lawData;
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