<template>
  <div>
    <h1>Add New law</h1>
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
        <div class="half">
          <Field
            class="input stack"
            type="text"
            name="title"
            placeholder="title of the law"
            tag="input"
          />
          <Field
            class="input stack"
            type="text"
            name="description"
            placeholder="description of the law"
            tag="input"
          />
          2004-01-01T00:00:00.000Z
          <Field
            class="input stack"
            type="text"
            name="start_date"
            placeholder="start date of the law"
            tag="input"
          />
          2004-02-01T00:00:00.000Z
          <Field
            class="input stack"
            type="text"
            name="end_date"
            placeholder="end date of the law"
            tag="input"
          />
          <button class="button stack" type="submit">Create new law</button>
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
        }
      }).catch(async response => {
          console.error('error create law', await response.json());
      });
    }
  }
};
</script>

<style>
</style>