const express = require('express')
const { createProxyMiddleware } = require('http-proxy-middleware')
const cors = require('cors')

const app = express()
const PORT = 3001

// 启用 CORS
app.use(cors())

// 创建代理中间件
const ipfsProxy = createProxyMiddleware({
  target: 'http://localhost:5001',
  changeOrigin: true,
  pathRewrite: {
    '^/api': '/api', // 保持路径不变
  },
})

// 使用代理
app.use('/api', ipfsProxy)

app.listen(PORT, () => {
  console.log(`IPFS Proxy server running on http://localhost:${PORT}`)
})