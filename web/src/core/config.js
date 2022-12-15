/*
 * @Author: your name
 * @Date: 2022-02-23 14:01:47
 * @LastEditTime: 2022-05-17 09:21:02
 * @LastEditors: yangpanda 573516293@qq.com
 * @Description: 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 * @FilePath: /web/src/core/config.js
 */
/**
 * 网站配置文件
 */

const config = {
  appName: '呼唤邻居',
  appLogo: 'https://huhuanlinju.vvv5g.com/api/uploads/file/876448b30e5a74e7d3109cb8b136a8f9_20220413103412.jpeg',
  showViteLogo: true
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require('chalk')
    console.log(
      chalk.green(
        `> 欢迎使用Gin-Vue-Admin，开源地址：https://github.com/flipped-aurora/gin-vue-admin`
      )
    )
    console.log(
      chalk.green(
        `> 当前版本:V2.5.0 beta.2`
      )
    )
    console.log(
      chalk.green(
        `> 加群方式:微信：shouzi_1994 QQ群：622360840`
      )
    )
    console.log(
      chalk.green(
        `> GVA讨论社区：https://support.qq.com/products/371961`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认自动化文档地址:http://127.0.0.1:${env.VITE_SERVER_PORT}/swagger/index.html`
      )
    )
    console.log(
      chalk.green(
        `> 默认前端文件运行地址:http://127.0.0.1:${env.VITE_CLI_PORT}`
      )
    )
    console.log(
      chalk.green(
        `> 如果项目让您获得了收益，希望您能请团队喝杯可乐:https://www.github.com/flipped-aurora/gin-vue-admin/server.com/docs/coffee`
      )
    )
    console.log('\n')
  }
}

export default config
