<template>
  <main class="shell">
    <header class="topbar">
      <div>
        <p class="brand">Spring Hackathon Template</p>
        <p class="subtitle">Vue + Router + Echo API</p>
      </div>
      <nav class="nav">
        <RouterLink to="/">Dashboard</RouterLink>
        <RouterLink to="/members">Members</RouterLink>
      </nav>
    </header>

    <RouterView />
  </main>
</template>

<script setup lang="ts">
import { RouterLink, RouterView } from "vue-router";
</script>
