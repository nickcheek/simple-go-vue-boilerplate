<template>
  <div class="home">
    <div class="container">
      <h1>Go Vue Boilerplate</h1>
      
      <div v-if="loading" class="loading">
        Loading...
      </div>
      
      <div v-else-if="error" class="error">
        Error: {{ error }}
      </div>
      
      <div v-else-if="welcomeData" class="welcome-card">
        <h2>{{ welcomeData.message }}</h2>
        <p><strong>Version:</strong> {{ welcomeData.version }}</p>
        <p><strong>Status:</strong> {{ welcomeData.status }}</p>
      </div>
      
      <button @click="fetchWelcomeData" class="refresh-btn">
        Refresh Data
      </button>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import axios from 'axios'

interface WelcomeData {
  message: string
  version: string
  status: string
}

const welcomeData = ref<WelcomeData | null>(null)
const loading = ref(false)
const error = ref<string | null>(null)

const fetchWelcomeData = async () => {
  try {
    loading.value = true
    error.value = null
    const response = await axios.get('http://localhost:8080/api/welcome')
    welcomeData.value = response.data
  } catch (err) {
    error.value = 'Failed to fetch data from backend'
    console.error('Error:', err)
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  fetchWelcomeData()
})
</script>

<style scoped>
.home {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  padding: 2rem;
}

.container {
  text-align: center;
  max-width: 600px;
  width: 100%;
}

h1 {
  color: #333;
  margin-bottom: 2rem;
  font-size: 2.5rem;
}

.welcome-card {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
}

.welcome-card h2 {
  color: #4CAF50;
  margin-bottom: 1rem;
}

.welcome-card p {
  margin: 0.5rem 0;
  color: #666;
}

.loading {
  padding: 2rem;
  color: #666;
  font-size: 1.1rem;
}

.error {
  padding: 2rem;
  color: #f44336;
  background: #ffebee;
  border-radius: 8px;
  margin-bottom: 2rem;
}

.refresh-btn {
  background: #007bff;
  color: white;
  border: none;
  padding: 0.75rem 1.5rem;
  border-radius: 4px;
  cursor: pointer;
  font-size: 1rem;
  transition: background-color 0.3s;
}

.refresh-btn:hover {
  background: #0056b3;
}
</style>
