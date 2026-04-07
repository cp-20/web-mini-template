<template>
  <BasePanel>
    <h1>Members</h1>
    <p class="subtitle">バックエンドのサンプルメンバー一覧です。</p>

    <p v-if="errorMessage" class="notice">{{ errorMessage }}</p>

    <ul class="list">
      <li v-for="member in members" :key="member.id" class="card">
        <p class="card-title">{{ member.name }}</p>
        <p class="meta">id: {{ member.id }}</p>
      </li>
    </ul>
  </BasePanel>
</template>

<script setup lang="ts">
import { onMounted, ref } from "vue";

import BasePanel from "../components/BasePanel.vue";
import { getMembers, type Member } from "../lib/api";

const members = ref<Member[]>([]);
const errorMessage = ref("");

onMounted(async () => {
  try {
    members.value = await getMembers();
  } catch (err) {
    errorMessage.value = (err as Error).message;
  }
});
</script>
