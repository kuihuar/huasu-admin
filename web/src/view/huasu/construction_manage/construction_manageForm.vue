
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="工程名:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入工程名" />
</el-form-item>
        <el-form-item label="工程描述:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入工程描述" />
</el-form-item>
        <el-form-item label="工程类型:" prop="type">
    <el-input v-model="formData.type" :clearable="true" placeholder="请输入工程类型" />
</el-form-item>
        <el-form-item label="项目位置:" prop="location">
    <el-input v-model="formData.location" :clearable="true" placeholder="请输入项目位置" />
</el-form-item>
        <el-form-item label="责任人:" prop="owner">
    <el-input v-model="formData.owner" :clearable="true" placeholder="请输入责任人" />
</el-form-item>
        <el-form-item label="设计单位:" prop="design">
    <el-input v-model="formData.design" :clearable="true" placeholder="请输入设计单位" />
</el-form-item>
        <el-form-item label="建造单位:" prop="contractor">
    <el-input v-model="formData.contractor" :clearable="true" placeholder="请输入建造单位" />
</el-form-item>
        <el-form-item label="管理单位:" prop="management">
    <el-input v-model="formData.management" :clearable="true" placeholder="请输入管理单位" />
</el-form-item>
        <el-form-item label="projectmanager字段:" prop="projectmanager">
    <el-input v-model="formData.projectmanager" :clearable="true" placeholder="请输入projectmanager字段" />
</el-form-item>
        <el-form-item label="开工日期:" prop="scheduleddate">
    <el-date-picker v-model="formData.scheduleddate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="实际开开日期:" prop="actualstartdate">
    <el-date-picker v-model="formData.actualstartdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="计划结束日期:" prop="plannedfinishdate">
    <el-date-picker v-model="formData.plannedfinishdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="contractnumber字段:" prop="contractnumber">
    <el-input v-model="formData.contractnumber" :clearable="true" placeholder="请输入contractnumber字段" />
</el-form-item>
        <el-form-item label="开工日期:" prop="initiationdate">
    <el-date-picker v-model="formData.initiationdate" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="项目展示图片:" prop="images">
    <SelectImage
     multiple
     v-model="formData.images"
     file-type="image"
     />
</el-form-item>
        <el-form-item label="项目文档:" prop="pdfs">
    <SelectFile v-model="formData.pdfs" />
</el-form-item>
        <el-form-item label="项目视频:" prop="video">
    <SelectImage
    v-model="formData.video"
    file-type="video"
    />
</el-form-item>
        <el-form-item label="vudgetedcost字段:" prop="vudgetedcost">
    <el-input v-model.number="formData.vudgetedcost" :clearable="true" placeholder="请输入vudgetedcost字段" />
</el-form-item>
        <el-form-item label="实际花费:" prop="actualcost">
    <el-input v-model.number="formData.actualcost" :clearable="true" placeholder="请输入实际花费" />
</el-form-item>
        <el-form-item label="预估花费:" prop="estimatecompletion">
    <el-input v-model.number="formData.estimatecompletion" :clearable="true" placeholder="请输入预估花费" />
</el-form-item>
        <el-form-item label="花费单位:" prop="costunit">
    <el-input v-model="formData.costunit" :clearable="true" placeholder="请输入花费单位" />
</el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createConstructionManage,
  updateConstructionManage,
  findConstructionManage
} from '@/api/huasu/construction_manage'

defineOptions({
    name: 'ConstructionManageForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            title: '',
            description: '',
            type: '',
            location: '',
            owner: '',
            design: '',
            contractor: '',
            management: '',
            projectmanager: '',
            scheduleddate: new Date(),
            actualstartdate: new Date(),
            plannedfinishdate: new Date(),
            contractnumber: '',
            initiationdate: new Date(),
            images: [],
            pdfs: [],
            video: "",
            vudgetedcost: undefined,
            actualcost: undefined,
            estimatecompletion: undefined,
            costunit: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findConstructionManage({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      btnLoading.value = true
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return btnLoading.value = false
            let res
           switch (type.value) {
             case 'create':
               res = await createConstructionManage(formData.value)
               break
             case 'update':
               res = await updateConstructionManage(formData.value)
               break
             default:
               res = await createConstructionManage(formData.value)
               break
           }
           btnLoading.value = false
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
