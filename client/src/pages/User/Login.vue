<template>
<div class="login-container container-fluid">
<div class=" flex justify-center">
  <Formik
      :initial-values="{
        login: '',
        password: ''
      }"
      @onSubmit="handleSubmit"
    >
    <h1 class="text-center login-title">Login to your account</h1>
  <form class="bg-white text-center  shadow-md rounded px-8 pt-6 pb-8 mb-4">
    <div class="mb-4">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="login">
        Login
      </label>
      <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 leading-tight focus:outline-none focus:shadow-outline" 
      type="email" name="login" placeholder="Email" tag="input" />
    </div>
    <div class="mb-6">
      <label class="block text-gray-700 text-sm font-bold mb-2" for="password">
        Password
      </label>
      <Field class="shadow appearance-none border rounded w-full py-2 px-3 text-gray-700 mb-3 leading-tight focus:outline-none focus:shadow-outline" 
      type="password"
            name="password"
            placeholder="Password"
            tag="input" />
    </div>
    <div class="flex items-center justify-between">
      <button class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded focus:outline-none focus:shadow-outline" type="submit">
        Sign in
      </button>
    </div>
  </form>
  </Formik>
</div>
</div> 
</template>

<script>
import Formik from "@/components/formik/Formik.vue";
import Field from "@/components/formik/Field.vue";
import store from "@/store/user";
import { login } from "@/services/user.service";
export default {
  name: "Login",
  components: {
    Formik,
    Field
  },
  methods: {

    handleSubmit({ event, values }) {
      event.preventDefault();

      login({
        login: values.login,
        password: values.password
      })
        .then(async response => {
          if (response.status === 200) {            
            const data = await response.json();
            store.commit('authenticate', data);
            this.$router.push('/');
            console.log('connected');
          }
        })
        .catch(response => {
          console.log(response);
        });
    }
  }
};
</script>

<style lang="css">

.login-container {
  background-color: #A7ABDD;
  width: 100%;
  height: 100%;
}

.login-title {
  font-size: 20px;
  margin: 40px;
}

</style>