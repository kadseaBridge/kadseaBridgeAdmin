<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
      <el-form-item  prop="sourceCoinAddress">
        <el-input
          v-model="queryParams.sourceCoinAddress"
          placeholder="转出币种"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item prop="targetCoinAddress">
        <el-input
          v-model="queryParams.targetCoinAddress"
          placeholder="转入币种"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="sourceAddress">
        <el-input
            v-model="queryParams.sourceAddress"
            placeholder="转出钱包地址"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="targetAddress">
        <el-input
            v-model="queryParams.targetAddress"
            placeholder="转入钱包地址"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="inHash">
        <el-input
          v-model="queryParams.inHash"
          placeholder="转入TXID"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item  prop="outHash">
        <el-input
          v-model="queryParams.outHash"
          placeholder="转出TXID"
          clearable
          size="small"
          @keyup.enter.native="handleQuery"
        />
      </el-form-item>

<!--      <el-form-item prop="targetCoinAddress">-->
<!--        <el-input-->
<!--            v-model="queryParams.targetCoinAddress"-->
<!--            placeholder="转入币种地址"-->
<!--            clearable-->
<!--            size="small"-->
<!--            @keyup.enter.native="handleQuery"-->
<!--        />-->
<!--      </el-form-item>-->
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
      <el-form-item label="发起时间">
        <el-date-picker
          clearable size="small" style="width: 200px"
          v-model="queryParams.startDate1"
          type="date"
          value-format="yyyy-MM-dd"
          placeholder="选择发起时间">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="发起时间">
        <el-date-picker
          clearable size="small" style="width: 200px"
          v-model="queryParams.startDate2"
          type="date"
          value-format="yyyy-MM-dd"
          placeholder="选择发起时间">
        </el-date-picker>
      </el-form-item>

      <el-form-item label="结束时间">
        <el-date-picker
          clearable size="small" style="width: 200px"
          v-model="queryParams.endDate1"
          type="date"
          value-format="yyyy-MM-dd"
          placeholder="选择结束时间">
        </el-date-picker>
      </el-form-item>
      <el-form-item label="结束时间">
        <el-date-picker
          clearable size="small" style="width: 200px"
          v-model="queryParams.endDate2"
          type="date"
          value-format="yyyy-MM-dd"
          placeholder="选择结束时间">
        </el-date-picker>
      </el-form-item>

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
    </el-row>

    <el-row :gutter="10" class="mb8"></el-row>
    <el-table v-loading="loading" :data="bridgeOrderList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="序号" align="center" prop="id" />
<!--      <el-table-column label="转出币种" align="center" prop="sourceCoinAddress" />-->
      <el-table-column label="转出币种" align="center" prop="sourceCoinAddress" :formatter="sourceCoinAddressFormat" width="150">
        <template slot-scope="scope">
          {{ sourceCoinAddressFormat(scope.row) }}
        </template>
      </el-table-column>
<!--      <el-table-column label="转入币种" align="center" prop="targetCoinAddress" />-->
      <el-table-column label="转入币种" align="center" prop="targetCoinAddress" :formatter="targetCoinAddressFormat" width="150">
        <template slot-scope="scope">
          {{ targetCoinAddressFormat(scope.row) }}
        </template>
      </el-table-column>
<!--      <el-table-column label="转出链" align="center" prop="sourceChainName" />-->
      <el-table-column label="转出链" align="center" prop="sourceChainId" :formatter="sourceChainIdFormat" width="150">
        <template slot-scope="scope">
          {{ sourceChainIdFormat(scope.row) }}
        </template>
      </el-table-column>
