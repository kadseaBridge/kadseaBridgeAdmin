import request from '@/utils/request'
// 查询支付详情列表
export function listPayDetail(query) {
  return request({
    url: '/admin/payDetail/list',
    method: 'get',
    params: query
  })
}
// 查询支付详情详细
export function getPayDetail(id) {
  return request({
    url: '/admin/payDetail/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增支付详情
export function addPayDetail(data) {
  return request({
    url: '/admin/payDetail/add',
    method: 'post',
    data: data
  })
}
// 修改支付详情
export function updatePayDetail(data) {
  return request({
    url: '/admin/payDetail/edit',
    method: 'put',
    data: data
  })
}
// 删除支付详情
export function delPayDetail(ids) {
  return request({
    url: '/admin/payDetail/delete',
    method: 'delete',
    data:{
       ids:ids
    }
  })
}
// 关联chain表选项
export function listChain(query){
   return request({
     url: '/admin/chain/list',
     method: 'get',
     params: query
   })
}
// 关联coin表选项
export function listCoin(query){
   return request({
     url: '/admin/coin/list',
     method: 'get',
     params: query
   })
}
