import request from '@/utils/request'
// 查询跨链记录列表
export function listBridgeOrder(query) {
  return request({
    url: '/admin/bridgeOrder/list',
    method: 'get',
    params: query
  })
}
// 查询跨链记录详细
export function getBridgeOrder(id) {
  return request({
    url: '/admin/bridgeOrder/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增跨链记录
export function addBridgeOrder(data) {
  return request({
    url: '/admin/bridgeOrder/add',
    method: 'post',
    data: data
  })
}
// 修改跨链记录
export function updateBridgeOrder(data) {
  return request({
    url: '/admin/bridgeOrder/edit',
    method: 'put',
    data: data
  })
}
// 删除跨链记录
export function delBridgeOrder(ids) {
  return request({
    url: '/admin/bridgeOrder/delete',
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

// 跨链记录跨链记录状态修改
export function changeBridgeOrderStatus(id,status) {
  const data = {
    id,
    status
  }
  return request({
    url: '/admin/bridgeOrder/changeStatus',
    method: 'put',
    data:data
  })
}
