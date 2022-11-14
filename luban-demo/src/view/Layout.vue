<template>
  <div>
    <el-container class="header-container">
      <!--头部-->
      <el-header style="display: flex; background-color: #001529; align-items: stretch; justify-content: space-between;height: 50px;">
        <div style="display: flex; align-items: center">
          <span style="font-size: large; margin-left: 30px">多集群管理平台</span>
        </div>
      </el-header>
      <el-container>
        <el-aside style="rgb(238, 241, 246)">
          <el-col style="height: 100%">
            <el-menu
                default-active="1"
                default-openeds="1"
                class="el-menu-vertical-no-collapse"
                @open="handleOpen"
                @close="handleClose"
                background-color="#545c64"
                text-color="#fff"
                active-text-color="#ffd04b">
              <el-submenu index="1">
                <template slot="title">
                  <i class="el-icon-location"></i>
                  <span>容器管理</span>
                </template>
                <el-menu-item-group>
                  <template slot="title">分组一</template>
                  <el-menu-item index="1-1">选项1</el-menu-item>
                  <el-menu-item index="1-2">选项2</el-menu-item>
                </el-menu-item-group>
                <el-menu-item-group title="分组2">
                  <el-menu-item index="1-3">选项3</el-menu-item>
                </el-menu-item-group>
                <el-submenu index="1-4">
                  <template slot="title">选项4</template>
                  <el-menu-item index="1-4-1">选项1</el-menu-item>
                </el-submenu>
              </el-submenu>
              <el-menu-item index="2">
                <i class="el-icon-menu"></i>
                <span slot="title">导航二</span>
              </el-menu-item>
              <el-menu-item index="3" disabled>
                <i class="el-icon-document"></i>
                <span slot="title">导航三</span>
              </el-menu-item>
              <el-menu-item index="4">
                <i class="el-icon-setting"></i>
                <span slot="title">导航四</span>
              </el-menu-item>
            </el-menu>
          </el-col>
        </el-aside>
        <el-main>
          <router-view></router-view>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  name: "Layout"
}
</script>

<style scoped>
.header-container {
  height: 100%;
  width: 100%;
  font-weight: 400;
  font-style: normal;
   position: fixed;
  min-width: 1120px;
  z-index: 1002;
  top: 0;
  left: 0;
  background-color: #ebe7e7;
}

.el-header {
  background: rgb(55, 61, 61);
  display: flex;
  justify-content: space-between;
  line-height: 60px;
  color: #fff;
  font-size: 25px;
}
.el-aside {
  height: 100%;
  width: auto !important;
}
.el-menu {
  border-right: none;
  height: 100%;
}

.el-main {
  background: #fff;
}

.el-menu-class {
  height: 51px;
}

</style>