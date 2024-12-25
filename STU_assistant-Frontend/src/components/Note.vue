<template>
    <div class="container-note">
        <div class="options">
            <div class="options-text">操作</div>
            <div class="options-fun1">
                <el-button text class="but2" @click="createNote">
                    <el-icon color="#E5EAF3" :size="20">
                        <Plus />
                    </el-icon>
                    <div class="but-text">创建备忘录</div>
                </el-button>
            </div>
            <div class="options-fun2">
                <el-button text class="but2" @click="checkDeadline">
                    <el-icon color="#E5EAF3" :size="20">
                        <Search />
                    </el-icon>
                    <div class="but-text">查看即将截止的事件</div>
                </el-button>
            </div>
            <div class="options-fun3">
                <el-button text class="but2" @click="handleReminder">
                    <el-icon color="#E5EAF3" :size="20">
                        <AlarmClock />
                    </el-icon>
                    <div class="but-text">设置提醒时间</div>
                </el-button>
            </div>
        </div>
        <div class="note-table">
            <el-table v-if="noteData" :data="noteData" height="750" style="width: 100%;" class="custom-table"
                header-cell-class-name="table-header" :row-class-name="tableRowClassName"
                :default-sort="{ prop: 'end_time', order: 'descending' }">
                <el-table-column type="index" width="10" />
                <el-table-column type="expand" width="50">
                    <template #default="props">
                        <div class="detail">
                            <el-popover placement="right" :width="200" trigger="click">
                                <template #reference>
                                    <el-statistic class="detail-sta1" title="题目"
                                        :value="props.row.title"></el-statistic>
                                </template>
                                <div>{{ props.row.title }}</div>
                            </el-popover>
                            <el-popover placement="right" :width="400" trigger="click">
                                <template #reference>
                                    <el-statistic class="detail-sta2" title="内容"
                                        :value="props.row.content"></el-statistic>
                                </template>
                                <div>{{ props.row.content }}</div>
                            </el-popover>
                            <el-statistic class="detail-sta3" title="状态" :value="props.row.status" />
                            <el-statistic class="detail-sta4" title="开始时间" :value="props.row.start_time" />
                            <el-statistic class="detail-sta4" title="结束时间" :value="props.row.end_time" />
                            <el-statistic class="detail-sta4" title="创建时间" :value="props.row.create_time" />
                            <el-statistic class="detail-sta4" title="最近更新时间" :value="props.row.update_time" />
                        </div>
                    </template>
                </el-table-column>
                <el-table-column prop="title" label="题目" width="230" />
                <el-table-column prop="content" label="内容" width="480" />
                <el-table-column prop="start_time" label="开始时间" width="215" />
                <el-table-column prop="end_time" label="结束时间" width="215" sortable />
                <el-table-column prop="status" label="状态" width="100" />
                <el-table-column label="操作" width="180">
                    <template #default="scope">
                        <el-button type="success" @click="handleEdit($index, scope.row)">编辑</el-button>
                        <el-button type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
            <el-empty v-else description="空空如也" />
        </div>
    </div>
    <!-- 编辑备忘录 -->
    <el-dialog title="编辑备忘录" v-model="editDialogVisible" width="400px" align-center>
        <el-form :model="editForm">
            <el-form-item label="题目">
                <el-input v-model="editForm.title"></el-input>
            </el-form-item>
            <el-form-item label="内容">
                <el-input v-model="editForm.content" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="开始时间">
                <div class="date-picker-container">
                    <el-date-picker class="data-picker" type="datetime" v-model="editForm.start_time" arrow-control
                        placeholder="start_time" :shortcuts="shortcuts" />
                </div>
            </el-form-item>
            <el-form-item label="结束时间">
                <div class="date-picker-container">
                    <el-date-picker class="data-picker" type="datetime" v-model="editForm.end_time" arrow-control
                        placeholder="end_time" :shortcuts="shortcuts" />
                </div>
            </el-form-item>
            <el-form-item label="状态">
                <el-select v-model="editForm.status" placeholder="Select Status">
                    <el-option label="未完成" value="0"></el-option>
                    <el-option label="已完成" value="1"></el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button @click="editDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitEdit">确认</el-button>
        </template>
    </el-dialog>
    <!-- 添加备忘录 -->
    <el-dialog title="创建备忘录" v-model="createDialogVisible" width="400px" align-center draggable append-to-body="true">
        <el-form :model="createForm">
            <el-form-item label="题目">
                <el-input v-model="createForm.title"></el-input>
            </el-form-item>
            <el-form-item label="内容">
                <el-input v-model="createForm.content" type="textarea"></el-input>
            </el-form-item>
            <el-form-item label="开始时间">
                <div class="date-picker-container">
                    <el-date-picker class="data-picker" type="datetime" v-model="createForm.start_time" arrow-control
                        placeholder="start_time" :shortcuts="shortcuts" />
                </div>
            </el-form-item>
            <el-form-item label="结束时间">
                <div class="date-picker-container">
                    <el-date-picker class="data-picker" type="datetime" v-model="createForm.end_time" arrow-control
                        placeholder="end_time" :shortcuts="shortcuts" />
                </div>
            </el-form-item>
            <el-form-item label="状态">
                <el-select v-model="createForm.status" placeholder="Select Status">
                    <el-option label="未完成" value="0"></el-option>
                    <el-option label="已完成" value="1"></el-option>
                </el-select>
            </el-form-item>
        </el-form>
        <template #footer>
            <el-button @click="createDialogVisible = false">取消</el-button>
            <el-button type="primary" @click="submitCreateNote">确认</el-button>
        </template>
    </el-dialog>
    <!-- 设置提醒时间 -->
    <el-dialog title="设置提醒时间" v-model="createReminderVisible" width="400px" align-center draggable
        append-to-body="true">
        <div class="date-picker-container">
            <el-date-picker class="data-picker" type="datetime" v-model="reminderForm.deadline" arrow-control
                placeholder="deadline" :shortcuts="shortcuts" />
        </div>
        <template #footer>
            <el-button @click="createReminderVisible = false">取消</el-button>
            <el-button type="primary" @click="createReminder">确认</el-button>
        </template>
    </el-dialog>
    <!-- 查看截止事件 -->
    <el-dialog title="即将截止事件" v-model="deadlineVisible" width="550px" align-center draggable append-to-body="true">
        <el-table v-if="deadlineData" :data="deadlineData" height="500px" style="width: 100%;"
            :default-sort="{ prop: 'end_time', order: 'descending' }">
            <el-table-column type="index" width="30" />
            <el-table-column prop="title" label="题目" width="100" />
            <el-table-column prop="content" label="内容" width="200" />
            <el-table-column prop="end_time" label="结束时间" width="180" sortable />
        </el-table>
        <el-empty v-else description="空空如也" />
    </el-dialog>
