import request from '@/utils/request'
// 查询可跨币种设置列表
export function listBridgeConfig(query) {
  return request({
    url: '/work/bridgeConfig/list',
    method: 'get',
    params: query
  })
}
// 查询可跨币种设置详细
export function getBridgeConfig(id) {
  return request({
    url: '/work/bridgeConfig/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增可跨币种设置
export function addBridgeConfig(data) {
  return request({
    url: '/work/bridgeConfig/add',
    method: 'post',
    data: data
  })
}
// 修改可跨币种设置
export function updateBridgeConfig(data) {
  return request({
    url: '/work/bridgeConfig/edit',
    method: 'put',
    data: data
  })
}
// 删除可跨币种设置
export function delBridgeConfig(ids) {
  return request({
    url: '/work/bridgeConfig/delete',
    method: 'delete',
    data:{
       ids:ids
    }
  })
}
