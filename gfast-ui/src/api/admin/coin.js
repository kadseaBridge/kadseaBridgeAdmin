import request from '@/utils/request'
// 查询币种管理列表
export function listCoin(query) {
  return request({
    url: '/admin/coin/list',
    method: 'get',
    params: query
  })
}
// 查询币种管理详细
export function getCoin(id) {
  return request({
    url: '/admin/coin/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增币种管理
export function addCoin(data) {
  return request({
    url: '/admin/coin/add',
    method: 'post',
    data: data
  })
}
// 修改币种管理
export function updateCoin(data) {
  return request({
    url: '/admin/coin/edit',
    method: 'put',
    data: data
  })
}
// 删除币种管理
export function delCoin(ids) {
  return request({
    url: '/admin/coin/delete',
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
