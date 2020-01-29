<template>
  <div class="laws">

    
    <div class="bg-teal-100 border-t-4 border-teal-500 rounded-b text-teal-900 px-4 py-3 shadow-md">
      <div class="flex">
          <svg class="fill-current h-6 w-6 text-teal-500 mr-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 20 20"><path d="M3 13h2v-2H3v2zm0 4h2v-2H3v2zm0-8h2V7H3v2zm4 4h14v-2H7v2zm0 4h14v-2H7v2zM7 7v2h14V7H7z"/></svg>
          <p class="font-bold">List of all laws</p>
      </div>
    </div>
      <div v-if="loading" class="loading">Loading...</div>

      <div v-if="error" class="error">{{ error }}</div>

      <div v-if="laws" class="flex-wrap flex">
        <div v-for="law in laws" :key="law.uuid" class="content flex-1 self-auto">
          <div class="rounded overflow-hidden shadow-lg bg-teal-200">
            <div class="px-6 py-4">
              <div class="font-bold text-xl mb-2"> {{ law.title }}</div>
              <p class="text-gray-700 text-base">{{ law.description }}</p>
            </div>
            <div class="px-6 py-4">
              <span class="vote-btn bg-white inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2"
              v-on:click="vote(law.uuid)">
                Vote !
              </span>
              <router-link :to="'/law/' + law.uuid" class="vote-btn bg-white inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2">
                More
              </router-link>
            </div>
          </div>
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
            this.$router.push(`/law/${uuid}`);
          } else if(response.status === 401) {
            const responseError = await response.json();
            this.error = responseError.error;
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

  .vote-btn {
    cursor: pointer;
  }

  .content {
    margin: 30px;
    width: 30%;
    min-width: 30%;
  }


</style>