<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
        <el-form-item label="当前链" prop="sourceChainId">
          <el-select v-model="queryParams.sourceChainId" placeholder="请选择当前链" clearable size="small"  @change="onSourceChainChange">
              <el-option
                  v-for="item in sourceChainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
          </el-select>
        </el-form-item>
        <el-form-item label="币种" prop="sourceCoinAddress">
          <el-select v-model="queryParams.sourceCoinAddress" placeholder="请选择币种" clearable size="small" @change="onSourceCoinChange">
              <el-option
                  v-for="item in filteredCoinOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
          </el-select>
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
          v-hasPermi="['admin/bridgeConfig/add']"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="warning"
          icon="el-icon-download"
          size="mini"
          @click="handleExport"
          v-hasPermi="['admin/coin/export']"
        >导出</el-button>
      </el-col>

    </el-row>
    <el-table v-loading="loading" :data="bridgeConfigList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
       <el-table-column label="序号" align="center" prop="id" />
<!--      <el-table-column label="序号" width="50">-->
<!--        <template v-slot="scope">-->
<!--          {{ scope.$index + 1 }}-->
<!--        </template>-->
<!--      </el-table-column>-->
      <el-table-column label="当前链" align="center" prop="sourceChainId" :formatter="sourceChainIdFormat" width="150">
        <template slot-scope="scope">
          {{ sourceChainIdFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="币种" align="center" prop="sourceCoinAddress" :formatter="sourceCoinAddressFormat" width="150">
        <template slot-scope="scope">
          {{ sourceCoinAddressFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="对手链" align="center" prop="targetChainId" :formatter="targetChainIdFormat" width="150">
        <template slot-scope="scope">
          {{ targetChainIdFormat(scope.row) }}
        </template>
      </el-table-column>

      <el-table-column label="跨入手续费" align="center">
        <template v-slot="scope">
          {{ calculateTotalFee(scope.row) }}
        </template>
      </el-table-column>

      <!-- <el-table-column label="固定金额手续费，" align="center" prop="feeFixed" />
      <el-table-column label="手续费百分比，如果不需要则设置为0" align="center" prop="feePercent" /> -->
      <el-table-column label="跨入每日审核数量" align="center" prop="dayTotal" />
      <el-table-column label="跨入单次审核数量" align="center" prop="onceTotal" />
      <el-table-column label="状态" align="center" prop="isEnable" :formatter="isEnableFormat" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['admin/bridgeConfig/edit']"
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
    <!-- 添加或修改跨链币对对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
      <el-form-item label="当前链" prop="sourceChainId">
          <el-select v-model="form.sourceChainId" placeholder="请选择当前链"  @change="onSourceChainChangeDialog">
              <el-option
                  v-for="item in sourceChainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
      </el-form-item>
      <el-form-item label="币种" prop="sourceCoinAddress">
          <el-select v-model="form.sourceCoinAddress" placeholder="请选择币种" @change="onSourceCoinChangeDialog">
              <el-option
                  v-for="item in filteredCoinOptionsDialog"
                  :key="item.key"
                  :label="item.value"
                  :value="item.value"
              ></el-option>

          </el-select>

      </el-form-item>
      <el-form-item label="对手链" prop="targetChainId">
          <el-select v-model="form.targetChainId" placeholder="请选择对手链">
              <el-option
                  v-for="item in filteredTargetChainOptionsDialog"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
      </el-form-item>

      <el-form-item label="固定金额手续费：" label-width="150px" prop="feeFixed">
           <el-input v-model="form.feeFixed" placeholder="请输入固定金额手续费，" />
      </el-form-item>
      <el-form-item label="手续费百分比：" label-width="150px" prop="feePercent">
           <el-input v-model="form.feePercent" placeholder="请输入手续费百分比，如果不需要则设置为0" />
      </el-form-item>
      <el-form-item label="跨入每日审核数量" label-width="150px" prop="dayTotal">
           <el-input v-model="form.dayTotal" placeholder="请输入跨入每日审核数量" />
      </el-form-item>
      <el-form-item label="跨入单次审核数量" label-width="150px" prop="onceTotal">
           <el-input v-model="form.onceTotal" placeholder="请输入跨入单次审核数量" />
      </el-form-item>
       <el-form-item label="状态" prop="isEnable">
           <el-radio-group v-model="form.isEnable">
               <el-radio
                v-for="dict in isEnableOptions"
                :key="dict.key"
                :label="dict.key"
               >{{dict.value }}</el-radio>
           </el-radio-group>
       </el-form-item>
<!--      <el-form-item label="目标链合约地址" prop="targetCoinAddress">-->
<!--           <el-input v-model="form.targetCoinAddress" placeholder="请输入目标链合约地址" />-->
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
    listBridgeConfig,
    getBridgeConfig,
    delBridgeConfig,
    addBridgeConfig,
    updateBridgeConfig,
    listChain,
    listCoin,
    exportBridgeConfig
} from "@/api/admin/bridgeConfig";
import {exportCoin} from "@/api/admin/coin";
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
      // 跨链币对表格数据
      bridgeConfigList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // sourceChainIdOptions关联表数据
      sourceChainIdOptions: [],
      // targetChainIdOptions关联表数据
      targetChainIdOptions: [],
      // sourceCoinAddressOptions关联表数据
      sourceCoinAddressOptions: [],
      // isEnableOptions字典数据
      isEnableOptions: [],

      filteredCoinOptions: [],
      filteredCoinOptionsDialog: [],
      filteredTargetChainOptionsDialog: [],
      filteredTargetChainOptions:[],

      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        sourceChainId: undefined,
        sourceCoinAddress: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        sourceChainId : [
          { required: true, message: "当前链不能为空", trigger: "blur" }
        ],
        targetChainId : [
          { required: true, message: "对手链不能为空", trigger: "blur" }
        ],
        sourceCoinAddress : [
          { required: true, message: "币种不能为空", trigger: "blur" }
        ],
        feeFixed : [
          { required: true, message: "固定金额手续费，不能为空", trigger: "blur" }
        ],
        feePercent : [
          { required: true, message: "手续费百分比，如果不需要则设置为0不能为空", trigger: "blur" }
        ],
        dayTotal : [
          { required: true, message: "跨入每日审核数量不能为空", trigger: "blur" }
        ],
        onceTotal : [
          { required: true, message: "跨入单次审核数量不能为空", trigger: "blur" }
        ],
        isEnable : [
          { required: true, message: "状态不能为空", trigger: "blur" }
        ],

      }
    };
  },
  created() {
    this.getChainItems()
    this.getCoinItems()
    this.getDicts("is_enable").then(response => {
      this.isEnableOptions = response.data.values||[];
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
      console.log('Coin items response:', this.sourceChainIdOptions);
    },


    calculateTotalFee(row) {
      return `${row.feePercent}% + ${row.feeFixed}`;
    },
    //关联coin表选项

    getCoinItems() {
      this.getItems(listCoin, { pageSize: 10000 }).then(res => {
        console.log('Coin items response:', res);
        if (res && res.data && res.data.list) {
          this.sourceCoinAddressOptions = res.data.list.map(item => ({
            key: item.address,
            value: item.symbol,
            chainId: item.chainId // 确保每个币种有chainId字段
          }));
        } else {
          console.error('Invalid coin items response:', res);
        }
      }).catch(error => {
        console.error('Error fetching coin items:', error);
      });
    },
    /** 查询跨链币对列表 */
    getList() {
      this.loading = true;
      listBridgeConfig(this.queryParams).then(response => {
        this.bridgeConfigList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 当前链关联表翻译
    sourceChainIdFormat(row, column) {
      return this.selectItemsLabel(this.sourceChainIdOptions, row.sourceChainId);
    },
    // 对手链关联表翻译
    targetChainIdFormat(row, column) {
      return this.selectItemsLabel(this.targetChainIdOptions, row.targetChainId);
    },
    // 币种关联表翻译
    sourceCoinAddressFormat(row, column) {
      return this.selectItemsLabel2(this.sourceCoinAddressOptions, row.sourceCoinAddress, row.sourceChainId);
    },
    // 状态字典翻译
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
        isEnable: "0" ,
        // targetCoinAddress: undefined,
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
    // 当前链选择变化
    onSourceChainChange(value) {
      this.filteredCoinOptions = this.sourceCoinAddressOptions.filter(item => item.chainId === value);
      this.queryParams.sourceCoinAddress = undefined; // 重置币种选择框
      this.filteredTargetChainOptions = []; // 重置对手链选择框
      this.queryParams.targetChainId = undefined; // 重置对手链选择框
    },
    // 币种选择变化
    onSourceCoinChange(value) {
      this.filteredTargetChainOptions = this.sourceCoinAddressOptions.filter(item => item.value.includes(value));
      this.queryParams.targetChainId = undefined; // 重置对手链选择框
    },
    // 当前链选择变化（对话框）
    onSourceChainChangeDialog(value) {
      this.filteredCoinOptionsDialog = this.sourceCoinAddressOptions.filter(item => item.chainId === value);
      this.form.sourceCoinAddress = undefined; // 重置币种选择框
      this.filteredTargetChainOptionsDialog = []; // 重置对手链选择框
      this.form.targetChainId = undefined; // 重置对手链选择框
    },

    // 币种选择变化（对话框）
    onSourceCoinChangeDialog(value) {
      const filteredCoinRecords = this.sourceCoinAddressOptions.filter(item => item.value.includes(value));
      const chainIds = filteredCoinRecords.map(item => item.chainId);
      this.filteredTargetChainOptionsDialog = this.sourceChainIdOptions.filter(option => chainIds.includes(option.key) && option.key !== this.form.sourceChainId);
      this.form.targetChainId = undefined; // 重置对手链选择框
    },

    selectItemsLabel2(datas, address, chainId) {
      for (let item of datas) {
        if (item.key === String(address) && item.chainId === String(chainId)) {
          return item.value;
        }
      }
      return null; // 或者返回一个默认值
    },

    /** 重置按钮操作 */
    // resetQuery() {
    //   this.resetForm("queryForm");
    //   this.handleQuery();
    // },
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
      this.title = "添加跨链币对";

    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getBridgeConfig(id).then(response => {
        let data = response.data;
        data.sourceChainId = ''+data.sourceChainId
        data.targetChainId = ''+data.targetChainId
        data.sourceCoinAddress = ''+data.sourceCoinAddress
        data.isEnable = ''+data.isEnable
        this.form = data;
        this.open = true;
        this.title = "修改跨链币对";
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
      this.$confirm('是否确认删除跨链币对?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delBridgeConfig(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    },

    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams;
      this.$confirm('是否确认导出所有跨链币对?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        return exportBridgeConfig(queryParams);
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
      link.download = '跨链币对表.xlsx';
      link.click();
      window.URL.revokeObjectURL(link.href);
    }
  }
};
</script>
