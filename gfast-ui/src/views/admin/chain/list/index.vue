<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
      <el-form-item label="链名" prop="name">
        <el-input
            v-model="queryParams.name"
            placeholder="请输入链名"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
<!--        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>-->
      </el-form-item>
    </el-form>
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          v-hasPermi="['admin/chain/add']"
        >新增</el-button>
      </el-col>
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="success"-->
<!--          icon="el-icon-edit"-->
<!--          size="mini"-->
<!--          :disabled="single"-->
<!--          @click="handleUpdate"-->
<!--          v-hasPermi="['admin/chain/edit']"-->
<!--        >修改</el-button>-->
<!--      </el-col>-->
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="danger"-->
<!--          icon="el-icon-delete"-->
<!--          size="mini"-->
<!--          :disabled="multiple"-->
<!--          @click="handleDelete"-->
<!--          v-hasPermi="['admin/chain/delete']"-->
<!--        >删除</el-button>-->
<!--      </el-col>-->
    </el-row>
    <el-table v-loading="loading" :data="chainList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
<!--      <el-table-column label="" align="center" prop="id" />-->
      <el-table-column label="序号" width="50">
        <template v-slot="scope">
          {{ scope.$index + 1 }}
        </template>
      </el-table-column>
      <el-table-column label="链名" align="center" prop="name" />
      <el-table-column label="链Id" align="center" prop="chainId" />
      <el-table-column label="链rpc" align="center" prop="rpc" />
      <el-table-column label="链类型" align="center" prop="type" :formatter="typeFormat" />
      <el-table-column label="跨链桥合约地址" align="center" prop="bridgeContractAddress"  />
      <el-table-column label="maxGasPrice" align="center" prop="maxGasPrice" />
      <el-table-column label="enablePay" align="center" prop="enablePay" />
      <el-table-column label="链图标" align="center" prop="icon" />
      <el-table-column label="状态" align="center" prop="isEnable" :formatter="isEnableFormat" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['admin/chain/edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['admin/chain/delete']"
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
    <!-- 添加或修改链管理对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
      <el-form-item label="链名" prop="name">
           <el-input v-model="form.name" placeholder="请输入链名" />
      </el-form-item>
      <el-form-item label="链Id" prop="chainId">
           <el-input v-model="form.chainId" placeholder="请输入链Id" />
      </el-form-item>
      <el-form-item label="链rpc" prop="rpc">
           <el-input v-model="form.rpc" placeholder="请输入链rpc" />
      </el-form-item>
      <el-form-item label="链浏览器地址" prop="explorer">
           <el-input v-model="form.explorer" placeholder="请输入链浏览器地址" />
      </el-form-item>
      <el-form-item label="状态" prop="isEnable">
          <el-select v-model="form.isEnable" placeholder="请选择状态">
              <el-option
                  v-for="dict in isEnableOptions"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
      </el-form-item>
<!--       <el-form-item label="创建时间" prop="createAt">-->
<!--           <el-date-picker clearable size="small" style="width: 200px"-->
<!--               v-model="form.createAt"-->
<!--               type="date"-->
<!--               value-format="yyyy-MM-dd"-->
<!--               placeholder="选择创建时间">-->
<!--           </el-date-picker>-->
<!--       </el-form-item>-->
<!--       <el-form-item label="更新时间" prop="updateAt">-->
<!--           <el-date-picker clearable size="small" style="width: 200px"-->
<!--               v-model="form.updateAt"-->
<!--               type="date"-->
<!--               value-format="yyyy-MM-dd"-->
<!--               placeholder="选择更新时间">-->
<!--           </el-date-picker>-->
<!--       </el-form-item>-->
      <el-form-item label="链类型" prop="type">
          <el-select v-model="form.type" placeholder="链类型">
              <el-option
                  v-for="dict in typeOptions"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
      </el-form-item>
      <el-form-item label="链图标" prop="icon">
           <el-input v-model="form.icon" placeholder="请输入链图标" />
      </el-form-item>
      <el-form-item label="软删除" prop="isDelete">
          <el-select v-model="form.isDelete" placeholder="请选择软删除">
              <el-option
                  v-for="dict in isDeleteOptions"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
      </el-form-item>
<!--      <el-form-item label="maxGasPrice" prop="maxGasPrice">-->
<!--           <el-input v-model="form.maxGasPrice" placeholder="请输入支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时， 为0时，所有订单都不进行支付， 如果想要所有的订单都进行支付，不管gasprice 则设置很大   单位：wei" />-->
<!--      </el-form-item>-->

        <el-form-item label="最大接受的gasprice" prop="maxGasPrice" label-width="150px">
          <el-tooltip
            content="请输入支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时，所有订单都不进行支付，如果想要所有的订单都进行支付，不管gasprice 则设置很大。单位：wei"
            placement="top"
          >
            <el-input v-model="form.maxGasPrice" placeholder="请输入支付时，最大接受的gasprice" />
          </el-tooltip>
        </el-form-item>
      <el-form-item label="enablePay" prop="enablePay" label-width="130px">
           <el-input v-model="form.enablePay" placeholder="enablePay" />
      </el-form-item>
      <el-form-item label="跨链桥合约地址" prop="bridgeContractAddress" label-width="130px">
           <el-input v-model="form.bridgeContractAddress"  placeholder="请输入跨链桥合约地址" />
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
    listChain,
    getChain,
    delChain,
    addChain,
    updateChain,
} from "@/api/admin/chain";
export default {
  components:{},
  name: "Chain",
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
      // 链管理表格数据
      chainList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // isEnableOptions字典数据
      isEnableOptions: [],
      // typeOptions字典数据
      typeOptions: [],
      // isDeleteOptions字典数据
      isDeleteOptions: [],
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        name: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        name : [
          { required: true, message: "链名不能为空", trigger: "blur" }
        ],
        chainId : [
          { required: true, message: "链Id不能为空", trigger: "blur" }
        ],
        rpc : [
          { required: true, message: "链rpc不能为空", trigger: "blur" }
        ],
        isEnable : [
          { required: true, message: "状态不能为空", trigger: "blur" }
        ],
        type : [
          { required: true, message: "链类型", trigger: "blur" }
        ],
        maxGasPrice : [
          { required: true, message: "支付时，该公链最大接受的gasprice,如果超过则不进行目标链支付,为0时， 为0时，所有订单都不进行支付， 如果想要所有的订单都进行支付，不管gasprice 则设置很大   单位：wei不能为空", trigger: "blur" }
        ],
        enablePay : [
          { required: true, message: "该公链作为目标时，是否允许支付不能为空", trigger: "blur" }
        ],
        bridgeContractAddress : [
          { required: true, message: "跨链桥合约地址不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getDicts("is_enable").then(response => {
      this.isEnableOptions = response.data.values||[];
    });
    this.getDicts("chain_type").then(response => {
      this.typeOptions = response.data.values||[];
    });
    this.getDicts("isDelete").then(response => {
      this.isDeleteOptions = response.data.values||[];
    });
    this.getList();
  },
  methods: {
    /** 查询链管理列表 */
    getList() {
      this.loading = true;
      listChain(this.queryParams).then(response => {
        this.chainList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 状态字典翻译
    isEnableFormat(row, column) {
      return this.selectDictLabel(this.isEnableOptions, row.isEnable);
    },
    // 0 ethereum 1 tron  2 btc字典翻译
    typeFormat(row, column) {
      return this.selectDictLabel(this.typeOptions, row.type);
    },
    // 软删除字典翻译
    isDeleteFormat(row, column) {
      return this.selectDictLabel(this.isDeleteOptions, row.isDelete);
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
        name: undefined,
        chainId: undefined,
        rpc: undefined,
        explorer: undefined,
        isEnable: undefined,
        createAt: undefined,
        updateAt: undefined,
        type: undefined,
        icon: undefined,
        isDelete: undefined,
        maxGasPrice: undefined,
        enablePay: undefined,
        bridgeContractAddress: undefined,
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
      this.title = "添加链";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getChain(id).then(response => {
        let data = response.data;
        data.isEnable = ''+data.isEnable
        data.type = ''+data.type
        data.isDelete = ''+data.isDelete
        this.form = data;
        this.open = true;
        this.title = "修改链管理";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateChain(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addChain(this.form).then(response => {
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
      this.$confirm('是否确认删除该链?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delChain(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    }
  }
};
</script>
