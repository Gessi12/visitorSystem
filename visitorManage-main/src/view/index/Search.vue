<template>
  <div>
  <center>
  <el-form ref="searchForm" :model="searchForm" class="demo-form-inline">
    <el-form-item label="姓名" prop="visitorName">
      <el-input v-model="searchForm.visitorName" placeholder="姓名" ></el-input>
    </el-form-item>
    <el-form-item label="联系方式" prop="phone">
      <el-input v-model="searchForm.phone" placeholder="联系方式"></el-input>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" v-on:click="searchUsers( 'searchForm' )">查询</el-button>
    </el-form-item>
  </el-form>
      <el-table :data="visitors" border>
        <el-table-column prop="visitor_name" label="姓名"></el-table-column>
        <el-table-column prop="created_at" label="到访日期"></el-table-column>
        <el-table-column prop="event" label="到访事由"></el-table-column>
      </el-table>
  </center>
  </div>
</template>
<script>
import axios from 'axios';
export default {
  name: "Search",
  data() {
    return {
      searchForm:{
        visitorName: '',
        phone: ''
      },
      visitors:[]
    }
  },
  methods: {
    searchUsers(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          axios.post(`http://localhost:8080/main/search`,this.searchForm)
            .then(response => {
              this.visitors = response.data.results;
            })
            .catch(error => {
              console.log(error);
            });
        } else {
          // 表单验证失败，给出提示
          console.log('失败');
        }
      });

    }
  }
}
</script>
