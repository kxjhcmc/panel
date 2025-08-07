import { http } from '@/utils'

export default {
  // 负载状态
  load: (): any => http.Get('/apps/postgresql/load'),
  // 获取配置
  config: (): any => http.Get('/apps/postgresql/config'),
  // 保存配置
  saveConfig: (config: string): any => http.Post('/apps/postgresql/config', { config }),
  // 获取用户配置
  userConfig: (): any => http.Get('/apps/postgresql/user_config'),
  // 保存配置
  saveUserConfig: (config: string): any => http.Post('/apps/postgresql/user_config', { config }),
  // 获取日志
  log: (): any => http.Get('/apps/postgresql/log'),
  // 清空错误日志
  clearLog: (): any => http.Post('/apps/postgresql/clear_log')
}
