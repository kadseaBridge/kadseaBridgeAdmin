<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
      <el-form-item label="日期" prop="date">
        <el-date-picker
            clearable size="small" style="width: 200px"
            v-model="queryParams.date"
            type="date"
            value-format="yyyy-MM-dd"
            placeholder="选择日期">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="币种" prop="coinName">
        <el-input
            v-model="queryParams.coinName"
            placeholder="请输入币种"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="链类型" prop="chainName">
        <el-input
            v-model="queryParams.chainName"
            placeholder="请输入链类型"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
<!--    <el-row :gutter="10" class="mb8">-->
<!--&lt;!&ndash;      <el-col :span="1.5">&ndash;&gt;-->
<!--&lt;!&ndash;        <el-button&ndash;&gt;-->
<!--&lt;!&ndash;          type="primary"&ndash;&gt;-->
<!--&lt;!&ndash;          icon="el-icon-plus"&ndash;&gt;-->
<!--&lt;!&ndash;          size="mini"&ndash;&gt;-->
<!--&lt;!&ndash;          @click="handleAdd"&ndash;&gt;-->
<!--&lt;!&ndash;          v-hasPermi="['admin/dailyBridgeStats/add']"&ndash;&gt;-->
<!--&lt;!&ndash;        >新增</el-button>&ndash;&gt;-->
<!--&lt;!&ndash;      </el-col>&ndash;&gt;-->
<!--&lt;!&ndash;      <el-col :span="1.5">&ndash;&gt;-->
<!--&lt;!&ndash;        <el-button&ndash;&gt;-->
<!--&lt;!&ndash;          type="success"&ndash;&gt;-->
<!--&lt;!&ndash;          icon="el-icon-edit"&ndash;&gt;-->
<!--&lt;!&ndash;          size="mini"&ndash;&gt;-->
<!--&lt;!&ndash;          :disabled="single"&ndash;&gt;-->
<!--&lt;!&ndash;          @click="handleUpdate"&ndash;&gt;-->
<!--&lt;!&ndash;          v-hasPermi="['admin/dailyBridgeStats/edit']"&ndash;&gt;-->
<!--&lt;!&ndash;        >修改</el-button>&ndash;&gt;-->
<!--&lt;!&ndash;      </el-col>&ndash;&gt;-->
<!--&lt;!&ndash;      <el-col :span="1.5">&ndash;&gt;-->
<!--&lt;!&ndash;        <el-button&ndash;&gt;-->
<!--&lt;!&ndash;          type="danger"&ndash;&gt;-->
<!--&lt;!&ndash;          icon="el-icon-delete"&ndash;&gt;-->
<!--&lt;!&ndash;          size="mini"&ndash;&gt;-->
<!--&lt;!&ndash;          :disabled="multiple"&ndash;&gt;-->
<!--&lt;!&ndash;          @click="handleDelete"&ndash;&gt;-->
<!--&lt;!&ndash;          v-hasPermi="['admin/dailyBridgeStats/delete']"&ndash;&gt;-->
<!--&lt;!&ndash;        >删除</el-button>&ndash;&gt;-->
<!--&lt;!&ndash;      </el-col>&ndash;&gt;-->
<!--    </el-row>-->
    <el-table v-loading="loading" :data="dailyBridgeStatsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="id" align="center" prop="id" />
      <el-table-column label="日期" align="center" prop="date" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.date, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="币种" align="center" prop="coinName" />
      <el-table-column label="链类型" align="center" prop="chainName" />
      <el-table-column label="跨链转入" align="center" prop="transferIn" />
      <el-table-column label="跨链转出" align="center" prop="transferOut" />
      <el-table-column label="跨链差额" align="center" prop="transferDifference" />
      <el-table-column label="跨链手续费" align="center" prop="fee" />
      <el-table-column label="平台资产" align="center" prop="platformAssets" />
