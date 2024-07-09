import request from '@/utils/request'
// 查询跨链日统计列表
export function listDailyBridgeStats(query) {
  return request({
    url: '/admin/dailyBridgeStats/list',
    method: 'get',
    params: query
  })
}
// 查询跨链日统计详细
export function getDailyBridgeStats(id) {
  return request({
    url: '/admin/dailyBridgeStats/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增跨链日统计
export function addDailyBridgeStats(data) {
  return request({
    url: '/admin/dailyBridgeStats/add',
    method: 'post',
    data: data
  })
}
// 修改跨链日统计
export function updateDailyBridgeStats(data) {
  return request({
    url: '/admin/dailyBridgeStats/edit',
    method: 'put',
    data: data
  })
}
// 删除跨链日统计
export function delDailyBridgeStats(ids) {
  return request({
    url: '/admin/dailyBridgeStats/delete',
    method: 'delete',
    data:{
       ids:ids
    }
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
// 关联chain表选项
export function listChain(query){
   return request({
     url: '/admin/chain/list',
     method: 'get',
     params: query
   })
}
