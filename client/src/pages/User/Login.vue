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

      console.log('values login : ',values);

      login({
        login: values.login,
        password: values.password
      })
        .then(response => {
            console.log('response = ', response);
          if (response.status === 200) {
            console.log("status ok  login");
          }
        })
        .catch(response => {
          console.log("error logged");
          console.log('RESPONSE ERREUR : ', response);
        });

      //   axios
      //     .post('/login', {
      //       email: values.email,
      //       password: values.password
      //     })
      //     .then(response => {
      //       if (response.status === 200) {
      //         store.commit('authenticate', response.data);
      //         this.$router.push('/votes');
      //         console.log('Logged !');
      //       }
      //     })
      //     .catch(response => {
      //       console.log(response);
      //     });
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