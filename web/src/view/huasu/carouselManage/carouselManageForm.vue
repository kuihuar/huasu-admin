
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="轮播图标题:" prop="title">
    <el-input v-model="formData.title" :clearable="false" placeholder="请输入轮播图标题" />
</el-form-item>
        <el-form-item label="轮播图片的存储路径或 CDN 链接:" prop="image">
    <SelectImage
     v-model="formData.image"
     file-type="image"
    />
</el-form-item>
        <el-form-item label="点击轮播图后的跳转链接:" prop="link">
    <el-input v-model="formData.link" :clearable="true" placeholder="请输入点击轮播图后的跳转链接" />
</el-form-item>
        <el-form-item label="排序权重:" prop="sort">
    <el-input v-model.number="formData.sort" :clearable="true" placeholder="请输入排序权重" />
</el-form-item>
        <el-form-item label="状态（如：1 - 启用、0 - 禁用、2 - 草稿，控制是否在前端显示）:" prop="status">
    <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态（如：1 - 启用、0 - 禁用、2 - 草稿，控制是否在前端显示）" />
</el-form-item>
        <el-form-item label="关联业务模块:" prop="business">
    <el-input v-model="formData.business" :clearable="true" placeholder="请输入关联业务模块" />
</el-form-item>
        <el-form-item label="轮播图描述（管理后台备注，如 “618 活动主视觉”）:" prop="description">
    <el-input v-model="formData.description" :clearable="true" placeholder="请输入轮播图描述（管理后台备注，如 “618 活动主视觉”）" />
</el-form-item>
        <el-form-item label="图片的 alt 属性:" prop="alt">
    <el-input v-model="formData.alt" :clearable="true" placeholder="请输入图片的 alt 属性" />
</el-form-item>
        <el-form-item label="链接打开方式:" prop="target">
    <el-input v-model="formData.target" :clearable="true" placeholder="请输入链接打开方式" />
</el-form-item>
        <el-form-item label="创建人 ID:" prop="created_by">
    <el-input v-model.number="formData.created_by" :clearable="true" placeholder="请输入创建人 ID" />
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
  createCarouselManage,
  updateCarouselManage,
  findCarouselManage
} from '@/api/huasu/carouselManage'

defineOptions({
    name: 'CarouselManageForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const formData = ref({
            title: '',
            image: "",
            link: '',
            sort: undefined,
            status: undefined,
            business: '',
            description: '',
            alt: '',
            target: '',
            created_by: undefined,
        })
// 验证规则
const rule = reactive({
               title : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               image : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               link : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               sort : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               business : [{
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
      const res = await findCarouselManage({ ID: route.query.id })
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
               res = await createCarouselManage(formData.value)
               break
             case 'update':
               res = await updateCarouselManage(formData.value)
               break
             default:
               res = await createCarouselManage(formData.value)
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