</template>

<script>
import { createTask, deleteTask, getReminder, getTasks, putReminder, putTask } from '@/utils/api/api';
import { convertToTimestamp, formatTime } from '@/utils/time';
export default {
    name: 'Note',
    data() {
        return {
            noteData: [],
            deadlineData: [],
            editDialogVisible: false, // 控制编辑对话框显示/隐藏
            createDialogVisible: false,
            createReminderVisible: false,
            deadlineVisible: false,
            editForm: { // 用于编辑的任务数据
                id: null,
                title: "",
                content: "",
                status: "",
                start_time: '',
                end_time: '',
            },
            createForm: {
                title: "",
                content: "",
                status: "",
                start_time: '',
                end_time: '',
            },
            reminderForm: {
                id: JSON.parse(localStorage.getItem('user')).id,
                deadline: '',
            },
            shortcuts: [
                {
                    text: '一天后',
                    value: () => {
                        const date = new Date()
                        date.setDate(date.getDate() + 1)
                        return date
                    },
                },
                {
                    text: '三天后',
                    value: () => {
                        const date = new Date()
                        date.setDate(date.getDate() + 3)
                        return date
                    },
                },
                {
                    text: '一周后',
                    value: () => {
                        const date = new Date()
                        date.setDate(date.getDate() + 7)
                        return date
                    },
                }
            ],
            intervalId: null, // 保存定时器 ID
        }
    },
    created() {
        this.refreshNote();
        this.getDeadline();
        this.startAutoRefresh();
    },
    beforeUnmount() {
        this.stopAutoRefresh(); // 组件销毁时停止自动刷新
    },
    methods: {
        startAutoRefresh() {
            // 设置每 10 秒刷新一次
            this.intervalId = setInterval(() => {
                this.refreshNote();
            }, 60000); // 10000 毫秒 = 10 秒
        },
        stopAutoRefresh() {
            // 清除定时器
            if (this.intervalId) {
                clearInterval(this.intervalId);
                this.intervalId = null;
            }
        },
        tableRowClassName({ row, rowIndex }) {
            if (row.status === '进行中') {
                return 'warning-row';
            } else if (row.status === '已完成') {
                return 'success-row';
            }
            return '';
        },
        refreshNote() {
            getTasks().then(response => {
                // console.log(response);
                this.noteData = response.data.task_list.map(task => {
                    // 根据 status 的值转化为对应的文本
                    if (task.status === 0) {
                        task.status = '进行中';  // status 为 0 时转为 '进行中'
                    } else if (task.status === 1) {
                        task.status = '已完成';  // status 为 1 时转为 '已完成'
                    }
                    task.create_time = formatTime(task.create_time);
                    task.update_time = formatTime(task.update_time);
                    task.start_time = formatTime(task.start_time);
                    task.end_time = formatTime(task.end_time);
                    return task;
                });
            }).catch(error => {
                console.log(error);
            })
        },
        checkDeadline() {
            this.getDeadline();
            this.deadlineVisible = true;
        },
        getDeadline() {
            getReminder().then(response => {
                if (response.data.count > 0) {
                    // console.log(response);
                    this.deadlineData = response.data.task_list.map(task => {
                        const [title, content] = task.title.split('|'); // 按第一个 '|' 分割
                        task.title = title; // 更新 title 属性
                        task.content = content || ''; // 添加 content 属性，确保安全处理无 '|' 的情况
                        task.end_time = formatTime(task.end_time); // 格式化 end_time
                        return task;
                    });
                } else {
                    this.deadlineData = null;
                }
            }).catch(error => {
                console.log(error);
            })
        },
        // 编辑任务
        handleEdit(index, row) {
            // 弹出一个对话框让用户编辑内容
            this.editDialogVisible = true;
            this.editForm = {
                id: row.id,
                title: row.title,
                content: row.content,
                status: row.status,
                start_time: row.start_time,
                end_time: row.end_time,
            };
        },
        submitEdit() {
            if (!this.editForm) return;
            // 调用 putTask 更新任务
            const id = this.editForm.id;
            const title = this.editForm.title;
            const content = this.editForm.content;
            let status = null;
            if (this.editForm.status === '进行中') {
                status = 0;
            } else if (this.editForm.status === '已完成') {
                status = 1;
            } else {
                status = this.editForm.status;
            }
            const start_time = convertToTimestamp(this.editForm.start_time);
            const end_time = convertToTimestamp(this.editForm.end_time);
            putTask(id, title, content, status, start_time, end_time)
                .then(() => {
                    // 更新成功后，刷新前端数据
                    this.refreshNote();
                    this.editDialogVisible = false; // 关闭对话框
                    this.$message({
                        type: "success",
                        message: "备忘录更新成功!",
                    });
                })
                .catch((error) => {
                    console.error("Error updating task:", error);
                    this.$message.error("更新备忘录失败！");
                });
        },

        // 删除任务
        handleDelete(index, row) {
            // 弹出删除确认框
            this.$confirm("确定删除该备忘录吗？", "删除", {
                confirmButtonText: "删除",
                cancelButtonText: "取消",
                type: "warning",
            })
                .then(() => {
                    // 调用 deleteTask 删除任务
                    deleteTask(row.id)
                        .then(() => {
                            // 删除成功后，移除前端数据
                            this.noteData.splice(index, 1);
                            this.refreshNote();
                            this.$message({
                                type: "success",
                                message: "删除成功",
                            });
                        })
                        .catch((error) => {
                            console.error("Error deleting task:", error);
                            this.$message.error("删除失败");
                        });
                })
                .catch(() => {
                    console.log("Delete canceled");
                });
        },
        createNote() {
            this.createDialogVisible = true;
        },
        submitCreateNote() {
            if (!this.createForm.title || !this.createForm.content || !this.createForm.status) {
                this.$message.info("请输入信息");
                return;
            }
            const { title, content, status, start_time, end_time } = this.createForm;
            createTask(title, content, status, convertToTimestamp(start_time), convertToTimestamp(end_time)).then(response => {
                if (response.status == 200) {
                    this.$message({
                        type: "success",
                        message: "创建成功",
                    });
                    this.createDialogVisible = false;
                    this.refreshNote();
                }
            }).catch((error) => {
                console.error("Error create task:", error);
                this.$message.error("创建失败");
            });

        },
        handleReminder() {
            this.createReminderVisible = true;
        },
        createReminder() {
            if (!this.reminderForm.deadline) {
                this.$message.info("请输入信息");
                return;
            }
            if (!this.reminderForm.id) {
                const user = JSON.parse(localStorage.getItem('user'));
                this.reminderForm.id = user.id;
                // console.log(user);
            }
            const id = this.reminderForm.id;
            const currentTime = Math.floor(Date.now() / 1000);
            const deadline = convertToTimestamp(this.reminderForm.deadline) - currentTime;
            if (deadline < 0) {
                this.$message.error("截止时间不能小于当前时间！");
                return;
            }
            putReminder(id, deadline).then(response => {
                if (response.status == 200) {
                    this.$message({
                        type: "success",
                        message: "设置成功",
                    });
                    this.createReminderVisible = false;
                    this.refreshNote();
                }
            }).catch((error) => {
                console.error("Error create deadline:", error);
                this.$message.error("设置失败");
            });
        }
    }
}
</script>

