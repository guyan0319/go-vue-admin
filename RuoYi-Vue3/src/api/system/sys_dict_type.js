import request from '@/utils/request'


// 获取字典类型表列表
export function listSysDictType(query) {
  return request({
    url: '/system/dict/type/list',
    method: 'get',
    params: query

  })
}

// 添加字典类型表
export function addSysDictType(query) {
  return request({
        url: '/system/dict/type',
    method: 'post',
    data: query

  })
}

// 修改字典类型表
export function updateSysDictType(query) {
  return request({
        url: '/system/dict/type',
    method: 'put',
    data: query

  })
}

// 删除字典类型表
export function deleteSysDictType(query) {
  return request({
        url: '/system/dict/type/' + query,
    method: 'delete'

  })
}

// 获取字典类型表
export function getSysDictType(query) {
  return request({
        url: '/system/dict/type/' + query,
    method: 'get'

  })
}

