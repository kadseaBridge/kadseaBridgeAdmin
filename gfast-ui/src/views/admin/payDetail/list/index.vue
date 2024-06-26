<template>
  <div class="app-container">
    <el-form :model="queryParams" ref="queryForm" :inline="true" label-width="68px">    
      <el-form-item label="" prop="orderId">
        <el-input
            v-model="queryParams.orderId"
            placeholder="请输入"
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
    <el-row :gutter="10" class="mb8">
      <el-col :span="1.5">
        <el-button
          type="primary"
          icon="el-icon-plus"
          size="mini"
          @click="handleAdd"
          v-hasPermi="['admin/payDetail/add']"
        >新增</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="success"
          icon="el-icon-edit"
          size="mini"
          :disabled="single"
          @click="handleUpdate"
          v-hasPermi="['admin/payDetail/edit']"
        >修改</el-button>
      </el-col>
      <el-col :span="1.5">
        <el-button
          type="danger"
          icon="el-icon-delete"
          size="mini"
          :disabled="multiple"
          @click="handleDelete"
          v-hasPermi="['admin/payDetail/delete']"
        >删除</el-button>
      </el-col>
    </el-row>
    <el-table v-loading="loading" :data="payDetailList" @selection-change="handleSelectionChange">
      <el-table-column type="selection" width="55" align="center" />      
      <el-table-column label="" align="center" prop="id" />      
      <el-table-column label="" align="center" prop="orderId" />      
      <el-table-column label="" align="center" prop="payChainId" :formatter="payChainIdFormat" width="100">
        <template slot-scope="scope">
          {{ payChainIdFormat(scope.row) }}
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="payCoinAddress" :formatter="payCoinAddressFormat" width="100">
        <template slot-scope="scope">
          {{ payCoinAddressFormat(scope.row) }}
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="transactionHash" />      
      <el-table-column label="" align="center" prop="payAmount" />      
      <el-table-column label="" align="center" prop="payTime" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.payTime, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="payAddress" />      
      <el-table-column label="" align="center" prop="createAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.createAt, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="updateAt" width="180">
        <template slot-scope="scope">
            <span>{{ parseTime(scope.row.updateAt, '{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>      
      <el-table-column label="" align="center" prop="blockNumber" />      
      <el-table-column label="" align="center" prop="payGasFee" />      
      <el-table-column label="操作" align="center" class-name="small-padding fixed-width">
        <template slot-scope="scope">
          <el-button
            size="mini"
            type="text"
            icon="el-icon-edit"
            @click="handleUpdate(scope.row)"
            v-hasPermi="['admin/payDetail/edit']"
          >修改</el-button>
          <el-button
            size="mini"
            type="text"
            icon="el-icon-delete"
            @click="handleDelete(scope.row)"
            v-hasPermi="['admin/payDetail/delete']"
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
    <!-- 添加或修改支付详情对话框 -->
    <el-dialog :title="title" :visible.sync="open" width="800px" append-to-body :close-on-click-modal="false">
      <el-form ref="form" :model="form" :rules="rules" label-width="80px">      
      <el-form-item label="" prop="orderId">
           <el-input v-model="form.orderId" placeholder="请输入" />
      </el-form-item>      
      <el-form-item label="" prop="payChainId">
          <el-select v-model="form.payChainId" placeholder="请选择">
              <el-option
                  v-for="item in payChainIdOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
      </el-form-item>      
      <el-form-item label="" prop="payCoinAddress">
          <el-select v-model="form.payCoinAddress" placeholder="请选择">
              <el-option
                  v-for="item in payCoinAddressOptions"
                  :key="item.key"
                  :label="item.value"
                  :value="item.key"
              ></el-option>
          </el-select>
      </el-form-item>      
      <el-form-item label="" prop="transactionHash">
           <el-input v-model="form.transactionHash" placeholder="请输入" />
      </el-form-item>      
      <el-form-item label="" prop="payAmount">
           <el-input v-model="form.payAmount" placeholder="请输入" />
      </el-form-item>      
       <el-form-item label="" prop="payTime">
           <el-date-picker clearable size="small" style="width: 200px"
               v-model="form.payTime"
               type="date"
               value-format="yyyy-MM-dd"
               placeholder="选择">
           </el-date-picker>
       </el-form-item>      
      <el-form-item label="" prop="payAddress">
           <el-input v-model="form.payAddress" placeholder="请输入" />
      </el-form-item>      
       <el-form-item label="" prop="createAt">
           <el-date-picker clearable size="small" style="width: 200px"
               v-model="form.createAt"
               type="date"
               value-format="yyyy-MM-dd"
               placeholder="选择">
           </el-date-picker>
       </el-form-item>      
       <el-form-item label="" prop="updateAt">
           <el-date-picker clearable size="small" style="width: 200px"
               v-model="form.updateAt"
               type="date"
               value-format="yyyy-MM-dd"
               placeholder="选择">
           </el-date-picker>
       </el-form-item>      
      <el-form-item label="" prop="remark">
           <el-input v-model="form.remark" placeholder="请输入" />
      </el-form-item>      
      <el-form-item label="" prop="blockNumber">
           <el-input v-model="form.blockNumber" placeholder="请输入" />
      </el-form-item>      
      <el-form-item label="" prop="payGasFee">
           <el-input v-model="form.payGasFee" placeholder="请输入" />
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
    listPayDetail,
    getPayDetail,
    delPayDetail,
    addPayDetail,
    updatePayDetail,    
    listChain,    
    listCoin,    
} from "@/api/admin/payDetail";
export default {
  components:{},
  name: "PayDetail",
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
      // 支付详情表格数据
      payDetailList: [],
      // 弹出层标题
      title: "",
      // 是否显示弹出层
      open: false,      
      // payChainIdOptions关联表数据
      payChainIdOptions: [],      
      // payCoinAddressOptions关联表数据
      payCoinAddressOptions: [],      
      // 查询参数
      queryParams: {
        pageNum: 1,
        pageSize: 10,
        orderId: undefined,
      },
      // 表单参数
      form: {},
      // 表单校验
      rules: { 
        orderId : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payChainId : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payCoinAddress : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        transactionHash : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payAmount : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payTime : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payAddress : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
        payGasFee : [
          { required: true, message: "不能为空", trigger: "blur" }
        ],
      }
    };
  },
  created() {    
    this.getChainItems()    
    this.getCoinItems()
    this.getList();
  },
  methods: {    
    //关联chain表选项
    getChainItems() {
      this.getItems(listChain, {pageSize:10000}).then(res => {
        this.payChainIdOptions = this.setItems(res, 'chainId', 'name')
      })
    },    
    //关联coin表选项
    getCoinItems() {
      this.getItems(listCoin, {pageSize:10000}).then(res => {
        this.payCoinAddressOptions = this.setItems(res, 'address', 'name')
      })
    },    
    /** 查询支付详情列表 */
    getList() {
      this.loading = true;
      listPayDetail(this.queryParams).then(response => {
        this.payDetailList = response.data.list;
        this.total = response.data.total;
        this.loading = false;
      });
    },    
    // 关联表翻译
    payChainIdFormat(row, column) {
      return this.selectItemsLabel(this.payChainIdOptions, row.payChainId);
    },    
    // 关联表翻译
    payCoinAddressFormat(row, column) {
      return this.selectItemsLabel(this.payCoinAddressOptions, row.payCoinAddress);
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
        orderId: undefined,        
        payChainId: undefined,        
        payCoinAddress: undefined,        
        transactionHash: undefined,        
        payAmount: undefined,        
        payTime: undefined,        
        payAddress: undefined,        
        createAt: undefined,        
        updateAt: undefined,        
        remark: undefined,        
        blockNumber: undefined,        
        payGasFee: undefined,        
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
      this.title = "添加支付详情";
    },
    /** 修改按钮操作 */
    handleUpdate(row) {
      this.reset();
      const id = row.id || this.ids
      getPayDetail(id).then(response => {
        let data = response.data;        
        data.payChainId = ''+data.payChainId        
        data.payCoinAddress = ''+data.payCoinAddress        
        this.form = data;
        this.open = true;
        this.title = "修改支付详情";
      });
    },
    /** 提交按钮 */
    submitForm: function() {
      this.$refs["form"].validate(valid => {
        if (valid) {
          if (this.form.id != undefined) {
            updatePayDetail(this.form).then(response => {
              if (response.code === 0) {
                this.msgSuccess("修改成功");
                this.open = false;
                this.getList();
              } else {
                this.msgError(response.msg);
              }
            });
          } else {
            addPayDetail(this.form).then(response => {
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
      this.$confirm('是否确认删除支付详情编号为"' + ids + '"的数据项?', "警告", {
          confirmButtonText: "确定",
          cancelButtonText: "取消",
          type: "warning"
        }).then(function() {
          return delPayDetail(ids);
        }).then(() => {
          this.getList();
          this.msgSuccess("删除成功");
        }).catch(function() {});
    }
  }
};
</script>