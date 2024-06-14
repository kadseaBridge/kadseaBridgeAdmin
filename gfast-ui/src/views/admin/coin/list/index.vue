<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">
      <el-form-item label="代币简称" prop="symbol">
        <el-input
            v-model="queryParams.symbol"
            placeholder="请输入代币简称"
            clearable
            size="small"
            @keyup.enter.native="handleQuery"
        />
      </el-form-item>
        <el-form-item label="链" prop="chainId">
          <el-select v-model="queryParams.chainId" placeholder="请选择链" clearable size="small">
              <el-option
                  v-for="item in chainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              />
          </el-select>
        </el-form-item>
      <el-form-item label="代币地址" prop="address">
        <el-input
            v-model="queryParams.address"
            placeholder="请输入代币地址"
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
          v-hasPermi="['admin/coin/add']"
        >新增</el-button>
      </el-col>
<!--      <el-col :span="1.5">-->
<!--        <el-button-->
<!--          type="success"-->
<!--          icon="el-icon-edit"-->
<!--          size="mini"-->
<!--          :disabled="single"-->
<!--          @click="handleUpdate"-->
<!--          v-hasPermi="['admin/coin/edit']"-->
<!--        >修改</el-button>-->
<!--      </el-col>-->
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
    <el-table v-loading="loading" :data="coinList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />
      <el-table-column label="序号" align="center" prop="id" />
      <el-table-column label="币种" align="center" prop="symbol" />
      <el-table-column label="链" align="center" prop="chainId" :formatter="chainIdFormat" width="100">
        <template slot-scope="scope">
          {{ chainIdFormat(scope.row) }}
        </template>
      </el-table-column>
      <el-table-column label="精度" align="center" prop="decimals" />
      <el-table-column label="合约地址" align="center" prop="address" />

      <el-table-column label="图标" align="center" prop="icon" />
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['admin/coin/edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['admin/coin/delete']"
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
    <!-- 添加或修改币种管理对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">
      <el-form-item label="代币名称" prop="name">
           <el-input v-model="form.name" placeholder="请输入代币名称" />
      </el-form-item>
      <el-form-item label="代币简称" prop="symbol">
           <el-input v-model="form.symbol" placeholder="请输入代币简称" />
      </el-form-item>
      <el-form-item label="币种精度" prop="decimals" >
        <el-input v-model="form.decimals" placeholder="请输入币种精度" />
      </el-form-item>
      <el-form-item label="链" prop="chainId">
          <el-select v-model="form.chainId" placeholder="请选择链">
              <el-option
                  v-for="item in chainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
      </el-form-item>
      <el-form-item label="代币地址" prop="address">
           <el-input v-model="form.address" placeholder="请输入代币地址,如为空则表示公链原生币" />
      </el-form-item>

       <el-form-item label="是否上架" prop="isEnable">
           <el-radio-group v-model="form.isEnable">
               <el-radio
                v-for="dict in isEnableOptions"
                :key="dict.key"
                :label="dict.key"
               >{{dict.value }}</el-radio>
           </el-radio-group>
       </el-form-item>
      <el-form-item label="币种类型" prop="tokenType">
          <el-select v-model="form.tokenType" placeholder="请选择币种类型">
              <el-option
                  v-for="dict in tokenTypeOptions"
                  :key="dict.key"
                  :label="dict.value"
                      :value="dict.key"
              ></el-option>
          </el-select>
      </el-form-item>
      <el-form-item label="代币图标" prop="icon">
           <el-input v-model="form.icon" placeholder="请输入代币图标" />
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
    listCoin,
    getCoin,
    delCoin,
    addCoin,
    updateCoin,
    listChain,
    exportCoin
} from "@/api/admin/coin";
import {exportPost} from "@/api/system/post";
export default {
  components:{},
  name: "Coin",
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
      // 币种管理表格数据
      coinList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,
      // chainIdOptions关联表数据
      chainIdOptions: [],
      // isEnableOptions字典数据
      isEnableOptions: [],
      // tokenTypeOptions字典数据
      tokenTypeOptions: [],
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        symbol: undefined,
        chainId: undefined,
        address: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: {
        name : [
          { required: true, message: "代币名称不能为空", trigger: "blur" }
        ],
        symbol : [
          { required: true, message: "代币简称不能为空", trigger: "blur" }
        ],
        chainId : [
          { required: true, message: "链不能为空", trigger: "blur" }
        ],
        address : [
          {  message: "代币地址为空则为公链原生币", trigger: "blur" }
        ],
        isEnable : [
          { required: true, message: "是否上架不能为空", trigger: "blur" }
        ],
        tokenType : [
          { required: true, message: "币种类型不能为空", trigger: "blur" }
        ],
        decimals : [
          { required: true, message: "精度不能为空", trigger: "blur" }
        ],
        icon : [
          { required: false, message: "代币图标不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {
    this.getChainItems()
    this.getDicts("is_enable").then(response => {
      this.isEnableOptions = response.data.values||[];
    });
    this.getDicts("token_type").then(response => {
      this.tokenTypeOptions = response.data.values||[];
    });
    this.getList();
  },
  methods: {
    //关联chain表选项
    getChainItems() {
      this.getItems(listChain, {pageSize:10000}).then(res => {
        this.chainIdOptions = this.setItems(res, 'chainId', 'name')
      })
    },
    /** 查询币种管理列表 */
    getList() {
      this.loading = true;
      listCoin(this.queryParams).then(response => {
        this.coinList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },
    // 链Id关联表翻译
    chainIdFormat(row, column) {
      return this.selectItemsLabel(this.chainIdOptions, row.chainId);
    },
    // 是否上架字典翻译
    isEnableFormat(row, column) {
      return this.selectDictLabel(this.isEnableOptions, row.isEnable);
    },
    // 币种类型字典翻译
    tokenTypeFormat(row, column) {
      return this.selectDictLabel(this.tokenTypeOptions, row.tokenType);
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
        symbol: undefined,
        chainId: undefined,
        address: undefined,
        isEnable: "0" ,
        createAt: undefined,
        updateAt: undefined,
        tokenType: undefined,
        decimals: undefined,
        icon: undefined,
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
      this.title = "添加币种管理";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getCoin(id).then(response => {
        let data = response.data;
        data.chainId = ''+data.chainId
        data.isEnable = ''+data.isEnable
        data.tokenType = ''+data.tokenType
        this.form = data;
        this.open = true;
        this.title = "修改币种管理";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updateCoin(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addCoin(this.form).then(response => {
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
      // '是否确认删除币种管理编号为"' + ids + '"的数据项?', "警告",
      const ids = row.id || this.ids;
      this.$confirm('是否确认删除该币种?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delCoin(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    },

    /** 导出按钮操作 */
    handleExport() {
      const queryParams = this.queryParams;
      this.$confirm('是否确认导出所有币种?', "警告", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning"
      }).then(() => {
        return exportCoin(queryParams);
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
      link.download = 'coins.xlsx';
      link.click();
      window.URL.revokeObjectURL(link.href);
    }

  }
};
</script>
