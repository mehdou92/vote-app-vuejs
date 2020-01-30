<template>
  <div>
    <div class="bg-teal-100 border-t-4 border-teal-500 rounded-b text-teal-900 px-4 py-3 shadow-md">
        <div class="flex">
            <svg xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"><path d="M4 6H2v14c0 1.1.9 2 2 2h14v-2H4V6zm16-4H8c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-1 9h-4v4h-2v-4H9V9h4V5h2v4h4v2z"/></svg>
            <p class="font-bold">Create user</p>
        </div>
    </div>
    <Formik
      :initial-values="{
            first_name: '',
            last_name: '',
            email: '',
            password: '',
            date_of_birth: '',
            access_level: false,
        }"
      @onSubmit="handleSubmit"
    >
      <form class="centered">
        <div class="mb-4">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="title">
            First name
          </label>
          <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" type="text" name="first_name" placeholder="firstname" tag="input" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="description">
            Last name
          </label>
          <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" type="text" name="last_name" placeholder="lastname" tag="input" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="email">
            Email
          </label>
          <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
          type="email" name="email" placeholder="Email" tag="input" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
            Password
          </label>
          <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" type="password" name="password" placeholder="password" tag="input" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="date_of_birth">
            Date of birth <!-- 2000-01-01T00:00:00.000Z -->
          </label>
          <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" type="text" name="date_of_birth" placeholder="2000-01-01T00:00:00.000Z" tag="input" />
        </div>
        <div class="mb-6">
          <label class="block text-gray-700 text-sm font-bold mb-2" for="is_admin"> Is Admin </label>
          <Field class="input stack" type="checkbox" name="access_level" tag="input" /> 
        </div>
        <div class="flex items-center justify-between">
          <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
            Create user
          </button>
        </div>

        </form>
    </Formik>
  </div>
</template>

<script>
import Formik from "@/components/formik/Formik.vue";
import Field from "@/components/formik/Field.vue";
import { createUser } from "@/services/user.service";

export default {
  name: "CreateUser",
  components: {
    Formik,
    Field
  },
  methods: {
    handleSubmit({ event, values }) {
      event.preventDefault();

      const accessLevel = values.access_level ? 1 : 2;

      createUser({
          first_name: values.first_name,
          last_name: values.last_name,
          email: values.email,
          password: values.password,
          date_of_birth: values.date_of_birth,
          access_level: accessLevel
      })
        .then(async response => {
            if (response.status === 200) {
                console.log('created');
                this.$router.push('/login');
            }
        })
        .catch(response => {
            console.error('error create user ', response);
        });

    }
  }
};
</script>

<style>
</style>