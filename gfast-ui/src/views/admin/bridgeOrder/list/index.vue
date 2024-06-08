<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
      <el-form-item  prop="sourceAddress">
        <el-input
            v-model="queryParams.sourceAddress"
            placeholder="请输入转出币种"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="targetAddress">
        <el-input
            v-model="queryParams.targetAddress"
            placeholder="请输入转入币种"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="sourceCoinAddress">
        <el-input
            v-model="queryParams.sourceCoinAddress"
            placeholder="请输入转出钱包地址"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item prop="targetCoinAddress">
        <el-input
            v-model="queryParams.targetCoinAddress"
            placeholder="请输入转入钱包地址"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
        <el-form-item label="转出链" prop="sourceChainId">
          <el-select v-model="queryParams.sourceChainId" placeholder="请选择转出链" clearable size="small">
              <el-option
                  v-for="item in sourceChainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
          </el-select>
        </el-form-item>
        <el-form-item label="转入链" prop="targetChainId">
          <el-select v-model="queryParams.targetChainId" placeholder="请选择转入链" clearable size="small">
              <el-option
                  v-for="item in targetChainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
          </el-select>
        </el-form-item>
      <el-form-item  prop="transactionHash">
        <el-input
            v-model="queryParams.transactionHash"
            placeholder="请输入交易哈希"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
<!--      <el-form-item label="数量" prop="amount">-->
<!--        <el-input-->
<!--            v-model="queryParams.amount"-->
<!--            placeholder="请输入数量"-->
<!--            clearable-->
<!--            size="small"-->
<!--            @keyup.enter.native="handleQuery"-->
<!--        />-->
<!--      </el-form-item>-->
      <el-form-item label="状态" prop="status">
        <el-select v-model="queryParams.status" placeholder="状态" clearable size="small">
            <el-option
                v-for="dict in statusOptions"
                :key="dict.key"
                :label="dict.value"
                :value="dict.key"
            />
        </el-select>
      </el-form-item>
<!--      <el-form-item label="发起时间" prop="createAt">-->
<!--        <el-date-picker-->
<!--          clearable-->
<!--          size="small"-->
<!--          style="width: 400px"-->
<!--          v-model="queryParams.startRange"-->
<!--          type="daterange"-->
<!--          range-separator="至"-->
<!--          start-placeholder="开始日期"-->
<!--          end-placeholder="结束日期"-->
<!--          value-format="yyyy-MM-dd"-->
<!--          placeholder="选择日期范围">-->
<!--        </el-date-picker>-->
<!--      </el-form-item>-->

<!--      <el-form-item label="结束时间" prop="createAt">-->
<!--        <el-date-picker-->
<!--          clearable-->
<!--          size="small"-->
<!--          style="width: 400px"-->
<!--          v-model="queryParams.endRange"-->
<!--          type="daterange"-->
<!--          range-separator="至"-->
<!--          start-placeholder="开始日期"-->
<!--          end-placeholder="结束日期"-->
<!--          value-format="yyyy-MM-dd"-->
<!--          placeholder="选择日期范围">-->
<!--        </el-date-picker>-->
<!--      </el-form-item>-->


