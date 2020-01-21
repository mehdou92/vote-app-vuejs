<template>
  <div>
    <h1>Create user</h1>
    <Formik
      :initial-values="{
            first_name: '',
            last_name: '',
            email: '',
            password: '',
            date_of_birth: '',
            access_level: 2
        }"
      @onSubmit="handleSubmit"
    >
      <form class="centered">
        <div class="half">
          <Field
            class="input stack"
            type="text"
            name="first_name"
            placeholder="firstname"
            tag="input"
          />
          <Field
            class="input stack"
            type="text"
            name="last_name"
            placeholder="lastname"
            tag="input"
          />
          <Field class="input stack" type="email" name="email" placeholder="Email" tag="input" />
          <Field
            class="input stack"
            type="password"
            name="password"
            placeholder="password"
            tag="input"
          />2000-01-01T00:00:00.000Z
          <Field
            class="input stack"
            type="text"
            name="date_of_birth"
            placeholder="date of birth"
            tag="input"
          />
          <Field class="input stack" type="checkbox" name="access_level" tag="input" /> Is Admin? 

          <button class="button stack" type="submit">Create user</button>
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