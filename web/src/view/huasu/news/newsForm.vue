
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="新闻标题:" prop="title">
    <el-input v-model="formData.title" :clearable="true" placeholder="请输入新闻标题" />
</el-form-item>
        <el-form-item label="新闻分类:" prop="category">
    <el-input v-model="formData.category" :clearable="true" placeholder="请输入新闻分类" />
</el-form-item>
        <el-form-item label="新闻摘要:" prop="summary">
    <el-input v-model="formData.summary" :clearable="true" placeholder="请输入新闻摘要" />
</el-form-item>
        <el-form-item label="新闻正文:" prop="content">
    <RichEdit v-model="formData.content"/>
</el-form-item>
        <el-form-item label="封面图:" prop="cover">
    <SelectImage
     v-model="formData.cover"
     file-type="image"
    />
</el-form-item>
        <el-form-item label="新闻来源:" prop="source">
    <el-input v-model="formData.source" :clearable="true" placeholder="请输入新闻来源" />
</el-form-item>
        <el-form-item label="阅读量:" prop="views">
    <el-input v-model.number="formData.views" :clearable="true" placeholder="请输入阅读量" />
</el-form-item>
        <el-form-item label="publishtime:" prop="publishtime">
    <el-date-picker v-model="formData.publishtime" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
</el-form-item>
        <el-form-item label="新闻状态:" prop="status">
    <el-input v-model="formData.status" :clearable="true" placeholder="请输入新闻状态" />
</el-form-item>
        <el-form-item label="作者:" prop="author">
    <el-input v-model="formData.author" :clearable="true" placeholder="请输入作者" />
</el-form-item>
        <el-form-item label="点赞数:" prop="like">
    <el-input v-model.number="formData.like" :clearable="true" placeholder="请输入点赞数" />
</el-form-item>
        <el-form-item label="不点赞数:" prop="dislike">
    <el-input v-model.number="formData.dislike" :clearable="true" placeholder="请输入不点赞数" />
</el-form-item>
        <el-form-item label="链接:" prop="link">
    <el-input v-model="formData.link" :clearable="true" placeholder="请输入链接" />
</el-form-item>
        <el-form-item label="排名:" prop="sort">
    <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排名" />
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
  createNews,
  updateNews,
  findNews
} from '@/api/huasu/news'

defineOptions({
    name: 'NewsForm'
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
            category: '',
            summary: '',
            content: '',
            cover: "",
            source: '',
            views: undefined,
            publishtime: new Date(),
            status: '',
            author: '',
            like: undefined,
            dislike: undefined,
            link: '',
            sort: undefined,
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               summary : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               publishtime : [{
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
      const res = await findNews({ ID: route.query.id })
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
               res = await createNews(formData.value)
               break
             case 'update':
               res = await updateNews(formData.value)
               break
             default:
               res = await createNews(formData.value)
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