<style scoped>
.container-note {
    height: 100%;
    width: 100%;
}

.note-table {
    margin-top: 40px;
}

::v-deep(.table-header .cell) {
    color: #000000;
    /* 表头字体颜色 */
}

/* 设置表格内容字体大小 */
::v-deep(.el-table .cell) {
    font-size: 15px;
    /* 你可以根据需要调整字体大小 */
}

.detail {
    display: flex;
    /* justify-content: center; */
    align-items: center;
    width: 100%;
    height: 100%;
}

.detail-sta1 {
    margin-left: 30px;
    width: 5%;
    height: 50px;
    align-items: center;
    font-size: 10px;
    word-wrap: break-word;
    /* 自动换行 */
    word-break: break-all;
    /* 长单词断行 */
    white-space: normal;
    /* 启用换行 */
    overflow: hidden;
    /* 隐藏多余部分 */
}

.detail-sta2 {
    margin-left: 30px;
    width: 15%;
    height: 50px;
    align-items: center;
    font-size: 10px;
    word-wrap: break-word;
    /* 自动换行 */
    word-break: break-all;
    /* 长单词断行 */
    white-space: normal;
    /* 启用换行 */
    overflow: hidden;
    /* 隐藏多余部分 */
}

.detail-sta3 {
    margin-left: 30px;
    width: 5%;
    align-items: center;
    font-size: 10px;
}

