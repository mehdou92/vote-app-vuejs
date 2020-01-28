<template>
  <div class="add-law-container">
    <h1 class="law-title">Add New law</h1>
    <Formik
      :initial-values="{
            title: '',
            description: '',
            start_date: '',
            end_date: ''
        }"
      @onSubmit="handleSubmit"
    >
      <form class="centered">

      <div class="mb-4">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="title">
          Title
        </label>
        <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
        type="text" name="title" placeholder="Title of the law" tag="input" />
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="description">
          Description
        </label>
        <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
        type="text"
              name="description"
              placeholder="Description of the law"
              tag="input" />
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="start_date">
          Start date
        </label>
        <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
        type="text"
              name="start_date"
              placeholder="2004-01-01T00:00:00.000Z"
              tag="input" />
      </div>
      <div class="mb-6">
        <label class="block text-gray-700 text-sm font-bold mb-2" for="end_date">
          End date
        </label>
        <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
        type="text"
              name="end_date"
              placeholder="2004-01-01T00:00:00.000Z"
              tag="input" />
      </div>
      <div class="flex items-center justify-between">
        <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
          Create new law
        </button>
      </div>

      </form>
    </Formik>
  </div>
</template>

<script>
import Formik from "@/components/formik/Formik.vue";
import Field from "@/components/formik/Field.vue";
import { createLaw } from "@/services/vote.service";

export default {
  name: "AddNewLaw",
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();

      createLaw({
        title: values.title,
        description: values.description,
        start_date: values.start_date,
        end_date: values.end_date
      }).then(async response => {
        if (response.status === 200) {
          console.log("law created");
          this.$router.push('/laws');
        }
      }).catch(async response => {
          console.error('error create law', await response.json());
      });
    }
  }
};
</script>

<style>

  .law-title {
    margin: 40px;
  }

  .add-law-container {
    margin: 0 50px 0 50px;
  }

</style>