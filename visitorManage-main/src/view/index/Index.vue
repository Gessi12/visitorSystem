<template>
  <div style="height: 1000px">
    <el-container style="height: 100%; border: 1px solid #eee">
      <el-container>
        <el-header style="text-align: right; font-size: 12px">
          <el-button type="primary">
            <router-link to="/register">登记</router-link>
          </el-button>
        </el-header>

        <el-main>
          <el-table :data="visitors">
            <el-table-column prop="visitor_name" label="姓名" width="150"></el-table-column>
            <el-table-column prop="created_at" label="到访日期" width="200"></el-table-column>
            <el-table-column prop="event" label="到访事由"></el-table-column>
          </el-table>
        </el-main>
        <el-main>
          <router-view />
        </el-main>
      </el-container>
    </el-container>
  </div>

</template>


<script>
import axios from 'axios';
export default {
  name: "Index",
  data() {
    return {
      visitors:[],
    }
  },
  created() {
    axios.get('http://localhost:8080/main/index')
      .then(response => {
        this.visitors = response.data.results;
      })
      .catch(error => {
        console.log(error);
      });
  }

}
</script>

<style scoped>
.el-header {
  background-color: #B3C0D1;
  color: #333;
  line-height: 60px;
}

.el-aside {
  color: #333;
}
</style>
