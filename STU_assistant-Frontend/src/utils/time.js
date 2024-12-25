// utils.js
export function formatTime(timestamp) {
    if (timestamp === 0) return "未设置"; // 特殊处理 0 的情况
    const date = new Date(timestamp * 1000); // 时间戳转换为毫秒
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // 月份从 0 开始
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`; // 格式化为字符串
}
export function convertToTimestamp(datetime) {
    // 使用 Date 对象将字符串解析为时间对象
    const date = new Date(datetime);
    // 返回时间戳（以秒为单位）
    return Math.floor(date.getTime() / 1000);
}