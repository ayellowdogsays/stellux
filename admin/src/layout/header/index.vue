<template>
  <div
    class="flex h-full w-full items-center justify-between px-4 dark:bg-dark"
  >
    <div class="w-full flex justify-between">
      <div class="flex items-center gap-4">
        <MenuUnfoldOutlined
          v-if="sidebarStore.collapse"
          class="cursor-pointer rounded-md p-1"
          @click="sidebarStore.setCollapse(false)"
        />
        <MenuFoldOutlined
          v-else
          class="cursor-pointer rounded-md p-1"
          @click="sidebarStore.setCollapse(true)"
        />
        <Breadcrumb />
      </div>
    </div>
    <div class="flex items-center gap-2 md:gap-4">
      <Search />
      <FullScreen />
      <Dropdown placement="bottomRight">
        <Avatar
          class="cursor-pointer"
          :src="userStore.userInfo?.avatar"
          :alt="userStore.userInfo?.username"
        >
          {{ userStore.userInfo?.nickname }}
        </Avatar>
        <template #overlay>
          <Menu>
            <Menu.Item>
              <div @click.prevent="Logout">
                <PoweroffOutlined />
                退出登录
              </div>
            </Menu.Item>
          </Menu>
        </template>
      </Dropdown>
      <Setting />
    </div>
  </div>
</template>

<script lang="ts" setup>
import { Breadcrumb, FullScreen, Search, Setting } from "./components";
import { useUserStore } from "@/store/modules/user";
import { Menu, Dropdown, Avatar, message } from "ant-design-vue";
import {
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  PoweroffOutlined,
} from "@ant-design/icons-vue";
import { clearStore } from "@/utils/clear";
import { useRouter } from "vue-router";
import { useSidebarStore } from "@/store/modules/sidebar";
const sidebarStore = useSidebarStore();
const userStore = useUserStore();
const router = useRouter();

function Logout() {
  clearStore();
  message.info("退出成功,请重新登录");
  router.push("/login");
}
</script>
