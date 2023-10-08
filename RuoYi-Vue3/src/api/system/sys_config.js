import request from '@/utils/request'


// 获取参数配置表列表
export function listSysConfig(query) {
  return request({
    url: '/system/sys/config/list',
    method: 'get',
    params: query

  })
}

// 添加参数配置表
export function addSysConfig(query) {
  return request({
        url: '/system/sys/config',
    method: 'post',
    data: query

  })
}

// 修改参数配置表
export function updateSysConfig(query) {
  return request({
        url: '/system/sys/config',
    method: 'put',
    data: query

  })
}

// 删除参数配置表
export function deleteSysConfig(query) {
  return request({
        url: '/system/sys/config/' + query,
    method: 'delete'

  })
}

// 获取参数配置表
export function getSysConfig(query) {
  return request({
        url: '/system/sys/config/' + query,
    method: 'get'

  })
}

