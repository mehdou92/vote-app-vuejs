<template>
  <div class="laws">
      <h1>List of all laws</h1>

      <div v-if="loading" class="loading">Loading...</div>

      <div v-if="error" class="error">{{ error }}</div>

      <div v-if="laws">
        <div v-for="law in laws" :key="law.uuid" class="content">
          <h2>{{ law.uuid }}</h2>
          <p>{{ law.title }}</p>
          <p>{{ law.description }}</p>
          <p>{{ law.start_date }}</p>
          <p>{{ law.end_date }}</p>
          <button v-on:click="vote(law.uuid)">Vote !</button>
        </div>
      </div>
  </div>
</template>

<script>
import { getLaws, updateLaw } from "@/services/vote.service";

export default {
  name:"ListLaws",
  data() {
    return {
      loading: false,
      laws: null,
      error: null
    };
  },
  created(){
    this.fetchData();
  },
  watch: {
        // call again the method if the route changes
    $route: "fetchData"
  },
  methods: {
    vote : function (uuid){
      this.error = null;
      updateLaw(uuid)
        .then(async response => {
          if (response.status === 200) {
            console.log('voted');
            this.$router.push(`/laws/${uuid}`);
          } else if(response.status === 401) {
            this.error = 'The proposal was  already voted';
          }
        })
        .catch(response => {
          console.error('error vote for the law :', response);
        })
    },

    fetchData() {
      this.error = this.laws = null;
      this.loading = true;

      getLaws()
        .then(async response => {
          this.loading = false;
          if (response.status === 200) {
            const lawData = await response.json();
            this.laws = lawData;
          }
        })
        .catch(response => {
          this.error = response.toString();
        });
    }
  }
}
</script>

<style>

</style>