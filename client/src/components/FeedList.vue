<template>
  <ul class="list">
    <li v-for="item in items" :key="item.task_id" class="card">
      <p class="card-title">{{ item.task_title }}</p>
      <p class="meta">
        Project: {{ item.project_name }} / Owner: {{ item.owner_name }} /
        <span class="badge">{{ item.task_status }}</span>
      </p>
    </li>
  </ul>
</template>

<script setup lang="ts">
import type { FeedItem } from "../lib/api";

defineProps<{
  items: FeedItem[];
}>();
</script>
