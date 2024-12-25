import request from '@/utils/request'

// 用户注册
export function userRegister(user_name, password, password_confirm){
    const formData = new FormData();
    formData.append('user_name',user_name);
    formData.append('password',password);
    formData.append('password_confirm',password_confirm);
    return request({
        url:'/api/v1/user/register',
        method:'POST',
        data:formData,
    })
}
// 用户登录
export function userLogin(user_name, password){
    const formData = new FormData();
    formData.append('user_name',user_name);
    formData.append('password',password);
    return request({
        url:'/api/v1/user/login',
        method:'POST',
        data:formData,
    })
}
// 更新token
export function getToken(){
    return request({
        url:'/api/v1/renew',
        method:'GET',
    })
}
// 创建备忘录
export function createTask(title, content, status, start_time, end_time){
    const formData = new FormData();
    formData.append('title',title);
    formData.append('content',content);
    formData.append('status',status);
    formData.append('start_time',start_time);
    formData.append('end_time',end_time);
    return request({
        url:'/api/v1/task',
        method:'POST',
        data:formData,
    })
}
// 获取备忘录
export function getTasks(){
    return request({
        url:'/api/v1/tasks',
        method:'GET'
    })
}
// 更新备忘录
export function putTask(taskId, title, content, status, start_time, end_time){
    const formData = new FormData();
    formData.append('title',title);
    formData.append('content',content);
    formData.append('status',status);
    formData.append('start_time',start_time);
    formData.append('end_time',end_time);
    return request({
        url:`/api/v1/task/${taskId}`,
        method:'PUT',
        data:formData,
    })
}
// 获取备忘录详情
export function getTaskDetails(taskId){
    return request({
        url:`/api/v1/task/${taskId}`,
        method:'GET'
    })
}
// 删除备忘录
export function deleteTask(taskId){
    return request({
        url:`/api/v1/task/${taskId}`,
        method:'DELETE'
    })
}
// 获取快要截至的待办事项
export function getReminder(){
    return request({
        url:'/api/v1/reminders',
        method:'GET'
    })
}
// 设置提醒时间
export function putReminder(taskId, dead_line){
    const formData = new FormData();
    formData.append('dead_line',dead_line);
    return request({
        url:`/api/v1/reminders/1`,
        method:'PUT',
        data:formData,
    })
}
// 创建新课程
export function createCourse(course_name, start_class, end_class, start_week, end_week, day_of_week){
    const formData = new FormData();
    formData.append('course_name',course_name);
    formData.append('start_class',start_class);
    formData.append('end_class',end_class);
    formData.append('start_week',start_week);
    formData.append('end_week',end_week);
    formData.append('day_of_week',day_of_week);
    return request({
        url:'/api/v1/course',
        method:'POST',
        data:formData,
    })
}
// 获取课表
export function getCourse(week){
    return request({
        url:`/api/v1/course/${week}`,
        method:'GET'
    })
}
// 删除课程
export function deleteCourse(id){
    return request({
        url:`/api/v1/course/${id}`,
        method:'DELETE'
    })
}
// 更新课程
export function editCourse(id, course_name, start_class, end_class, start_week, end_week, day_of_week){
    const formData = new FormData();
    formData.append('course_name',course_name);
    formData.append('start_class',start_class);
    formData.append('end_class',end_class);
    formData.append('start_week',start_week);
    formData.append('end_week',end_week);
    formData.append('day_of_week',day_of_week);
    return request({
        url:`/api/v1/course/${id}`,
        method:'PUT',
        data:formData,
    })
}