<!--targetChainIdFormat  sourceCoinAddressFormat -->
<!--      <el-table-column label="转入链" align="center" prop="targetChainName" />-->
      <el-table-column label="转入链" align="center" prop="targetChainId" :formatter="targetChainIdFormat" width="150">
        <template slot-scope="scope">
          {{ targetChainIdFormat(scope.row) }}
        </template>
      </el-table-column>

      <el-table-column label="数量" align="center" prop="amount">
      <template slot-scope="scope">
        <span>{{ formatNumber(scope.row.amount) }}</span>
      </template>
      </el-table-column>
      <el-table-column label="手续费" align="center" prop="fee" >
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.fee) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="gas费" align="center" prop="gasFee" />
<!--        <template slot-scope="scope">-->
<!--          <span>{{ formatNumber(scope.row.gasFee) }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->

      <!-- 转入钱包地址 -->
      <el-table-column label="转入钱包地址" align="center">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.targetAddress" placement="top">
            <span class="copyable" @click="copyToClipboard(scope.row.targetAddress)">
              {{ formatAddress(scope.row.targetAddress) }}
            </span>
          </el-tooltip>
        </template>
      </el-table-column>

      <!-- 转出钱包地址 -->
      <el-table-column label="转出钱包地址" align="center">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.sourceAddress" placement="top">
            <span class="copyable" @click="copyToClipboard(scope.row.sourceAddress)">
              {{ formatAddress(scope.row.sourceAddress) }}
            </span>
          </el-tooltip>
        </template>
      </el-table-column>

      <!-- 转出TXID -->
      <el-table-column label="转出TXID" align="center">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.outHash" placement="top">
            <span class="copyable" @click="copyToClipboard(scope.row.outHash)">
              {{ formatTxId(scope.row.outHash) }}
            </span>
          </el-tooltip>
        </template>
      </el-table-column>

      <!-- 转入TXID -->
      <el-table-column label="转入TXID" align="center">
        <template slot-scope="scope">
          <el-tooltip class="item" effect="dark" :content="scope.row.inHash" placement="top">
            <span class="copyable" @click="copyToClipboard(scope.row.inHash)">
              {{ formatTxId(scope.row.inHash) }}
            </span>
          </el-tooltip>
        </template>
      </el-table-column>

<!--      <el-table-column label="转入钱包地址" align="center" prop="targetAddress" />-->
<!--      <el-table-column label="转出钱包地址" align="center" prop="sourceAddress" />-->
<!--      <el-table-column label="转出TXID" align="center" prop="outHash" />-->
<!--      <el-table-column label="转入TXID" align="center" prop="inHash" />-->
<!--      <el-table-column label="跨链订单id" align="center" prop="orderId" />-->

      <el-table-column label="状态" align="center"  prop="status" :formatter="statusFormat" />
      <el-table-column label="发起时间" align="center" prop="createAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createAt, '{y}-{m}-{d} {h}:{i}:{s}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="结束时间" align="center" prop="updateAt" width="180">
        <template slot-scope="scope">
        <span v-if="scope.row.status === 5">
          {{ parseTime(scope.row.updateAt, '{y}-{m}-{d} {h}:{i}:{s}') }}
        </span>
          <span v-else>
          未完成..
        </span>
        </template>
      </el-table-column>

      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <div v-if="scope.row.status === 1">
            <el-button
              size="mini"
              type="text"
              icon="el-icon-edit"
              @click="handleUpdate(scope.row)"
              v-hasPermi="['admin/bridgeOrder/edit']"
            >人工审核</el-button>
          </div>
<!--          <div v-else-if="scope.row.status > 1">操作已完成</div>-->
          <div v-else> --- </div>
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
    <!-- 修改跨链记录状态对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">

      <el-form-item label="人工审核状态" label-width="150px"  prop="status">
          <el-select v-model="form.status" placeholder="请审核">
              <el-option
                  v-for="dict in manualReviewStatus"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
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
    listBridgeOrder,
    getBridgeOrder,
    delBridgeOrder,
    addBridgeOrder,
    updateBridgeOrder,
    listChain,
    changeBridgeOrderStatus,
    exportBridgeOrder,
} from "@/api/admin/bridgeOrder";
import ClipboardJS from 'clipboard';
import {exportCoin} from "@/api/admin/coin";
import { BigNumber } from 'bignumber.js';
import {listCoin} from "@/api/admin/bridgeConfig";
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
      sourceCoinAddressOptions: [],
      // statusOptions字典数据
      statusOptions: [],
      // manualReviewStatus
      manualReviewStatus:[],
      sourceCoinAddress:undefined,
      coinAddressMap: {

      },
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
        outHash: undefined,
        inHash:undefined,
        status: undefined,
        startDate1:undefined,
        startDate2:undefined,
        endDate1:undefined,
        endDate2:undefined,
      },
      // 表单参数
      form: {
        reviewStatus: null
      },
      // 表单校验
      rules: {
        status : [
          { required: true, message: '请选择审核状态', trigger: 'change' }
        ],
        reviewStatus: [
          { required: true, message: '请选择人工审核状态', trigger: 'change' }
        ]
      }
    };
  },
  created() {
    this.getChainItems()
    this.getCoinItems()
    // this.getChainItems()
    this.getDicts("bridge_order_status").then(response => {
      this.statusOptions = response.data.values||[];
    });

    this.getDicts("manual_review_status").then(response => {
      this.manualReviewStatus = response.data.values||[];
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

    getCoinItems() {
      this.getItems(listCoin, { pageSize: 10000 }).then(res => {
        console.log('Coin items response:', res);
        if (res && res.data && res.data.list) {
          this.sourceCoinAddressOptions = res.data.list.map(item => ({
            key: item.address,
            value: item.symbol,
            chainId: item.chainId // 确保每个币种有chainId字段
          }));

          // 更新 coinAddressMap
          this.coinAddressMap = res.data.list.reduce((map, item) => {
            map[`${item.address}-${item.chainId}`] = item.symbol;
            return map;
          }, {});
        } else {
          console.error('Invalid coin items response:', res);
        }
      }).catch(error => {
        console.error('Error fetching coin items:', error);
      });
    },

    formatNumber(number) {
      if (typeof number === 'number') {
        let bigNumber = new BigNumber(number);
        return bigNumber.toFixed().replace(/(\.\d*?[1-9])0+$/g, '$1').replace(/\.0+$/, '');
      }
      return number;
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
      // return this.selectDictLabel(this.statusOptions, row.status);
      if (row.status <=1) {
        return "待审核"
      }
      if (1 < row.status && row.status  <= 4) {
        return "进行中"
      }

      if (row.status === 5) {
        return "已完成"
      }


      if (row.status > 5) {
        return "跨链失败"
      }

    },

    sourceCoinAddressFormat(row, column) {
      return this.selectItemsLabel2(this.sourceCoinAddressOptions, row.sourceCoinAddress, row.sourceChainId);
    },

    targetCoinAddressFormat(row, column) {
      return this.selectItemsLabel2(this.sourceCoinAddressOptions, row.targetCoinAddress, row.targetChainId);
    },

    selectItemsLabel2(datas, address, chainId) {
      for (let item of datas) {
        if (item.key === String(address) && item.chainId === String(chainId)) {
          return item.value;
        }
      }
      return null; // 或者返回一个默认值
    },

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
        outHash: undefined,
        inHash:undefined,
        orderId: undefined,
        amount: undefined,
        status: undefined,
        reviewStatus: null,
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
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getBridgeOrder(id).then(response => {
        let data = response.data;
        data.sourceChainId = ''+data.sourceChainId
        data.targetChainId = ''+data.targetChainId
        data.reviewStatus = data.reviewStatus || null;
        // data.status = ''+data.status
        data.status = ''
        this.form = data;
        this.open = true;
        this.title = "人工审核跨链记录";
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
    },

    formatAddress(address) {
      return address.slice(0, 6) + '...' + address.slice(-3);
    },
    formatTxId(txId) {
      return txId.slice(0, 6) + '...' + txId.slice(-3);
    },
    copyToClipboard(text) {
      const clipboard = new ClipboardJS('.copyable', {
        text: () => text
      });
      clipboard.on('success', (e) => {
        this.$message({
          message: '复制成功',
          type: 'success'
        });
        clipboard.destroy();
      });
      clipboard.on('error', (e) => {
        this.$message.error('复制失败');
        clipboard.destroy();
      });
    },

  }
};
</script>
