import request from '@/utils/request'
// 查询链列表
export function listChain(query) {
  return request({
    url: '/admin/chain/list',
    method: 'get',
    params: query
  })
}
// 查询链详细
export function getChain(id) {
  return request({
    url: '/admin/chain/get',
    method: 'get',
    params: {
     id: id.toString()
    }
  })
}
// 新增链
export function addChain(data) {
  return request({
    url: '/admin/chain/add',
    method: 'post',
    data: data
  })
}
// 修改链
export function updateChain(data) {
  return request({
    url: '/admin/chain/edit',
    method: 'put',
    data: data
  })
}
// 删除链
export function delChain(ids) {
  return request({
    url: '/admin/chain/delete',
    method: 'delete',
    data:{
       ids:ids
    }
  })
}
