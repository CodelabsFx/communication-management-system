<template>
  <div class="login-page">
    <h1>Login</h1>
    <form @submit.prevent="submit">
      <label>
        Email
        <input v-model="email" type="email" required />
      </label>
      <label>
        Password
        <input v-model="password" type="password" required />
      </label>
      <button type="submit">Sign in</button>
      <p v-if="error" class="error-message">{{ error }}</p>
    </form>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useAuthStore } from '../stores/auth.store';
import { login } from '../services/auth.service';

const email = ref('');
const password = ref('');
const error = ref('');
const router = useRouter();
const authStore = useAuthStore();

const submit = async () => {
  error.value = '';

  try {
    const response = await login({ email: email.value, password: password.value });
    localStorage.setItem('auth_token', response.token);
    authStore.setSession(response.token, 'user-123', 'company-123', ['user']);
    await router.push('/dashboard');
  } catch (err) {
    error.value = 'Invalid credentials. Please try again.';
    console.error(err);
  }
};
</script>

<style scoped>
.login-page {
  padding: 2rem;
  max-width: 420px;
  margin: 0 auto;
}
label {
  display: block;
  margin-bottom: 1rem;
}
input {
  width: 100%;
  padding: 0.75rem;
  margin-top: 0.5rem;
}
button {
  padding: 0.75rem 1.25rem;
}
.error-message {
  margin-top: 1rem;
  color: #c53030;
}
</style>
