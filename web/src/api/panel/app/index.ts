import { http } from '@/utils'

export default {
  // 获取应用列表
  list: (page: number, limit: number): any => http.Get('/app/list', { params: { page, limit } }),
  // 安装应用
  install: (slug: string, channel: string | null): any =>
    http.Post('/app/install', { slug, channel }),
  // 卸载应用
  uninstall: (slug: string): any => http.Post('/app/uninstall', { slug }),
  // 更新应用
  update: (slug: string): any => http.Post('/app/update', { slug }),
  // 设置首页显示
  updateShow: (slug: string, show: boolean): any => http.Post('/app/update_show', { slug, show }),
  // 应用是否已安装
  isInstalled: (slug: string): any => http.Get('/app/is_installed', { params: { slug } }),
  // 更新缓存
  updateCache: (): any => http.Get('/app/update_cache')
}
