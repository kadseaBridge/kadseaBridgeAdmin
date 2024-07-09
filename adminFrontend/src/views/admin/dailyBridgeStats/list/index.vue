<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">

      <el-form-item label="币种" prop="coinAddress">
        <el-input
            v-model="queryParams.coinName"
            placeholder="请输入币种"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
      <el-form-item label="链类型" prop="chainId">
        <el-input
            v-model="queryParams.chainName"
            placeholder="请输入链类型"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>

      <el-form-item label="日期" prop="date">
<!--        <el-date-picker-->
<!--          clearable size="small" style="width: 200px"-->
<!--          v-model="queryParams.date"-->
<!--          type="date"-->
<!--          value-format="yyyy-MM-dd"-->
<!--          placeholder="选择时间">-->
<!--        </el-date-picker>-->

        <el-date-picker
          clearable size="small" style="width: 300px"
          v-model="queryParams.date"
          type="daterange"
          range-separator="至"
          start-placeholder="开始日期"
          end-placeholder="结束日期"
          value-format="yyyy-MM-dd"
          placeholder="选择日期范围">
        </el-date-picker>
      </el-form-item>

      <el-form-item>
        <el-button type="primary" icon="el-icon-search" size="mini" @click="handleQuery">搜索</el-button>
        <el-button icon="el-icon-refresh" size="mini" @click="resetQuery">重置</el-button>
      </el-form-item>
    </el-form>
    <el-table v-loading="loading" :data="dailyBridgeStatsList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="id" align="center" prop="id" />
      <el-table-column label="日期" align="center" prop="date" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.date, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="币种" align="center" prop="coinAddress" :formatter="coinAddressFormat" width="100">
        <template slot-scope="scope">
          {{ coinAddressFormat(scope.row) }}
        </template>

      </el-table-column>
      <el-table-column label="链类型" align="center" prop="chainId" :formatter="chainIdFormat" width="100">
        <template slot-scope="scope">
          {{ chainIdFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="跨链转入" align="center" prop="transferIn">
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.transferIn) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="跨链转出" align="center" prop="transferOut" >
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.transferOut) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="跨链差额" align="center" prop="transferDifference">
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.transferDifference) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="跨链手续费" align="center" prop="fee" >
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.fee) }}</span>
        </template>
      </el-table-column>
      <el-table-column label="平台资产" align="center" prop="platformAssets" >
        <template slot-scope="scope">
          <span>{{ formatNumber(scope.row.platformAssets) }}</span>
        </template>
      </el-table-column>
<!--      <el-table-column label="财务地址余额" align="center" prop="financialBalance" />-->
    </el-table>
    <pagination
      v-show="total>0"
      :total="total"
      :page.sync="queryParams.pageNum"
      :limit.sync="queryParams.pageSize"
      @pagination="getList"
    />
  </div>
</template>
<script>
import {
    listDailyBridgeStats,
    getDailyBridgeStats,
    delDailyBridgeStats,
    addDailyBridgeStats,
    updateDailyBridgeStats,
    listCoin,
    listChain,
} from "@/api/admin/dailyBridgeStats";
import {BigNumber} from "bignumber.js";
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
      // coinAddressOptions关联表数据
      coinAddressOptions: [],
      // chainIdOptions关联表数据
      chainIdOptions: [],
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        // date: undefined,
        date: [],
        startDate:undefined,
        endDate:undefined,
        // coinAddress: undefined,
        // chainId: undefined,

        coinName: undefined,
        chainName: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        date : [
          { required: true, message: "时间不能为空", trigger: "blur" }
        ],
        coinAddress : [
          { required: true, message: "币种不能为空", trigger: "blur" }
        ],
        chainId : [
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
          { required: true, message: "财务地址余额不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getCoinItems()
    this.getChainItems()
    this.getList();
  },
  methods: {

    getCoinItems() {
      this.getItems(listCoin, { pageSize: 10000 }).then(res => {
        console.log('Coin items response:', res);
        if (res && res.data && res.data.list) {
          this.coinAddressOptions = res.data.list.map(item => ({
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

    //关联chain表选项
    getChainItems() {
      this.getItems(listChain, {pageSize:10000}).then(res => {
        this.chainIdOptions = this.setItems(res, 'chainId', 'name')
      })
    },
    /** 查询跨链日统计列表 */
    getList() {
      this.loading = true;
      listDailyBridgeStats(this.queryParams).then(response => {
        this.dailyBridgeStatsList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 币种关联表翻译
    coinAddressFormat(row, column) {
      return this.selectItemsLabel1(this.coinAddressOptions, row.coinAddress, row.chainId);
    },

    selectItemsLabel1(datas, address, chainId) {
      for (let item of datas) {
        if (item.key === String(address) && item.chainId === String(chainId)) {
          return item.value;
        }
      }
      return null; // 或者返回一个默认值
    },

    // 链类型关联表翻译
    chainIdFormat(row, column) {
      return this.selectItemsLabel(this.chainIdOptions, row.chainId);
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
        coinAddress: undefined,
        chainId: undefined,
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
      const { date } = this.queryParams;
      const [startDate1, endDate1] = date;

      this.queryParams.startDate = startDate1;
      this.queryParams.endDate = endDate1;
      console.log('查询参数:',  this.queryParams.startDate, this.queryParams.endDate );
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

  }
};
</script>
