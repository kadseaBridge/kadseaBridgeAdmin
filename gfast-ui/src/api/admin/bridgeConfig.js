import request from '@/utils/request'
// 查询可跨币对列表
export function listBridgeConfig(query) {
  return request({
    url: '/admin/bridgeConfig/list',
    method: 'get',
    params: query
  })
}
// 查询可跨币对详细
export function getBridgeConfig(id) {
  return request({
    url: '/admin/bridgeConfig/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增可跨币对
export function addBridgeConfig(data) {
  return request({
    url: '/admin/bridgeConfig/add',
    method: 'post',
    data: data
  })
}
// 修改可跨币对
export function updateBridgeConfig(data) {
  return request({
    url: '/admin/bridgeConfig/edit',
    method: 'put',
    data: data
  })
}
// 删除可跨币对
export function delBridgeConfig(ids) {
  return request({
    url: '/admin/bridgeConfig/delete',
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