<!--      <el-table-column label="财务余额" align="center" prop="financialBalance" />-->
<!--      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">-->
<!--        <template slot-scope="scope">-->
<!--          <el-button-->
<!--            size="mini"-->
<!--            type="text"-->
<!--            icon="el-icon-edit"-->
<!--            @click="handleUpdate(scope.row)"-->
<!--            v-hasPermi="['admin/dailyBridgeStats/edit']"-->
<!--          >修改</el-button>-->
<!--          <el-button-->
<!--            size="mini"-->
<!--            type="text"-->
<!--            icon="el-icon-delete"-->
<!--            @click="handleDelete(scope.row)"-->
<!--            v-hasPermi="['admin/dailyBridgeStats/delete']"-->
<!--          >删除</el-button>-->
<!--        </template>-->
<!--      </el-table-column>-->
    </el-table>
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
    <!-- 添加或修改跨链日统计对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button type="primary" @click="submitForm">确 定</el-button>
        <el-button @click="cancel">取 消</el-button>
      </div>
    </el-dialog>
  </div>
</template>
<script>
import {
    listDailyBridgeStats,
    getDailyBridgeStats,
    delDailyBridgeStats,
    addDailyBridgeStats,
    updateDailyBridgeStats,
} from "@/api/admin/dailyBridgeStats";
export default {
  components:{},
  name: "DailyBridgeStats",
  data() {
    return {
      // 遮罩层
      loading: true,
      // 选中数组
      ids: [],
      // 非单个禁用
      single: true,
      // 非多个禁用
      multiple: true,
      // 总条数
      total: 0,
      // 跨链日统计表格数据
      dailyBridgeStatsList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        date: undefined,
        coinName: undefined,
        chainName: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        date : [
          { required: true, message: "日期不能为空", trigger: "blur" }
        ],
        coinName : [
          { required: true, message: "币种不能为空", trigger: "blur" }
        ],
        chainName : [
          { required: true, message: "链类型不能为空", trigger: "blur" }
        ],
        transferIn : [
          { required: true, message: "跨链转入不能为空", trigger: "blur" }
        ],
        transferOut : [
          { required: true, message: "跨链转出不能为空", trigger: "blur" }
        ],
        transferDifference : [
          { required: true, message: "跨链差额不能为空", trigger: "blur" }
        ],
        fee : [
          { required: true, message: "跨链手续费不能为空", trigger: "blur" }
        ],
        platformAssets : [
          { required: true, message: "平台资产不能为空", trigger: "blur" }
        ],
        financialBalance : [
          { required: true, message: "财务余额不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getList();
  },
  methods: {
    /** 查询跨链日统计列表 */
    getList() {
      this.loading = true;
      listDailyBridgeStats(this.queryParams).then(response => {
        this.dailyBridgeStatsList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 取消按钮
    cancel() {
      this.open = false;
      this.reset();
    },
    // 表单重置
    reset() {
      this.form = {
        id: undefined,
        date: undefined,
        coinName: undefined,
        chainName: undefined,
        transferIn: undefined,
        transferOut: undefined,
        transferDifference: undefined,
        fee: undefined,
        platformAssets: undefined,
        financialBalance: undefined,
        createdAt: undefined,
        updatedAt: undefined,
      };
      this.resetForm("form");
    },
    /** 搜索按钮操作 */
    handleQuery() {
      this.queryParams.pageNum = 1;
      this.getList();
    },
    /** 重置按钮操作 */
    resetQuery() {
      this.resetForm("queryForm");
      this.handleQuery();
    },
    // 多选框选中数据
    handleSelectionChange(selection) {
      this.ids = selection.map(item => item.id)
      this.single = selection.length!=1
      this.multiple = !selection.length
    },
    /** 新增按钮操作 */
    handleAdd() {
      this.reset();
      this.open = true;
      this.title = "添加跨链日统计";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getDailyBridgeStats(id).then(response => {
        let data = response.data;
        this.form = data;
        this.open = true;
        this.title = "修改跨链日统计";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateDailyBridgeStats(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addDailyBridgeStats(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("新增成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          }
        }
      });
    },
    /** 删除按钮操作 */
    handleDelete(row) {
      const ids = row.id || this.ids;
      this.$confirm('是否确认删除跨链日统计编号为"' + ids + '"的数据项?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delDailyBridgeStats(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    }
  }
};
</script>
