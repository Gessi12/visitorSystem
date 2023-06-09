<template>
  <div>
    <h1>
      <el-page-header @back="goBack" content="访客登记"></el-page-header>
    </h1>
    <el-form :model="ruleForm" :rules="rules" ref="ruleForm" label-width="100px" class="demo-ruleForm">
      <el-form-item label="姓名" prop="visitor_name">
        <el-input v-model="ruleForm.visitor_name"></el-input>
      </el-form-item>

      <el-form-item label="性别" prop="sex">
        <el-radio-group v-model="ruleForm.sex">
          <el-radio label="男"></el-radio>
          <el-radio label="女"></el-radio>
        </el-radio-group>
      </el-form-item>

      <el-form-item label="联系方式" prop="phone">
        <el-input v-model="ruleForm.phone"></el-input>
      </el-form-item>

      <el-form-item label="身份证" prop="visit_id">
        <el-input v-model="ruleForm.visit_id"></el-input>
      </el-form-item>

      <el-form-item label="事由" prop="event">
        <el-input type="textarea" v-model="ruleForm.event"></el-input>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">立即登记</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>


<script>
import router from "../router";
import axios from 'axios';

export default {
  name : "Register",
  data() {
    return {
      ruleForm: {
        visitor_name: '',
        sex: '',
        phone:'',
        visit_id:'',
        event: ''
      },
      rules: {
        visitor_name: [
          { required: true, message: '请输入性名', trigger: 'blur' },
        ],
        sex: [
          { required: true, message: '请选择性别', trigger: 'change' }
        ],
        phone:[
          { required: true, message: '请输入联系方式', trigger: 'blur' },
          //{ pattern: /^1[3|4|5|6|7|8|9][0-9]\d{8}$/, message: "请输入正确的联系方式", trigger: "blur"}
        ],
        visit_id:[
          { required: true, message: '请输入身份证', trigger: 'blur' },
          //{ pattern:/(^\d{15}$)|(^\d{18}$)|(^\d{17}(\d|X|x)$)/, message: "请输入正确的身份证", trigger: "blur"}
        ],
        event: [
          { required: true, message: '请填写事由', trigger: 'blur' }
        ]
      }
    };
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          axios.post('http://localhost:8080/register', this.ruleForm)
            .then((response) => {
              if(response.data.message=="success"){
                console.log("登记成功");
                this.$router.push('/main/index');
              }else {
               console.log("登记失败")
              }
            })

        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    resetForm(formName) {
      this.$refs[formName].resetFields();
    },
    goBack() {
      router.back();
    }
  }
}
</script>
