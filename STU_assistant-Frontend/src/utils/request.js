// 引入axios
import axios from "axios";

// 如果 switch 为 true，则使用 devBaseURL，否则使用 apiBaseURL 
const apiBaseURL = window.config.devMode ? window.config.devBaseURL : window.config.prodBaseURL;
console.log("apiBaseURL:", apiBaseURL);

// 创建axios对象
const service = axios.create({
    baseURL: apiBaseURL,  //链接
    timeout: 5000, // 设置超时限制/ms
    headers:{
      'Content-Type': 'multipart/form-data',
    }
});

// 请求拦截器
service.interceptors.request.use(config => {
    // 可以在这里添加请求头等操作
    // 从本地存储获取 Token
    const token = localStorage.getItem('token');
    if (token) { //先确保登录
      config.headers['Authorization'] = token;
	  }
    // console.log(config);
    return config;
  }, error => {
    return Promise.reject(error);
});

// 响应拦截器 接收后端发往前端数据
service.interceptors.response.use(response => {
    const code = response.status;
    console.log(code);
    // console.log(response,'请求成功！') //查看返回数据
    return Promise.resolve(response.data)
  }, async error => {
    return Promise.reject(error);
  });

export default service