<!--      <el-form-item label="发起时间2" prop="createAt">-->
<!--        <el-date-picker-->
<!--          clearable size="small" style="width: 200px"-->
<!--          v-model="queryParams.startTime2"-->
<!--          type="date"-->
<!--          value-format="yyyy-MM-dd"-->
<!--          placeholder="选择发起时间2">-->
<!--        </el-date-picker>-->
<!--      </el-form-item>-->
<!--      <el-form-item label="结束时间" prop="updateAt">-->
<!--        <el-date-picker-->
<!--            clearable size="small" style="width: 200px"-->
<!--            v-model="queryParams.updateAt"-->
<!--            type="date"-->
<!--            value-format="yyyy-MM-dd"-->
<!--            placeholder="选择结束时间">-->
<!--        </el-date-picker>-->
<!--      </el-form-item>-->

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
<!--        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>-->
      </el-form-item>
      <el-form-item>
        <el-button
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
          v-hasPermi="['admin/coin/export']"
        >导出</el-button>
      </el-form-item>

    </el-form>
    <el-row :gutter="10" class="mb8">
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="primary"-->
<!--          icon="el-icon-plus"-->
<!--          size="mini"-->
<!--          @click="handleAdd"-->
<!--          v-hasPermi="['admin/bridgeOrder/add']"-->
<!--        >新增</el-button>-->
<!--      </el-col>-->
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="success"-->
<!--          icon="el-icon-edit"-->
<!--          size="mini"-->
<!--          :disabled="single"-->
<!--          @click="handleUpdate"-->
<!--          v-hasPermi="['admin/bridgeOrder/edit']"-->
<!--        >修改</el-button>-->
<!--      </el-col>-->
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="danger"-->
<!--          icon="el-icon-delete"-->
<!--          size="mini"-->
<!--          :disabled="multiple"-->
<!--          @click="handleDelete"-->
<!--          v-hasPermi="['admin/bridgeOrder/delete']"-->
<!--        >删除</el-button>-->
<!--      </el-col>-->
    </el-row>
    <el-table v-loading="loading" :data="bridgeOrderList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="" align="center" prop="id" />
      <el-table-column label="转出币种" align="center" prop="sourceAddress" />
      <el-table-column label="转入币种" align="center" prop="targetAddress" />
      <el-table-column label="转出链" align="center" prop="sourceChainId" :formatter="sourceChainIdFormat" width="100">
        <template slot-scope="scope">
          {{ sourceChainIdFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="转入链" align="center" prop="targetChainId" :formatter="targetChainIdFormat" width="100">
        <template slot-scope="scope">
          {{ targetChainIdFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="数量" align="center" prop="amount" />
      <el-table-column label="手续费" align="center" prop="fee" />
      <el-table-column label="gas费" align="center" prop="gasFee" />
      <el-table-column label="转入钱包地址" align="center" prop="targetCoinAddress" />
      <el-table-column label="转出钱包地址" align="center" prop="sourceCoinAddress" />

      <el-table-column label="转出TXID" align="center" prop="outHash" />
      <el-table-column label="转入TXID" align="center" prop="inHash" />
<!--      <el-table-column label="跨链订单id" align="center" prop="orderId" />-->

      <el-table-column label="状态" align="center"  prop="status" :formatter="statusFormat" />
      <el-table-column label="发起时间" align="center" prop="createAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createAt, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="结束时间" align="center" prop="updateAt" width="180">
        <template slot-scope="scope">
        <span v-if="scope.row.status === 5">
          {{ parseTime(scope.row.updateAt, '{y}-{m}-{d}') }}
        </span>
          <span v-else>
          未完成..
        </span>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['admin/bridgeOrder/edit']"
          >修改</el-button>
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
    <!-- 添加或修改跨链记录对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
<!--      <el-form-item label="转出币种" prop="sourceAddress">-->
<!--           <el-input v-model="form.sourceAddress" placeholder="请输入转出币种" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="转入币种" prop="targetAddress">-->
<!--           <el-input v-model="form.targetAddress" placeholder="请输入转入币种" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="转出钱包地址" prop="sourceCoinAddress">-->
<!--           <el-input v-model="form.sourceCoinAddress" placeholder="请输入转出钱包地址" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="转入钱包地址" prop="targetCoinAddress">-->
<!--           <el-input v-model="form.targetCoinAddress" placeholder="请输入转入钱包地址" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="转出链" prop="sourceChainId">-->
<!--          <el-select v-model="form.sourceChainId" placeholder="请选择转出链">-->
<!--              <el-option-->
<!--                  v-for="item in sourceChainIdOptions"-->
<!--                  :key="item.key"-->
<!--                  :label="item.value"-->
<!--                  :value="item.key"-->
<!--              ></el-option>-->
<!--          </el-select>-->
<!--      </el-form-item>-->
<!--      <el-form-item label="转入链" prop="targetChainId">-->
<!--          <el-select v-model="form.targetChainId" placeholder="请选择转入链">-->
<!--              <el-option-->
<!--                  v-for="item in targetChainIdOptions"-->
<!--                  :key="item.key"-->
<!--                  :label="item.value"-->
<!--                  :value="item.key"-->
<!--              ></el-option>-->
<!--          </el-select>-->
<!--      </el-form-item>-->
<!--      <el-form-item label="交易哈希" prop="transactionHash">-->
<!--           <el-input v-model="form.transactionHash" placeholder="请输入交易哈希" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="跨链订单id" prop="orderId">-->
<!--           <el-input v-model="form.orderId" placeholder="请输入跨链订单id" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="数量" prop="amount">-->
<!--           <el-input v-model="form.amount" placeholder="请输入数量" />-->
<!--      </el-form-item>-->
      <el-form-item label="跨链记录状态" label-width="150px"  prop="status">
          <el-select v-model="form.status" placeholder="请选择跨链记录状态">
              <el-option
                  v-for="dict in statusOptions"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
      </el-form-item>
<!--       <el-form-item label="发起时间" prop="createAt">-->
<!--           <el-date-picker clearable size="small" style="width: 200px"-->
<!--               v-model="form.createAt"-->
<!--               type="date"-->
<!--               value-format="yyyy-MM-dd"-->
<!--               placeholder="选择发起时间">-->
<!--           </el-date-picker>-->
<!--       </el-form-item>-->
<!--       <el-form-item label="结束时间" prop="updateAt">-->
<!--           <el-date-picker clearable size="small" style="width: 200px"-->
<!--               v-model="form.updateAt"-->
<!--               type="date"-->
<!--               value-format="yyyy-MM-dd"-->
<!--               placeholder="选择结束时间">-->
<!--           </el-date-picker>-->
<!--       </el-form-item>-->
<!--      <el-form-item label="区块高度" prop="blockNumber">-->
<!--           <el-input v-model="form.blockNumber" placeholder="请输入区块高度" />-->
<!--      </el-form-item>-->
<!--      <el-form-item label="手续费" prop="fee">-->
<!--           <el-input v-model="form.fee" placeholder="请输入手续费" />-->
<!--      </el-form-item>-->
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
    listBridgeOrder,
    getBridgeOrder,
    delBridgeOrder,
    addBridgeOrder,
    updateBridgeOrder,
    listChain,
    changeBridgeOrderStatus,
    exportBridgeOrder,
} from "@/api/admin/bridgeOrder";
import {exportCoin} from "@/api/admin/coin";
export default {
  components:{},
  name: "BridgeOrder",
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
      // 跨链记录表格数据
      bridgeOrderList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // sourceChainIdOptions关联表数据
      sourceChainIdOptions: [],
      // targetChainIdOptions关联表数据
      targetChainIdOptions: [],
      // statusOptions字典数据
      statusOptions: [],
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        sourceAddress: undefined,
        targetAddress: undefined,
        sourceCoinAddress: undefined,
        targetCoinAddress: undefined,
        sourceChainId: undefined,
        targetChainId: undefined,
        transactionHash: undefined,
        status: undefined,
        startRange: undefined,
        endRange: [],
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        sourceAddress : [
          { required: true, message: "转出币种不能为空", trigger: "blur" }
        ],
        targetAddress : [
          { required: true, message: "转入币种不能为空", trigger: "blur" }
        ],
        sourceCoinAddress : [
          { required: true, message: "转出钱包地址不能为空", trigger: "blur" }
        ],
        targetCoinAddress : [
          { required: true, message: "转入钱包地址不能为空", trigger: "blur" }
        ],
        sourceChainId : [
          { required: true, message: "转出链不能为空", trigger: "blur" }
        ],
        targetChainId : [
          { required: true, message: "转入链不能为空", trigger: "blur" }
        ],
        transactionHash : [
          { required: true, message: "交易哈希不能为空", trigger: "blur" }
        ],
        orderId : [
          { required: true, message: "跨链订单id不能为空", trigger: "blur" }
        ],
        amount : [
          { required: true, message: "数量不能为空", trigger: "blur" }
        ],
        status : [
          { required: true, message: "跨链记录状态不能为空", trigger: "blur" }
        ],
        blockNumber : [
          { required: true, message: "区块高度不能为空", trigger: "blur" }
        ],
        fee : [
          { required: true, message: "手续费不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getChainItems()
    this.getChainItems()
    this.getDicts("bridge_order_status").then(response => {
      this.statusOptions = response.data.values||[];
    });
    this.getList();
  },
  methods: {
    //关联chain表选项
    getChainItems() {
      this.getItems(listChain, {pageSize:10000}).then(res => {
        this.sourceChainIdOptions = this.setItems(res, 'chainId', 'name')
        this.targetChainIdOptions = this.sourceChainIdOptions
      })
    },

    /** 查询跨链记录列表 */
    getList() {
      this.loading = true;
      listBridgeOrder(this.queryParams).then(response => {
        this.bridgeOrderList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 转出链关联表翻译
    sourceChainIdFormat(row, column) {
      return this.selectItemsLabel(this.sourceChainIdOptions, row.sourceChainId);
    },
    // 转入链关联表翻译
    targetChainIdFormat(row, column) {
      return this.selectItemsLabel(this.targetChainIdOptions, row.targetChainId);
    },
    // 跨链记录状态字典翻译
    statusFormat(row, column) {
      return this.selectDictLabel(this.statusOptions, row.status);
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
        sourceAddress: undefined,
        targetAddress: undefined,
        sourceCoinAddress: undefined,
        targetCoinAddress: undefined,
        sourceChainId: undefined,
        targetChainId: undefined,
        transactionHash: undefined,
        orderId: undefined,
        amount: undefined,
        status: undefined,
        createAt: undefined,
        updateAt: undefined,
        blockNumber: undefined,
        fee: undefined,
        remark: undefined,
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
      this.title = "添加跨链记录";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getBridgeOrder(id).then(response => {
        let data = response.data;
        data.sourceChainId = ''+data.sourceChainId
        data.targetChainId = ''+data.targetChainId
        data.status = ''+data.status
        this.form = data;
        this.open = true;
        this.title = "修改跨链记录状态";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateBridgeOrder(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addBridgeOrder(this.form).then(response => {
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
      this.$confirm('是否确认删除跨链记录编号为"' + ids + '"的数据项?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delBridgeOrder(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    },

    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams;
      this.$confirm('是否确认导出所有跨链记录?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        return exportBridgeOrder(queryParams);
      }).then(response => {
        this.download(response);
      }).catch(() => {});
    },

    download(data) {
      console.log('Received data:', data);
      console.log('Data type:', typeof data);

      const blob = new Blob([data], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' });
      const link = document.createElement('a');
      link.href = window.URL.createObjectURL(blob);
      link.download = 'bridgeOrder.xlsx';
      link.click();
      window.URL.revokeObjectURL(link.href);
    }

  }
};
</script>
