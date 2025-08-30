# IPFS-X Frontend 项目文档

## 项目概述

IPFS-X 前端项目是一个基于 React 和 TypeScript 构建的 Sui 区块链 dApp，集成了 IPFS 功能，用于在去中心化网络上存储和检索数据。该项目使用了 Mysten 的 dApp 工具包和 SuiWare 套件，提供现代化的用户界面和流畅的交互体验。

## 项目结构

```
src/
├── assets/          # 静态资源文件
│   └── logo.svg     # 应用程序logo
│
├── components/      # 通用UI组件
│   ├── AnimatedBackground/  # 动画背景组件
│   ├── App.tsx              # 主应用组件
│   ├── CustomConnectButton.tsx  # 自定义连接钱包按钮
│   ├── Loading.tsx          # 加载指示器
│   ├── NetworkSupportChecker.tsx # 网络支持检查器
│   ├── Notification.tsx     # 通知组件
│   ├── ThemeSwitcher.tsx   # 主题切换器
│   └── layout/             # 布局相关组件
│       ├── Body.tsx       # 页面主体
│       ├── Extra.tsx      # 额外内容
│       ├── Footer.tsx     # 页脚
│       └── Header.tsx     # 页头
│
├── config/          # 配置文件
│   ├── main.ts            # 主配置
│   ├── network.ts         # 网络配置
│   └── themes.ts          # 主题配置
│
├── dapp/            # dApp 特定组件和功能
│   ├── components/        # dApp 特定组件
│   │   ├── GreetingForm.tsx  # 问候表单组件
│   │   └── Emoji.tsx         # 表情符号组件
│   ├── config/            # dApp 配置
│   ├── helpers/           # dApp 辅助函数
│   │   └── transactions.ts # 交易相关函数
│   ├── hooks/             # dApp 自定义钩子
│   │   └── useOwnGreeting.tsx # 获取用户问候语的自定义钩子
│   ├── pages/             # 页面组件
│   │   └── IndexPage.tsx # 主页
│   └── types/             # dApp 类型定义
│
├── helpers/         # 通用辅助函数
│   ├── misc.ts           # 杂项辅助函数
│   ├── network.ts        # 网络相关辅助函数
│   ├── notification.tsx  # 通知辅助函数
│   └── theme.ts          # 主题辅助函数
│
├── hooks/           # 全局自定义钩子
│   └── useNetworkConfig.tsx # 网络配置钩子
│
├── providers/       # 上下文提供者
│   └── ThemeProvider.tsx # 主题提供者
│
├── styles/          # 样式文件
│   └── index.css         # 全局样式
│
├── types/           # 全局类型定义
│   ├── ENetwork.ts       # 网络类型
│   ├── ENetworksWithFaucet.ts # 支持水龙头网络的类型
│   └── TTheme.ts         # 主题类型
│
├── main.tsx         # 应用程序入口点
└── vite-env.d.ts    # Vite 环境类型定义
```

## 核心功能

1. **钱包集成**：支持 Sui 钱包连接，包括自定义连接按钮
2. **IPFS 集成**：与 IPFS 网络集成，用于存储和检索去中心化数据
3. **NFT 生成**：基于用户输入生成 NFT，并显示在界面上
4. **网络管理**：支持多网络切换，包括本地网络和测试网络
5. **主题切换**：支持明暗主题切换
6. **通知系统**：提供交易状态通知

## 主要依赖项

- **React 19.1.0**：前端框架
- **TypeScript**：静态类型检查
- **Vite**：构建工具
- **Tailwind CSS**：样式框架
- **Radix UI Themes**：UI 组件库
- **Mysten dApp Kit**：Sui dApp 开发工具包
- **SuiWare Kit**：Sui 开发工具套件
- **React Query**：数据获取和状态管理

## 开发脚本

- `npm run dev`：启动开发服务器
- `npm run build`：构建生产版本
- `npm run lint`：运行 ESLint 检查
- `npm run preview`：预览生产构建
- `npm run format`：使用 Prettier 格式化代码

## 部署选项

项目支持多种部署方式：
- Firebase 托管
- Walrus 测试网/主网
- Arweave 去中心化网络

## 请参考

请查看根项目 [README](../../README.md) 获取更多项目信息。
