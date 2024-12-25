// auth.js
export function isAuthenticated() {
    // 你的认证逻辑，例如检查 token 是否存在和有效
    return !!localStorage.getItem('token')
}