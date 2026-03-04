# PicoClaw 项目介绍

> 这是一份为Go语言和AI Agent新手编写的PicoClaw项目详细指南。

## 目录

1. [项目概述](#-项目概述)
2. [项目架构](#-项目架构)
3. [核心模块详解](#-核心模块详解)
4. [技术栈](#-技术栈)
5. [运行原理](#-运行原理)
6. [快速开始](#-快速开始)
7. [常见用例](#-常见用例)
8. [学习建议](#-学习建议)

---

## 项目概述

### 什么是PicoClaw？

PicoClaw是一个**超轻量级个人AI助手**，用Go语言编写。它的设计目标是在低成本的硬件上运行（最低只需10美元），同时保持极低的内存占用（<10MB）。

### 主要特点

| 特性 | 说明 |
|------|------|
| 超轻量级 | 内存占用 <10MB，比同类产品小99% |
| 极低成本 | 可以在10美元的硬件上运行 |
| 快速启动 | 1秒内启动，即使在0.6GHz单核上 |
| 真正可移植 | 单一二进制文件，支持RISC-V、ARM和x86 |
| 多通道支持 | 支持Telegram、Discord、Slack等多种消息平台 |

---

## 项目架构

### 目录结构

```
picoclaw/
├── cmd/                    # 命令行应用程序
│   ├── picoclaw/          # 主CLI应用程序
│   │   ├── main.go        # 程序入口
│   │   └── internal/      # 内部命令模块
│   │       ├── auth/      # 认证模块
│   │       ├── cron/      # 定时任务模块
│   │       ├── skills/    # 技能管理模块
│   │       ├── gateway/   # 网关模块
│   │       ├── migrate/   # 数据迁移模块
│   │       ├── onboard/   # 入门引导模块
│   │       ├── status/    # 状态查看模块
│   │       └── version/   # 版本信息模块
│   ├── picoclaw-launcher/        # Web版启动器
│   └── picoclaw-launcher-tui/    # 终端版启动器
├── pkg/                    # 核心库
│   ├── config/            # 配置管理
│   ├── channels/          # 消息通道
│   ├── providers/         # AI提供商
│   ├── skills/            # 技能系统
│   ├── cron/              # 定时任务
│   ├── auth/              # 认证
│   ├── bus/               # 消息总线
│   └── ...
├── workspace/             # 工作区模板
├── docs/                  # 文档
└── docker/                # Docker配置
```

### 模块关系图

```
┌─────────────────────────────────────────────────────────────┐
│                      用户 (User)                             │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                   CLI (picoclaw命令)                         │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐           │
│  │  auth   │ │  cron   │ │ skills  │ │ gateway │  ...      │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘           │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    核心库 (pkg/)                             │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐ ┌─────────┐           │
│  │ config  │ │channels │ │providers│ │ skills  │           │
│  └─────────┘ └─────────┘ └─────────┘ └─────────┘           │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                    外部服务                                  │
│  ┌─────────┐ ┌─────────┐ ┌─────────┐                        │
│  │ Telegram│ │ Discord │ │ OpenAI  │  ...                   │
│  └─────────┘ └─────────┘ └─────────┘                        │
└─────────────────────────────────────────────────────────────┘
```

---

## 核心模块详解

### 1. CLI命令模块 (cmd/picoclaw/internal/)

PicoClaw使用[Cobra](https://github.com/spf13/cobra)库构建命令行界面。每个子命令都是一个独立的模块。

#### 1.1 认证模块 (auth/)

负责用户与AI服务提供商的认证。

```go
// 主要命令
picoclaw auth login     // 登录到AI服务提供商
picoclaw auth logout    // 登出
picoclaw auth status    // 查看认证状态
picoclaw auth models    // 查看可用模型
```

**支持的提供商：**
- OpenAI (OAuth或API Key)
- Anthropic/Claude (API Key)
- Google Antigravity (OAuth)

#### 1.2 网关模块 (gateway/)

网关是PicoClaw的核心服务，负责：
- 接收来自各种消息通道的消息
- 将消息路由到相应的AI代理
- 处理AI响应并发送回消息通道

```go
// 启动网关
picoclaw gateway        // 启动网关服务
picoclaw gateway -d     // 启用调试模式
```

#### 1.3 技能模块 (skills/)

技能(Skills)是预定义的AI能力集合，可以扩展PicoClaw的功能。

```go
picoclaw skills list           // 列出已安装的技能
picoclaw skills search         // 搜索可用技能
picoclaw skills install <name> // 安装技能
picoclaw skills remove <name>  // 移除技能
picoclaw skills show <name>    // 显示技能详情
```

**技能存储位置：**
- 工作区技能：`~/.picoclaw/workspace/skills/`
- 全局技能：`~/.picoclaw/skills/`
- 内置技能：`~/.picoclaw/picoclaw/skills/`

#### 1.4 定时任务模块 (cron/)

管理定时任务，可以在指定时间自动发送消息。

```go
picoclaw cron list             // 列出所有定时任务
picoclaw cron add              // 添加定时任务
picoclaw cron remove <id>      // 删除定时任务
picoclaw cron enable <id>      // 启用定时任务
picoclaw cron disable <id>     // 禁用定时任务
```

**调度方式：**
- 间隔任务：`--every 3600000` (每3600秒)
- Cron表达式：`--cron "0 8 * * *"` (每天早上8点)

#### 1.5 入门引导模块 (onboard/)

首次使用PicoClaw时运行此命令进行初始化。

```go
picoclaw onboard   // 初始化配置和工作区
```

**初始化内容：**
- 创建配置文件目录 `~/.picoclaw/`
- 创建工作区目录
- 复制示例工作区文件

#### 1.6 数据迁移模块 (migrate/)

从其他AI助手（如OpenClaw）迁移配置到PicoClaw。

```go
picoclaw migrate              // 从OpenClaw迁移
picoclaw migrate --dry-run    // 试运行
```

#### 1.7 状态查看模块 (status/)

查看PicoClaw的当前运行状态。

```go
picoclaw status   // 显示状态信息
```

#### 1.8 版本信息模块 (version/)

显示PicoClaw版本信息。

```go
picoclaw version   // 显示版本
```

---

### 2. 核心库 (pkg/)

#### 2.1 配置管理 (pkg/config/)

PicoClaw使用JSON格式的配置文件。

**默认配置路径：** `~/.picoclaw/config.json`

**配置结构：**
```json
{
  "agents": {
    "defaults": {
      "workspace": "~/picoclaw/workspace",
      "model": "gpt-4"
    },
    "list": [
      {
        "id": "default",
        "name": "My Agent"
      }
    ]
  },
  "channels": {
    "telegram": { ... },
    "discord": { ... }
  },
  "providers": { ... },
  "gateway": { ... }
}
```

#### 2.2 消息通道 (pkg/channels/)

支持多种消息平台的集成：

| 通道 | 说明 |
|------|------|
| Telegram | 即时通讯 |
| Discord | 游戏社区平台 |
| Slack | 企业协作 |
| Line | 日本即时通讯 |
| 钉钉 | 企业通讯 |
| 飞书 | 企业协作 |
| 企业微信 | 企业微信 |
| QQ | 腾讯QQ |
| WhatsApp | 即时通讯 |

#### 2.3 AI提供商 (pkg/providers/)

支持多种AI模型：

- **OpenAI**: GPT-4, GPT-3.5
- **Anthropic**: Claude 3
- **Google**: Gemini
- **本地模型**: Ollama

#### 2.4 消息总线 (pkg/bus/)

内部消息传递系统，负责各组件间的通信。

#### 2.5 技能系统 (pkg/skills/)

技能是扩展PicoClaw功能的模块化单元。每个技能包含：
- 技能描述 (SKILL.md)
- 系统提示词
- 工具函数

---

## 技术栈

### 编程语言

- **Go 1.21+**: 项目主要语言

### 关键库

| 库 | 用途 |
|-----|------|
| [spf13/cobra](https://github.com/spf13/cobra) | CLI命令行界面 |
| [spf13/viper](https://github.com/spf13/viper) | 配置管理 |
| [anthropic-sdk-go](https://github.com/anthropics/anthropic-sdk-go) | Anthropic API |
| [google/golang-api](https://github.com/google/google-api-go-client) | Google API |

### 架构模式

1. **插件化设计**: 通道和技能都是可插拔的
2. **消息总线**: 使用发布-订阅模式解耦组件
3. **配置驱动**: 通过配置文件控制行为

---

## 运行原理

### 整体流程

```
1. 用户启动网关
   │
   ▼
2. 加载配置文件
   │
   ▼
3. 初始化消息通道
   │
   ▼
4. 监听各通道的消息
   │
   ▼
5. 收到消息 → 发送到消息总线
   │
   ▼
6. 路由到对应的AI代理
   │
   ▼
7. 调用AI模型处理
   │
   ▼
8. 返回响应 → 发送回消息通道
```

### 网关工作流程

```
                    ┌──────────────┐
                    │   消息通道    │
                    │ (Telegram等) │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  消息管理器   │
                    │  (Manager)   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   消息总线    │
                    │   (MessageBus)│
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   AI代理      │
                    │   (Agent)    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  AI提供商    │
                    │ (OpenAI等)  │
                    └──────────────┘
```

---

## 快速开始

### 1. 安装

```bash
# 克隆项目
git clone https://github.com/sipeed/picoclaw.git
cd picoclaw

# 构建
go build -o picoclaw ./cmd/picoclaw/
```

### 2. 初始化

```bash
# 首次运行，初始化配置
./picoclaw onboard
```

### 3. 配置

编辑配置文件 `~/.picoclaw/config.json`，配置你的AI提供商和消息通道。

### 4. 认证

```bash
# 登录OpenAI
./picoclaw auth login --provider openai
```

### 5. 启动网关

```bash
# 启动网关服务
./picoclaw gateway
```

---

## 常见用例

### 用例1: 设置Telegram机器人

1. 在Telegram @BotFather 创建机器人，获取Token
2. 配置config.json:
```json
{
  "channels": {
    "telegram": {
      "token": "YOUR_BOT_TOKEN"
    }
  }
}
```
3. 启动网关
4. 在Telegram中与机器人对话

### 用例2: 创建定时任务

```bash
# 每天早上8点发送天气提醒
./picoclaw cron add \
  --name "morning-weather" \
  --message "今天天气晴朗，温度25度" \
  --cron "0 8 * * *"
```

### 用例3: 安装技能

```bash
# 安装天气技能
./picoclaw skills install sipeed/picoclaw-skills/weather
```

---

## 学习建议

### 对于Go语言新手

1. **先读标准库**: 熟悉`fmt`、`os`、`json`、`time`等常用包
2. **理解Go的并发**: 重点学习goroutine和channel
3. **阅读项目结构**: 从`cmd/`目录开始，了解如何组织代码
4. **动手实践**: 尝试修改或添加一个小功能

### 对于AI Agent新手

1. **理解提示词工程**: 学习如何编写有效的系统提示词
2. **了解工具调用**: 理解AI如何调用外部工具
3. **学习技能系统**: 技能是扩展AI能力的关键
4. **实践出真知**: 多配置、多尝试、多调试

### 推荐学习路径

```
1. 入门阶段
   ├── 理解项目结构
   ├── 运行onboard命令
   ├── 配置第一个通道
   └── 与AI对话

2. 进阶阶段
   ├── 学习认证系统
   ├── 配置多个AI提供商
   ├── 创建定时任务
   └── 安装和使用技能

3. 高级阶段
   ├── 开发自定义技能
   ├── 扩展消息通道
   ├── 优化性能
   └── 贡献代码
```

---

## 参考资料

- [官方文档](https://picoclaw.io)
- [GitHub仓库](https://github.com/sipeed/picoclaw)
- [Go语言官方文档](https://go.dev/doc/)
- [Cobra库文档](https://github.com/spf13/cobra)

---

> 祝你在PicoClaw的学习之旅中收获满满！🦞
