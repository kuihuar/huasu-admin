
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建日期" prop="createdAtRange">
      <template #label>
        <span>
          创建日期
          <el-tooltip content="搜索范围是开始日期（包含）至结束日期（不包含）">
            <el-icon><QuestionFilled /></el-icon>
          </el-tooltip>
        </span>
      </template>

      <el-date-picker
            v-model="searchInfo.createdAtRange"
            class="!w-380px"
            type="datetimerange"
            range-separator="至"
            start-placeholder="开始时间"
            end-placeholder="结束时间"
          />
       </el-form-item>
      

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column sortable align="left" label="日期" prop="CreatedAt" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        
            <el-table-column align="left" label="公告标题（必填）" prop="title" width="120" />

            <el-table-column align="left" label="公告摘要（可选）" prop="summary" width="120" />

            <el-table-column label="公告正文（必填）" prop="content" width="200">
   <template #default="scope">
      [富文本内容]
   </template>
</el-table-column>
            <el-table-column label="封面图 URL（可选）" prop="cover" width="200">
    <template #default="scope">
      <el-image preview-teleported style="width: 100px; height: 100px" :src="getUrl(scope.row.cover)" fit="cover"/>
    </template>
</el-table-column>
            <el-table-column align="left" label="公告来源（可选）" prop="source" width="120" />

            <el-table-column align="left" label="发布状态（必填），控制公告是否对外展示，常见值： - 0：草稿（仅创建者可见，未提交审核） - 1：待审核（提交给管理员，未发布） - 2：已发布（全员 / 指定用户可见） - 3：已下架（历史公告，不再展示）" prop="status" width="120" />

            <el-table-column align="left" label="公告类型（必填，用于分类筛选），如： - 按业务分：系统公告、活动公告、通知公告、预警公告 - 按部门分：全员公告、技术部公告、财务部公告" prop="type" width="120" />

            <el-table-column align="left" label="是否置顶（可选），1 = 置顶，0 = 不置顶，置顶公告在列表页优先排序（如紧急系统通知）。" prop="top" width="120" />

            <el-table-column align="left" label="生效时间（可选，定时发布）" prop="start_time" width="180">
   <template #default="scope">{{ formatDate(scope.row.start_time) }}</template>
</el-table-column>
            <el-table-column align="left" label="失效时间（可选，自动下架）" prop="end_time" width="180">
   <template #default="scope">{{ formatDate(scope.row.end_time) }}</template>
</el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateNoticeFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="公告标题（必填）">
    {{ detailForm.title }}
</el-descriptions-item>
                    <el-descriptions-item label="公告摘要（可选）">
    {{ detailForm.summary }}
</el-descriptions-item>
                    <el-descriptions-item label="公告正文（必填）">
    <RichView v-model="detailForm.content" />
</el-descriptions-item>
                    <el-descriptions-item label="封面图 URL（可选）">
    <el-image style="width: 50px; height: 50px" :preview-src-list="returnArrImg(detailForm.cover)" :src="getUrl(detailForm.cover)" fit="cover" />
</el-descriptions-item>
                    <el-descriptions-item label="公告来源（可选）">
    {{ detailForm.source }}
</el-descriptions-item>
                    <el-descriptions-item label="发布状态（必填），控制公告是否对外展示，常见值： - 0：草稿（仅创建者可见，未提交审核） - 1：待审核（提交给管理员，未发布） - 2：已发布（全员 / 指定用户可见） - 3：已下架（历史公告，不再展示）">
    {{ detailForm.status }}
</el-descriptions-item>
                    <el-descriptions-item label="公告类型（必填，用于分类筛选），如： - 按业务分：系统公告、活动公告、通知公告、预警公告 - 按部门分：全员公告、技术部公告、财务部公告">
    {{ detailForm.type }}
</el-descriptions-item>
                    <el-descriptions-item label="是否置顶（可选），1 = 置顶，0 = 不置顶，置顶公告在列表页优先排序（如紧急系统通知）。">
    {{ detailForm.top }}
</el-descriptions-item>
                    <el-descriptions-item label="生效时间（可选，定时发布）">
    {{ detailForm.start_time }}
</el-descriptions-item>
                    <el-descriptions-item label="失效时间（可选，自动下架）">
    {{ detailForm.end_time }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createNotice,
  deleteNotice,
  deleteNoticeByIds,
  updateNotice,
  findNotice,
  getNoticeList
} from '@/api/huasu/notice'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 富文本组件
import RichEdit from '@/components/richtext/rich-edit.vue'
import RichView from '@/components/richtext/rich-view.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'Notice'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               content : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
               status : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getNoticeList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteNoticeFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const IDs = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          IDs.push(item.ID)
        })
      const res = await deleteNoticeByIds({ IDs })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === IDs.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateNoticeFunc = async(row) => {
    const res = await findNotice({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteNoticeFunc = async (row) => {
    const res = await deleteNotice({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
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
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

const detailForm = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findNotice({ ID: row.ID })
  if (res.code === 0) {
    detailForm.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailForm.value = {}
}


</script>

<style>

</style>
