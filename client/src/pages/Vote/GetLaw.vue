<template>
  <div class="law get-law-container">
    <h1>More informations about this law</h1>

    <div v-if="loading" class="loading">Loading...</div>

    <div v-if="error" class="error">{{ error }}</div>

      <!-- <div v-if="law" class="content">
        <h2>{{ law.uuid }}</h2>
        <p>{{ law.title }}</p>
        <p>{{ law.description }}</p>
        <p>{{ law.start_date }}</p>
        <p>{{ law.end_date }}</p>
      </div> -->
      <div class="mb-4" v-if="law">
        <h2 class="block text-gray-700 text-sm font-bold mb-2">
          Title
        </h2>
        <p class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        >{{ law.title }}</p>
      </div>
      <div class="mb-6" v-if="law">
        <h2 class="block text-gray-700 text-sm font-bold mb-2">
          Description
        </h2>
        <p class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        >{{ law.description }}</p>
      </div>
      <div class="mb-6" v-if="law">
        <h2 class="block text-gray-700 text-sm font-bold mb-2">
          Start date
        </h2>
        <p class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        >{{ law.start_date }}</p>
      </div>
      <div class="mb-6" v-if="law">
        <h2 class="block text-gray-700 text-sm font-bold mb-2">
          End date
        </h2>
        <p class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        >{{ law.end_date }}</p>
      </div>
      <div class="mb-6" v-if="law">
        <h2 class="block text-gray-700 text-sm font-bold mb-2">
          Number of votes
        </h2>
        <p class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        >{{ law.uuid_votes ? law.uuid_votes.length : '0' }}</p>
      </div>
      <div v-if="law" class="flex items-center justify-between">
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline"
        v-on:click="vote(law.uuid)">
          Vote for this law
        </button>
      </div>
    </div>
    
</template>

<script>
import { getLaw, updateLaw } from "@/services/vote.service";

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
    },
    vote : function (uuid){
      this.error = null;
      updateLaw(uuid)
        .then(async response => {
          if (response.status === 200) {
            console.log('voted');
            this.$router.push(`/laws`);
          } else if(response.status === 401) {
            const responseError = await response.json();
            this.error = responseError.error;
          }
        })
        .catch(response => {
          console.error('error vote for the law :', response);
        })
    }
  }
};
</script>

<style>

.law-title {
    margin: 40px;
  }

  .get-law-container {
    margin: 0 50px 0 50px;
  }
</style>