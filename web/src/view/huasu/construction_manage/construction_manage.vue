
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
        
            <el-table-column align="left" label="工程名" prop="title" width="120" />

            <el-table-column align="left" label="工程描述" prop="description" width="120" />

            <el-table-column align="left" label="工程类型" prop="type" width="120" />

            <el-table-column align="left" label="项目位置" prop="location" width="120" />

            <el-table-column align="left" label="责任人" prop="owner" width="120" />

            <el-table-column align="left" label="设计单位" prop="design" width="120" />

            <el-table-column align="left" label="建造单位" prop="contractor" width="120" />

            <el-table-column align="left" label="管理单位" prop="management" width="120" />

            <el-table-column align="left" label="projectmanager字段" prop="projectmanager" width="120" />

            <el-table-column align="left" label="开工日期" prop="scheduleddate" width="180">
   <template #default="scope">{{ formatDate(scope.row.scheduleddate) }}</template>
</el-table-column>
            <el-table-column align="left" label="实际开开日期" prop="actualstartdate" width="180">
   <template #default="scope">{{ formatDate(scope.row.actualstartdate) }}</template>
</el-table-column>
            <el-table-column align="left" label="计划结束日期" prop="plannedfinishdate" width="180">
   <template #default="scope">{{ formatDate(scope.row.plannedfinishdate) }}</template>
</el-table-column>
            <el-table-column align="left" label="contractnumber字段" prop="contractnumber" width="120" />

            <el-table-column align="left" label="开工日期" prop="initiationdate" width="180">
   <template #default="scope">{{ formatDate(scope.row.initiationdate) }}</template>
</el-table-column>
            <el-table-column label="项目展示图片" prop="images" width="200">
   <template #default="scope">
      <div class="multiple-img-box">
         <el-image preview-teleported v-for="(item,index) in scope.row.images" :key="index" style="width: 80px; height: 80px" :src="getUrl(item)" fit="cover"/>
     </div>
   </template>
</el-table-column>
            <el-table-column label="项目文档" prop="pdfs" width="200">
    <template #default="scope">
         <div class="file-list">
           <el-tag v-for="file in scope.row.pdfs" :key="file.uid" @click="onDownloadFile(file.url)">{{ file.name }}</el-tag>
         </div>
    </template>
</el-table-column>
            <el-table-column label="项目视频" prop="video" width="200">
   <template #default="scope">
    <video
       style="width: 100px; height: 100px"
       muted
       preload="metadata"
       >
         <source :src="getUrl(scope.row.video) + '#t=1'">
       </video>
   </template>
</el-table-column>
            <el-table-column align="left" label="vudgetedcost字段" prop="vudgetedcost" width="120" />

            <el-table-column align="left" label="实际花费" prop="actualcost" width="120" />

            <el-table-column align="left" label="预估花费" prop="estimatecompletion" width="120" />

            <el-table-column align="left" label="花费单位" prop="costunit" width="120" />

        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updateConstructionManageFunc(scope.row)">编辑</el-button>
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
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="工程名">
    {{ detailForm.title }}
</el-descriptions-item>
                    <el-descriptions-item label="工程描述">
    {{ detailForm.description }}
</el-descriptions-item>
                    <el-descriptions-item label="工程类型">
    {{ detailForm.type }}
</el-descriptions-item>
                    <el-descriptions-item label="项目位置">
    {{ detailForm.location }}
</el-descriptions-item>
                    <el-descriptions-item label="责任人">
    {{ detailForm.owner }}
</el-descriptions-item>
                    <el-descriptions-item label="设计单位">
    {{ detailForm.design }}
</el-descriptions-item>
                    <el-descriptions-item label="建造单位">
    {{ detailForm.contractor }}
</el-descriptions-item>
                    <el-descriptions-item label="管理单位">
    {{ detailForm.management }}
</el-descriptions-item>
                    <el-descriptions-item label="projectmanager字段">
    {{ detailForm.projectmanager }}
</el-descriptions-item>
                    <el-descriptions-item label="开工日期">
    {{ detailForm.scheduleddate }}
</el-descriptions-item>
                    <el-descriptions-item label="实际开开日期">
    {{ detailForm.actualstartdate }}
</el-descriptions-item>
                    <el-descriptions-item label="计划结束日期">
    {{ detailForm.plannedfinishdate }}
</el-descriptions-item>
                    <el-descriptions-item label="contractnumber字段">
    {{ detailForm.contractnumber }}
</el-descriptions-item>
                    <el-descriptions-item label="开工日期">
    {{ detailForm.initiationdate }}
</el-descriptions-item>
                    <el-descriptions-item label="项目展示图片">
    <el-image style="width: 50px; height: 50px; margin-right: 10px" :preview-src-list="returnArrImg(detailForm.images)" :initial-index="index" v-for="(item,index) in detailForm.images" :key="index" :src="getUrl(item)" fit="cover" />
</el-descriptions-item>
                    <el-descriptions-item label="项目文档">
    <div class="fileBtn" v-for="(item,index) in detailForm.pdfs" :key="index">
        <el-button type="primary" text bg @click="onDownloadFile(item.url)">
          <el-icon style="margin-right: 5px"><Download /></el-icon>
          {{ item.name }}
        </el-button>
    </div>
</el-descriptions-item>
                    <el-descriptions-item label="项目视频">
    {{ detailForm.video }}
</el-descriptions-item>
                    <el-descriptions-item label="vudgetedcost字段">
    {{ detailForm.vudgetedcost }}
</el-descriptions-item>
                    <el-descriptions-item label="实际花费">
    {{ detailForm.actualcost }}
</el-descriptions-item>
                    <el-descriptions-item label="预估花费">
    {{ detailForm.estimatecompletion }}
</el-descriptions-item>
                    <el-descriptions-item label="花费单位">
    {{ detailForm.costunit }}
</el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createConstructionManage,
  deleteConstructionManage,
  deleteConstructionManageByIds,
  updateConstructionManage,
  findConstructionManage,
  getConstructionManageList
} from '@/api/huasu/construction_manage'
import { getUrl } from '@/utils/image'
// 图片选择组件
import SelectImage from '@/components/selectImage/selectImage.vue'
// 文件选择组件
import SelectFile from '@/components/selectFile/selectFile.vue'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'ConstructionManage'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
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
  const table = await getConstructionManageList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
            deleteConstructionManageFunc(row)
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
      const res = await deleteConstructionManageByIds({ IDs })
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
const updateConstructionManageFunc = async(row) => {
    const res = await findConstructionManage({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteConstructionManageFunc = async (row) => {
    const res = await deleteConstructionManage({ ID: row.ID })
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
  const res = await findConstructionManage({ ID: row.ID })
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

.file-list{
  display: flex;
  flex-wrap: wrap;
  gap: 4px;
}

.fileBtn{
  margin-bottom: 10px;
}

.fileBtn:last-child{
  margin-bottom: 0;
}

</style>
