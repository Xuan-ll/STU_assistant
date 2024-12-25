# microtodolist
本项目为MicroTodolist前端开发部分。

开发成员：

- 刘江浩
- 倪浩哲

## 主要技术

Vue3 + axios + router + Element Plus

## 项目结构

```
microtodolist
├─ .gitignore
├─ index.html			  # 应用的主 HTML 文件，Vue CLI 会在构建时自动注入生成的静态资源链接
├─ jsconfig.json
├─ package-lock.json
├─ package.json
├─ public				  # 静态资源文件夹，其中的文件会直接复制到构建输出目录中，无需经过编译处理
│  └─ favicon.ico	      # 网站图标
├─ README.md
├─ src					
│  ├─ App.vue			  # 应用程序根组件，整个应用的入口点，通常包含路由视图和其他全局共享组件
│  ├─ assets			  # 存放静态资源，如图像、字体等。这些文件会由 Webpack 处理，可以通过相对路径引用
│  ├─ components	      # 组件目录，按照功能或类别划分存放单文件组件（.vue文件）
│  ├─ main.js			  # 应用入口文件，用于初始化Vue实例、引入并配置路由、状态管理等核心模块并挂载到 DOM 上
│  ├─ router			  # 路由配置目录，主要包含index.js路由文件，用于配置应用的路由规则
│  │  └─ index.js
│  ├─ utils				  # 工具函数和类库目录，存放项目中常用的工具函数、辅助类等
│  │  ├─ api			  # api解耦
│  │  │  └─ api.js
│  │  └─ request.js	      # 将axios进行封装
│  └─ views				  # 存放视图组件，通常对应路由，每个视图都是一个独立的 .vue 文件
│     └─ TestApiView.vue
└─ vite.config.js

# 可重用的内容可以保存在 src/components文件夹中，与路由器绑定的内容可以保存在 src/views中
# 变量和方法一般使用小驼峰命名法
# 组件名应使用大驼峰命名法
```


## Project Setup

```sh
npm install
```

### Compile and Hot-Reload for Development

```sh
npm run dev
```
