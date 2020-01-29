<template>
    <div class="Home">
        <div class="container mx-auto h-screen bg-white text-teil-800 ">
            <div class="bg-grey-100 min-h-screen">
                <div class=" mx-auto">
                    <div class="relative">
                        <div class="relative lg:flex overflow-hidden">
                            <div class="h-56 lg:h-auto lg:w-5/12 relative flex items-center justify-center">
                                <img class="absolute h-full w-full object-cover"
                                    src="https://stripe.com/img/v3/payments/overview/photos/slack.jpg" alt="" />
                                <div class="absolute inset-0 bg-teal-500 opacity-75"></div>
                                <h1 class="relative text-white text-5xl">VueJS API GO</h1>
                            </div>
                            <div class="relative lg:w-7/12 bg-gray-200">
                                <svg class="absolute h-full text-teal-200 w-24 -ml-12" fill="currentColor" viewBox="0 0 100 100"
                                    preserveAspectRatio="none">
                                    <polygon points="50,0 100,0 50,100 0,100" />
                                </svg>
                                <div class="relative py-12 lg:py-24 px-8 lg:px-16 text-gray-700 leading-relaxed">
                                    <h2 class="font-bold">Objectif du projet</h2>
                                    <ul class="list-disc">
                                        <li>Le but du projet étant de générer un front en utilisant l'API généré pour le projet de GO.</li>
                                        <li>Le front doit reprendre toutes les fonctionnalités de l'API.</li>
                                        <li>Le choix des modules VueJS est libre afin que les étudiants apprennent à sélectionner des modules open-source.</li>
                                        <li>L'exercice Formik4Vue doit être inclus dans le projet et utiliser à chaque fois qu'il est question de formulaire.</li>
                                    </ul>
                                    <p class="mt-6">
                                        <button v-if="!isLogged" class="bg-teal-600 hover:bg-teil-700 text-white font-bold py-2 px-4 rounded-full">
                                            <router-link to="/login" class="inline-block text-sm px-4 py-2 leading-none rounded text-white mt-4 lg:mt-0">
                                                Connectez-vous pour voir le projet &rarr;
                                            </router-link>
                                        </button>
                                        <button v-if="isLogged" class="bg-teal-600 hover:bg-teil-700 text-white font-bold py-2 px-4 rounded-full">
                                            <router-link to="/laws" class="inline-block text-sm px-4 py-2 leading-none rounded text-white mt-4 lg:mt-0">
                                                Allez donc voter ! &rarr;
                                            </router-link>
                                        </button>
                                    </p>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import store from "@/store/user";
export default {
  name: 'home',
  computed: {
    isLogged () {
      return store.getters.isLogged
    }
  }
};
</script>

<style scoped>

</style>