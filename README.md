# 叮咚买菜助手

[![dingdong-release](https://github.com/gelove/dingdong/actions/workflows/go.yml/badge.svg)](https://github.com/gelove/dingdong/actions/workflows/go.yml)

## 下载可执行程序 (非码农请直接下载并运行, 不需要安装环境, 不需要下载源码！！！)

从 [奶牛快传](https://cowtransfer.com/s/e36fecd67ddc40
) 下载对应系统的最新版本

## 4月27日

程序启动后，可以在浏览器打开 <http://localhost:9999/notify> 测试推送通知是否正常

## 4月25日

**最好用一个账号抢菜, 另一个账号捡漏下单和监听运力**

## 抢菜

### 获取 Session

**如果无法找到所列出的请求，请参见后文 iOS 设备 Charles 抓包帮助**

1、在iOS设备上启动叮咚买菜APP

2、完成登录

3、启动Charles并完成抓包配置（需要配置SSL抓包）

4、点击“购物车”并刷新

5、在请求中找到https://maicai.api.ddxq.mobi/cart/index

6、右击该请求，选择Export Session，保存到项目 session 文件夹下，文件类型请选择JSON Session File (.chlsj)

### 获取 im_secret

**接续获取 Session 第三步**

4、点击“我的”并刷新

5、在请求中找到https://sunquan.api.ddxq.mobi/api/v1/user/detail

6、左击该请求，选择Contents选项卡，在下半部分选项卡中选择JSON Text视图

7、找到 user_info 下的 im_secret 字段，复制其值到配置文件中

### 通过抓包获取叮咚小程序的网络请求, 模仿接口调用, 配置文件中 header 和 form 里的必要参数需要根据自己抓包到的数据进行设置

### iphone

1. iphone上通过App Store安装Stream
2. 打开叮咚买菜微信小程序和Stream, Stream点击"开始抓包"
3. 回到叮咚买菜微信小程序, 商品加入购物车并结算提交(无需付款)使得小程序发起网络请求
4. 回到Stream停止抓包并查看抓包历史, 点击按域名, 会看到 maicai.api.ddxq.mobi 这个域名下的请求

感谢群友 邱建忠 提供的前端网页

感谢群友 hhhhg~ 提供的教程 
<a href="https://github.com/gelove/dingdong/blob/main/assets/doc/安卓抓包方法及Windows端小白教程.docx" target="_blank">安卓抓包方法及Windows端小白教程.docx</a>

[Stream 抓包教程(IOS)](https://www.jianshu.com/p/8a0fe2500f24)

[Charles 抓包教程(MacOS, Windows)](https://www.jianshu.com/p/ff85b3dac157)

![微信群](https://github.com/gelove/dingdong/blob/main/assets/image/wechat.jpeg)

QQ群: 60768433 (golang & rust & flutter & 区块链 交流学习)

### 配置文件 config.yml 设置

```js
base_concurrency // 除了提交订单的其他请求并发数 默认为1
submit_concurrency // 最后提交订单的并发数 默认为2
snap_up // 抢购 0关闭, 1 六点抢, 2 八点半抢, 3 六点和八点半都抢
advance_time // 抢购提前进入时间 单位:秒 默认为15
pick_up_needed // 闲时捡漏开关 false关闭 true打开 在抢购高峰期之外的时间捡漏 使用时需同时打开监听器
monitor_needed // 监听器开关 监听是否有运力
monitor_success_wait // 成功监听(发起捡漏或通知)之后的休息时间 单位:分钟 默认为10
monitor_interval_min // 监听器调用接口的最小时间间隔 单位:秒 默认为25 (防止接口调用过于频繁, 被叮咚风控)
monitor_interval_max // 监听器调用接口的最大时间间隔 单位:秒 默认为35
notify_needed // 通知开关 发现运力时或捡漏成功时的通知 使用需同时打开监听或捡漏
audio_needed // 播放音频开关 在订单提交成功后播放音频
```

#### 通过接口修改配置文件

GET 请求 localhost:9999/set

浏览器中打开 <http://localhost:9999/set?monitor_needed=1&monitor_interval_min=10&monitor_interval_max=20> , 并设置想要修改的参数

| 参数                    | 名称        |                说明                 |
|:----------------------|:----------|:---------------------------------:|
| base_concurrency      | 基础并发数     |       除了提交订单的其他请求并发数, 默认为1        |
| submit_concurrency    | 提交并发数     |         最后提交订单的并发数, 默认为2          |
| snap_up               | 抢购模式      |  0 关闭, 1 六点抢, 2 八点半抢, 3 六点和八点半都抢  |
| advance_time          | 抢购提前时间    | 提前一段时间执行一些预备任务(如获取购物车) 单位:秒 默认为15 |
| pick_up_needed        | 捡漏开关      |             0 关闭 1 打开             |
| monitor_needed        | 监听开关      |             0 关闭 1 打开             |
| monitor_success_wait  | 监听成功后休息时间 | 监听成功(发起捡漏或通知)之后的休息时间 单位:分钟 默认为10  |
| monitor_interval_min  | 监听的最小时间间隔 |       监听器调用接口的最小时间间隔 单位: 秒        |
| monitor_interval_max  | 监听的最大时间间隔 |       监听器调用接口的最大时间间隔 单位: 秒        |
| notify_needed         | 通知开关      |             0 关闭 1 打开             |
| audio_needed          | 播放音频开关    |             0 关闭 1 打开             |
| push_plus             | 通知用户      |      安卓和苹果都可使用 PushPlus 推送通知      |
| bark                  | 通知用户(仅苹果) |          苹果使用 Bark 推送通知           |

**例子**
| api | 说明 |
| :-----| :---- |
| localhost:9999/set?bark=xxx,yyy | 添加需要通知的用户, 第一个是自己的 barkID, 其他为需要通知到的朋友(只能通知与你同属一个叮咚发货站点的用户) |
| localhost:9999/set?snap_up=1 | 六点抢购 |
| localhost:9999/set?pick_up_needed=1 | 打开捡漏,在抢购高峰期之外的时间捡漏(需同时打开监听器) |
| localhost:9999/set?monitor_needed=1 | 打开监听器 在抢购高峰期之外的时间监听是否有运力 |
| localhost:9999/set?notify_needed=1 | 打开推送通知（需同时打开监听器） |
| localhost:9999/set?monitor_success_wait=5 | 设置推送时间间隔(防止太过频繁) |

<localhost:9999/set?bark=first,second&pick_up_needed=1&monitor_needed=1&notify_needed=1&monitor_success_wait=5>

## 运力监听

![effect](https://github.com/gelove/dingdong/blob/main/assets/image/effect.jpeg)

### 当有运力时, 发送通知到手机

### 苹果手机接收通知

#### 1.安装 bark 得到自己的 barkID (下图红框中), 并将其粘贴到配置文件的 "bark" 字段下

![bark](https://github.com/gelove/dingdong/blob/main/assets/image/user.jpeg)

#### 2.配置文件中 "bark" 字段为一组需要通知的 bark token

#### 3.bark 需打开允许通知

![notify](https://github.com/gelove/dingdong/blob/main/assets/image/notify.jpeg)

### 安卓或苹果手机通过微信接收通知

#### 微信上关注 "pushplus 推送加" 公众号

![PushPlus](https://github.com/gelove/dingdong/blob/main/assets/image/PushPlus.jpeg)

#### 点击 <kbd>激活消息</kbd>, 并回复 "激活消息", 等待激活成功

#### 点击 <kbd>功能</kbd>, 在按钮列表中点击 <kbd>个人中心</kbd>

#### 点击 <kbd>开发设置</kbd>, 点击 <kbd>Token</kbd> 并一键复制

#### 粘贴 Token 到配置文件的 "push_plus" 下

## 打包

打包到release目录下

```shell
make build # 默认打包为macOS darwin-amd64 Intel处理器
make build ARCH=arm64 # macOS M1处理器
make build OS=linux # linux
make build OS=windows # windows
```

## 执行程序

**将 config.example.json 重命名为 config.yml**

**确保与程序同一目录中包含 config.yml**

### macOS

1. 打开终端, 进入程序所在目录
2. 在终端中执行以下命令

```ssh
./dingdong
```

### windows

1. 进入解压后的程序目录, 在文件夹路径栏输入 cmd 并回车以打开CMD
2. 在CMD中执行以下命令

```ssh
dingdong.exe
```

### 版权说明

**本项目为 GPL3.0 协议，请所有进行二次开发的开发者遵守 GPL3.0 协议，并且不得将代码用于商用。**
