
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="任务编号:" prop="task_no">
    <el-input v-model="formData.task_no" :clearable="true" placeholder="请输入任务编号" />
</el-form-item>
        <el-form-item label="数量:" prop="num">
    <el-input v-model.number="formData.num" :clearable="true" placeholder="请输入数量" />
</el-form-item>
        <el-form-item label="重量:" prop="weight">
    <el-input-number v-model="formData.weight" style="width:100%" :precision="2" :clearable="true" />
</el-form-item>
        <el-form-item label="任务类型:" prop="test_task_status">
    <el-tree-select v-model="formData.test_task_status" placeholder="请选择任务类型" :data="test_task_statusOptions" style="width:100%" filterable :clearable="true" check-strictly></el-tree-select>
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
  createBTask,
  updateBTask,
  findBTask
} from '@/api/test_task/dtask'

defineOptions({
    name: 'BTaskForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
const test_task_statusOptions = ref([])
const formData = ref({
            task_no: '',
            num: undefined,
            weight: 0,
            test_task_status: '',
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findBTask({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    test_task_statusOptions.value = await getDictFunc('test_task_status')
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
               res = await createBTask(formData.value)
               break
             case 'update':
               res = await updateBTask(formData.value)
               break
             default:
               res = await createBTask(formData.value)
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
