<template>
  <a-layout class="h-screen flex flex-col overflow-hidden">
    <a-layout-sider
      width="216"
      :trigger="null"
      collapsible
      v-model:collapsed="sidebarStore.collapse"
    >
      <!-- 侧边栏内容 -->
      <SideBar :collapsed="sidebarStore.collapse" />
    </a-layout-sider>

    <a-layout class="flex-1 mb-0">
      <a-layout-header class="!h-16 !px-0">
        <Header />
      </a-layout-header>
      <a-layout class="!py-0 mt-2 mx-2 overflow-y-scroll">
        <a-layout-content>
          <Main class="h-full"></Main>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </a-layout>
</template>

<script setup lang="ts">
import SideBar from "./sidebar/index.vue";
import Header from "./header/index.vue";
import Main from "./main/index.vue";
import { useSidebarStore, useUserStore } from "@/store";
import { getUserInfoAPI } from "@/api/user";
import { useMobile } from "@/hooks/useMobile";

// 判断是否为移动端设备
const isMobile = ref(useMobile());

const userStore = useUserStore();
const { userInfo } = storeToRefs(userStore);

const sidebarStore = useSidebarStore();
// 加载界面时初始化侧边栏状态
sidebarStore.setCollapse(isMobile.value);

// 加载界面时初始化用户信息
const getUserInfo = async () => {
  const res = await getUserInfoAPI();
  userInfo.value = res.data;
};

onMounted(async () => {
  await getUserInfo();
});
</script>

<style lang="scss"></style>
