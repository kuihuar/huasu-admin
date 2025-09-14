
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公告标题（必填）:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入公告标题（必填）" />
</el-form-item>
        <el-form-item label="公告摘要（可选）:" prop="summary">
    <el-input v-model="formData.summary" :clearable="true" placeholder="请输入公告摘要（可选）" />
</el-form-item>
        <el-form-item label="公告正文（必填）:" prop="content">
    <RichEdit v-model="formData.content"/>
</el-form-item>
        <el-form-item label="封面图 URL（可选）:" prop="cover">
    <SelectImage
     v-model="formData.cover"
     file-type="image"
    />
</el-form-item>
        <el-form-item label="公告来源（可选）:" prop="source">
    <el-input v-model="formData.source" :clearable="true" placeholder="请输入公告来源（可选）" />
</el-form-item>
        <el-form-item label="发布状态（必填），控制公告是否对外展示，常见值： - 0：草稿（仅创建者可见，未提交审核） - 1：待审核（提交给管理员，未发布） - 2：已发布（全员 / 指定用户可见） - 3：已下架（历史公告，不再展示）:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入发布状态（必填），控制公告是否对外展示，常见值： - 0：草稿（仅创建者可见，未提交审核） - 1：待审核（提交给管理员，未发布） - 2：已发布（全员 / 指定用户可见） - 3：已下架（历史公告，不再展示）" />
</el-form-item>
        <el-form-item label="公告类型（必填，用于分类筛选），如： - 按业务分：系统公告、活动公告、通知公告、预警公告 - 按部门分：全员公告、技术部公告、财务部公告:" prop="type">
    <el-input v-model.number="formData.type" :clearable="true" placeholder="请输入公告类型（必填，用于分类筛选），如： - 按业务分：系统公告、活动公告、通知公告、预警公告 - 按部门分：全员公告、技术部公告、财务部公告" />
</el-form-item>
        <el-form-item label="是否置顶（可选），1 = 置顶，0 = 不置顶，置顶公告在列表页优先排序（如紧急系统通知）。:" prop="top">
    <el-input v-model.number="formData.top" :clearable="true" placeholder="请输入是否置顶（可选），1 = 置顶，0 = 不置顶，置顶公告在列表页优先排序（如紧急系统通知）。" />
</el-form-item>
        <el-form-item label="生效时间（可选，定时发布）:" prop="start_time">
    <el-date-picker v-model="formData.start_time" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="失效时间（可选，自动下架）:" prop="end_time">
    <el-date-picker v-model="formData.end_time" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
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
  createNotice,
  updateNotice,
  findNotice
} from '@/api/huasu/notice'

defineOptions({
    name: 'NoticeForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            title: '',
            summary: '',
            content: '',
            cover: "",
            source: '',
            status: undefined,
            type: undefined,
            top: undefined,
            start_time: new Date(),
            end_time: new Date(),
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findNotice({ ID: route.query.id })
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
               res = await createNotice(formData.value)
               break
             case 'update':
               res = await updateNotice(formData.value)
               break
             default:
               res = await createNotice(formData.value)
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