.detail-sta4 {
    margin-left: 30px;
    width: 15%;
    align-items: center;
    font-size: 10px;
}

::v-deep(.el-table .warning-row) {
    --el-table-tr-bg-color: var(--el-color-warning-light-9);
}

::v-deep(.el-table .success-row) {
    --el-table-tr-bg-color: var(--el-color-success-light-9);
}

.options {
    height: 50px;
    width: 100%;
    display: flex;
}

.options-fun1 {
    width: 200px;
    margin-top: 35px;
}

.options-fun2 {
    width: 200px;
    margin-top: 35px;
}

.options-fun3 {
    width: 200px;
    margin-top: 35px;
    margin-left: 70px;
}

.options-text {
    color: rgb(204, 202, 202);
    margin-left: 0px;
    font-size: 20px;
    height: 25px;
    width: 50px;
}

.but2 {
    margin-top: 10px;
    width: 100%;
    color: #E5EAF3;
    font-size: 20px;
    display: flex;
    justify-content: flex-start;
    /* 图标和文字靠左对齐 */
    gap: 5px;
    /* 图标和文字之间的间距 */
}

.but-text {
    padding-left: 10px;
}

::v-deep(.but2:hover) {
    background-color: #1b1b1b !important;
    /* 强制覆盖 */
    color: rgb(228, 228, 228);
    /* 如果需要，也可以修改文字颜色 */
    transition: background-color 0.3s ease;
    /* 增加过渡效果 */
}

.date-picker-container {
    display: flex;
    /* 启用 Flex 布局 */
    justify-content: center;
    /* 水平居中 */
    align-items: center;
    /* 垂直居中 */
    height: 100%;
    /* 根据需要调整高度 */
}

.data-picker {
    width: 100%;
    /* 根据需要调整宽度 */
}
</style>