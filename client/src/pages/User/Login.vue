<template>
  <div>
    <h1>Login</h1>
    <Formik
      :initial-values="{
        login: '',
        password: ''
      }"
      @onSubmit="handleSubmit"
    >
      <form class="centered">
        <div class="half">
          <Field class="input stack" type="email" name="login" placeholder="Email" tag="input" />
          <Field
            class="input stack"
            type="password"
            name="password"
            placeholder="Password"
            tag="input"
          />
          <button class="button stack" type="submit">Se connecter</button>
        </div>
      </form>
    </Formik>
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
.centered {
  display: flex;
  justify-content: center;
}
</style>