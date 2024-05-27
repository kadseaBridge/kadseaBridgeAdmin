<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">    
      <el-form-item label="当前链" prop="sourceChainId">
        <el-input
            v-model="queryParams.sourceChainId"
            placeholder="请输入当前链"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>    
      <el-form-item label="目标链" prop="targetChainId">
        <el-input
            v-model="queryParams.targetChainId"
            placeholder="请输入目标链"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>    
      <el-form-item label="当前代币" prop="sourceCoinAddress">
        <el-input
            v-model="queryParams.sourceCoinAddress"
            placeholder="请输入当前代币"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>    
      <el-form-item label="是否可用" prop="isEnable">
        <el-input
            v-model="queryParams.isEnable"
            placeholder="请输入是否可用"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>    
      <el-form-item label="目标代币" prop="targetCoinAddress">
        <el-input
            v-model="queryParams.targetCoinAddress"
            placeholder="请输入目标代币"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>    
      <el-form-item label="" prop="updateAt">
        <el-date-picker
            clearable size="small" style="width: 200px"
            v-model="queryParams.updateAt"
            type="date"
            value-format="yyyy-MM-dd"
            placeholder="选择">
        </el-date-picker>
      </el-form-item>    
      <el-form-item label="" prop="createAt">
        <el-date-picker
            clearable size="small" style="width: 200px"
            v-model="queryParams.createAt"
            type="date"
            value-format="yyyy-MM-dd"
            placeholder="选择">
        </el-date-picker>
      </el-form-item>      
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          v-hasPermi="['work/bridgeConfig/add']"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
          v-hasPermi="['work/bridgeConfig/edit']"
        >修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
          v-hasPermi="['work/bridgeConfig/delete']"
        >删除</el-button>
      </el-col>
    </el-row>
    <el-table v-loading="loading" :data="bridgeConfigList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />      
      <el-table-column label="编号" align="center" prop="id" />      
      <el-table-column label="当前链" align="center" prop="sourceChainId" />      
      <el-table-column label="目标链" align="center" prop="targetChainId" />      
      <el-table-column label="当前代币" align="center" prop="sourceCoinAddress" />      
      <el-table-column label="固定金额手续费，" align="center" prop="feeFixed" />      
      <el-table-column label="手续费百分比，如果不需要则设置为0" align="center" prop="feePercent" />      
      <el-table-column label="单用户每日跨链不需要审核的总额度" align="center" prop="dayTotal" />      
      <el-table-column label="单用户每次跨链不需要审核的额度" align="center" prop="onceTotal" />      
      <el-table-column label="是否可用" align="center" prop="isEnable" :formatter="isEnableFormat" />      
      <el-table-column label="目标代币" align="center" prop="targetCoinAddress" />      
      <el-table-column label="" align="center" prop="updateAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.updateAt, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="createAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createAt, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>      
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['work/bridgeConfig/edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['work/bridgeConfig/delete']"
          >删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
    <!-- 添加或修改可跨币种设置对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">      
      <el-form-item label="当前链" prop="sourceChainId">
           <el-input v-model="form.sourceChainId" placeholder="请输入当前链" />
      </el-form-item>      
      <el-form-item label="目标链" prop="targetChainId">
           <el-input v-model="form.targetChainId" placeholder="请输入目标链" />
      </el-form-item>      
      <el-form-item label="当前代币" prop="sourceCoinAddress">
           <el-input v-model="form.sourceCoinAddress" placeholder="请输入当前代币" />
      </el-form-item>      
      <el-form-item label="固定金额手续费，" prop="feeFixed">
           <el-input v-model="form.feeFixed" placeholder="请输入固定金额手续费，" />
      </el-form-item>      
      <el-form-item label="手续费百分比，如果不需要则设置为0" prop="feePercent">
           <el-input v-model="form.feePercent" placeholder="请输入手续费百分比，如果不需要则设置为0" />
      </el-form-item>      
      <el-form-item label="单用户每日跨链不需要审核的总额度" prop="dayTotal">
           <el-input v-model="form.dayTotal" placeholder="请输入单用户每日跨链不需要审核的总额度" />
      </el-form-item>      
      <el-form-item label="单用户每次跨链不需要审核的额度" prop="onceTotal">
           <el-input v-model="form.onceTotal" placeholder="请输入单用户每次跨链不需要审核的额度" />
      </el-form-item>      
      <el-form-item label="是否可用" prop="isEnable">
           <el-input v-model="form.isEnable" placeholder="请输入是否可用" />
      </el-form-item>      
      <el-form-item label="目标代币" prop="targetCoinAddress">
           <el-input v-model="form.targetCoinAddress" placeholder="请输入目标代币" />
      </el-form-item>      
       <el-form-item label="" prop="updateAt">
           <el-date-picker clearable size="small" style="width: 200px"
               v-model="form.updateAt"
               type="date"
               value-format="yyyy-MM-dd"
               placeholder="选择">
           </el-date-picker>
       </el-form-item>      
       <el-form-item label="" prop="createAt">
           <el-date-picker clearable size="small" style="width: 200px"
               v-model="form.createAt"
               type="date"
               value-format="yyyy-MM-dd"
               placeholder="选择">
           </el-date-picker>
       </el-form-item>       
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
    listBridgeConfig,
    getBridgeConfig,
    delBridgeConfig,
    addBridgeConfig,
    updateBridgeConfig,    
} from "@/api/work/bridgeConfig";
export default {
  components:{},
  name: "BridgeConfig",
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
      // 可跨币种设置表格数据
      bridgeConfigList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,      
      // isEnableOptions字典数据
      isEnableOptions: [],      
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        sourceChainId: undefined,
        targetChainId: undefined,
        sourceCoinAddress: undefined,
        isEnable: undefined,
        targetCoinAddress: undefined,
        updateAt: undefined,
        createAt: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: { 
        sourceChainId : [
          { required: true, message: "当前链不能为空", trigger: "blur" }
        ],
        targetChainId : [
          { required: true, message: "目标链不能为空", trigger: "blur" }
        ],
        sourceCoinAddress : [
          { required: true, message: "当前代币不能为空", trigger: "blur" }
        ],
        feeFixed : [
          { required: true, message: "固定金额手续费，不能为空", trigger: "blur" }
        ],
        feePercent : [
          { required: true, message: "手续费百分比，如果不需要则设置为0不能为空", trigger: "blur" }
        ],
        dayTotal : [
          { required: true, message: "单用户每日跨链不需要审核的总额度不能为空", trigger: "blur" }
        ],
        onceTotal : [
          { required: true, message: "单用户每次跨链不需要审核的额度不能为空", trigger: "blur" }
        ],
        isEnable : [
          { required: true, message: "是否可用不能为空", trigger: "blur" }
        ],
        targetCoinAddress : [
          { required: true, message: "目标代币不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {    
    this.getDicts("coin_able").then(response => {
      this.isEnableOptions = response.data.values||[];
    });
    this.getList();
  },
  methods: {    
    /** 查询可跨币种设置列表 */
    getList() {
      this.loading = true;
      listBridgeConfig(this.queryParams).then(response => {
        this.bridgeConfigList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },    
    // 是否可用字典翻译
    isEnableFormat(row, column) {
      return this.selectDictLabel(this.isEnableOptions, row.isEnable);
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
        sourceChainId: undefined,        
        targetChainId: undefined,        
        sourceCoinAddress: undefined,        
        feeFixed: undefined,        
        feePercent: undefined,        
        dayTotal: undefined,        
        onceTotal: undefined,        
        isEnable: undefined,        
        targetCoinAddress: undefined,        
        updateAt: undefined,        
        createAt: undefined,        
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
      this.title = "添加可跨币种设置";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getBridgeConfig(id).then(response => {
        let data = response.data;        
        this.form = data;
        this.open = true;
        this.title = "修改可跨币种设置";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateBridgeConfig(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addBridgeConfig(this.form).then(response => {
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
      this.$confirm('是否确认删除可跨币种设置编号为"' + ids + '"的数据项?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delBridgeConfig(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    }
  }
};
</script>