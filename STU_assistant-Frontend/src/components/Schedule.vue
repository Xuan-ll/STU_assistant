<template>
  <div class="container-schedule">
    <div class="ele-body">
      <el-card shadow="never">
        <!-- 表标题 -->
        <div class="table-heard">
          <el-tag type="info">
            {{ this.currentWeekRange[0] + '至' + this.currentWeekRange[1] }}
          </el-tag>
          <div class="center">
            <el-button @click="weekPre" size="mini">
              <i class="el-icon-arrow-left"></i>上一周
            </el-button>
            <p>{{ currentYear }}年{{ currentMonth }}月</p>
            <el-button @click="weekNext" size="mini">
              下一周<i class="el-icon-arrow-right"></i>
            </el-button>
          </div>
          <div>
            <el-button @click="addCourse" size="mini">
              添加课程
            </el-button>
            第{{ currentWeekNo }}周
          </div>
        </div>
        <table align="center" cellpadding="0" cellspacing="0">
          <!-- 顶部日期行 -->
          <tr class="sepacialy">
            <td></td>
            <td v-for="(item, index) in weekGround" :key="index">{{ item }}</td>
          </tr>
          <!-- 正表 -->
          <tr v-for="index in 13" :key="index">
            <td class="time">
              <p class="timeNo">{{ index }}</p>
              <p>{{ defaultClassBegin[index - 1] }}</p>
              <p>{{ defaultClassEnd[index - 1] }}</p>
            </td>
            <td align="center" valign="middle" v-for="index1 in 7" :key="index1">
              <div v-for="course in getCoursesForDay(index1)" :key="course.uid">
                <div class="outblock"
                  :style="{ backgroundColor: getBackgroundColor(course.courseName), cursor: pointer }"
                  v-if="index >= course.startClass && index <= course.endClass" @click="editCourse(course)"
                  style="cursor: pointer;">
                  <p class="coursename">{{ course.courseName }}</p>
                  <p>{{ defaultClassBegin[course.startClass - 1] }} - {{ defaultClassEnd[course.endClass - 1] }}</p>
                </div>
              </div>
            </td>
          </tr>

        </table>
      </el-card>
    </div>

    <!-- 弹窗 -->
    <el-dialog v-model="dialogVisible" :title="isEdit ? '编辑课程' : '添加课程'" @close="resetForm">
      <el-form :model="form" label-width="120px" :rules="rules" ref="formRef">
        <el-form-item label="课程名称" prop="courseName">
          <el-input v-model="form.courseName" placeholder="请输入课程名称"></el-input>
        </el-form-item>
        <el-form-item label="上课节数" prop="startClass">
          <el-input-number v-model.number="form.startClass" type="number" :min="1" :max="13"
            :precision="0"></el-input-number>
        </el-form-item>
        <el-form-item label="下课节数" prop="endClass">
          <el-input-number v-model.number="form.endClass" type="number" :min="1" :max="13"
            :precision="0"></el-input-number>
        </el-form-item>
        <el-form-item label="起始周数" prop="startWeek">
          <el-input-number v-model.number="form.startWeek" type="number" :min="1" :max="20"
            :precision="0"></el-input-number>
        </el-form-item>
        <el-form-item label="结束周数" prop="endWeek">
          <el-input-number v-model.number="form.endWeek" type="number" :min="1" :max="20"
            :precision="0"></el-input-number>
        </el-form-item>
        <el-form-item label="周几上课" prop="dayOfWeek">
          <el-input-number v-model.number="form.dayOfWeek" type="number" :min="1" :max="7"
            :precision="0"></el-input-number>
        </el-form-item>
      </el-form>

      <!-- 提交按钮 -->
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogVisible = false">取消</el-button>
        <el-button type="primary" @click="submitForm">提交</el-button>
        <el-button class="delete-button" v-if="isEdit" @click="deleteForm">删除</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import axios from 'axios';
import { deleteCourse, createCourse, editCourse, getCourse } from '@/utils/api/api.js';
export default {
  name: 'Schedule',
  props: {

  },
  data() {
    return {
      // 一周7天
      weekGround: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],
      course: [],

      currentYear: '', // 年份
      currentMonth: '', // 月份
      currentDay: '', // 日期
      currentWeek: '', // 星期
      currentWeekNo: 17, // 当前周数
      days: [],

      //   当前周的日期范围
      currentWeekRange: [],

      // 默认周
      defalultWeekGround: ['周一', '周二', '周三', '周四', '周五', '周六', '周日'],

      // 默认时间
      defaultClassBegin: ['08:00', '08:50', '09:55', '10:45', '11:35', '13:20', '14:10', '15:15', '16:05', '16:55', '18:30', '19:20', '20:10'],
      defaultClassEnd: ['08:45', '09:35', '10:40', '11:30', '12:20', '14:05', '14:55', '16:00', '16:50', '17:40', '19:15', '20;05', '20:55'],

      // 颜色map映射
      colorOfClass: {},

      // 弹窗显示
      dialogVisible: false,
      currentCourse: {},    // 当前选中的课程
      isEdit: false,

      form: {
        courseName: '',
        startClass: '',
        endClass: '',
        startWeek: '',
        endWeek: '',
        dayOfWeek: '',
        id: '',
      },

      rules: {},
      uid: "",
    };
  },
  methods: {
    // 日期格式化
    formatDate(year, month, day) {
      const y = year;
      let m = month;
      if (m < 10) m = `0${m}`;
      let d = day;
      if (d < 10) d = `0${d}`;
      return `${y}-${m}-${d}`;
    },
    // 日期初始化
    initDate(cur) {
      let date = '';
      if (cur) {
        date = new Date(cur);
      } else {
        date = new Date();
      }
      this.currentDay = date.getDate(); // 今日日期 几号
      this.currentYear = date.getFullYear(); // 当前年份
      this.currentMonth = date.getMonth() + 1; // 当前月份
      this.currentWeek = date.getDay(); // 1...6,0   // 星期几
      if (this.currentWeek === 0) {
        this.currentWeek = 7;
      }
      const str = this.formatDate(
        this.currentYear,
        this.currentMonth,
        this.currentDay
      ); // 今日日期 年-月-日

      this.days.length = 0;
      // 今天是周日，放在第一行第7个位置，前面6个 这里默认显示一周，如果需要显示一个月，则第二个循环为 i<= 35- this.currentWeek
      /* eslint-disabled */
      for (let i = this.currentWeek - 1; i >= 0; i -= 1) {
        const d = new Date(str);
        d.setDate(d.getDate() - i);
        this.days.push(d);
      }
      for (let i = 1; i <= 7 - this.currentWeek; i += 1) {
        const d = new Date(str);
        d.setDate(d.getDate() + i);
        this.days.push(d);
      }

      //  当前周的日期范围
      this.currentWeekRange = [
        this.formatDate(
          this.days[0].getFullYear(),
          this.days[0].getMonth() + 1,
          this.days[0].getDate()
        ),
        this.formatDate(
          this.days[6].getFullYear(),
          this.days[6].getMonth() + 1,
          this.days[6].getDate()
        )
      ];
      //获取到这周7个日期,不要年份，eg:周一(7/31)
      let dataArr = [];

      this.days.forEach((item) => {
        dataArr.push(`${item.getMonth() + 1}/${item.getDate()}`);
      });

      this.defalultWeekGround.map((item, index) => {
        this.weekGround[index] = item + `(${dataArr[index]})`;
      });
    },
    //  上个星期
    weekPre() {
      if (this.currentWeekNo == 1) {
        alert("已经是第1周了!");
        return;
      }
      const d = this.days[0]; // 如果当期日期是7号或者小于7号
      d.setDate(d.getDate() - 7);
      this.currentWeekNo = Math.max(1, this.currentWeekNo - 1);
      this.initDate(d);

      this.fetchCourse();
    },
    //  下个星期
    weekNext() {
      if (this.currentWeekNo == 20) {
        alert("已经是最后一周了!");
        return;
      }
      const d = this.days[6]; // 如果当期日期是7号或者小于7号
      d.setDate(d.getDate() + 7);
      this.currentWeekNo = Math.min(20, this.currentWeekNo + 1);
      this.initDate(d);

      this.fetchCourse();
    },
    // 获取当前周星期dayIndex的课程数据
    getCoursesForDay(dayIndex) {
      let coursesForDay = [];
      // console.log("getCoursesForDay: ", this.course);
      this.course.forEach(course1 => {
        if (course1.dayOfWeek == dayIndex && course1.startWeek <= this.currentWeekNo && course1.endWeek >= this.currentWeekNo) {
          coursesForDay.push(course1);
        }
      });
      // console.log(coursesForDay);
      return coursesForDay;
    },
    // 随机生成背景颜色
    getLightColor() {
      const getRandomChannel = () => Math.floor(Math.random() * 76) + 180;  // 随机值在180到255之间

      // 生成一个淡色 RGB 颜色
      const r = getRandomChannel();
      const g = getRandomChannel();
      const b = getRandomChannel();

      // 转换为十六进制颜色字符串
      const randomColor = `#${r.toString(16).padStart(2, '0')}${g.toString(16).padStart(2, '0')}${b.toString(16).padStart(2, '0')}`;
      return randomColor;
    },
    // 匹配/生成课程背景颜色
    getBackgroundColor(courseName) {
      // 如果courseName已经存在于colorOfClass字典中，直接返回对应的颜色
      if (this.colorOfClass[courseName]) {
        return this.colorOfClass[courseName];
      } else {
        // 如果不存在，生成一个随机的淡色，并加入字典中
        const newColor = this.getLightColor();
        this.colorOfClass[courseName] = newColor; // 将颜色存入字典
        return newColor;
      }
    },
    // 添加课程
    addCourse() {
      this.form = this.getInitialFormData(); // 重置表单数据
      this.dialogVisible = true;
      this.$nextTick(() => {
        // 确保表单已渲染
        if (this.$refs.formRef) {
          console.log("formRef is now available");
        } else {
          console.error("formRef is still undefined after nextTick");
        }
      });
    },
    // 修改课程
    editCourse(course) {
      // 深拷贝当前课程数据，避免直接修改原数据
      this.isEdit = true;
      this.currentCourse = course;

      this.form = {
        courseName: course.courseName || '',
        startClass: course.startClass || 1,
        endClass: course.endClass || 1,
        startWeek: course.startWeek || 1,
        endWeek: course.endWeek || 1,
        dayOfWeek: course.dayOfWeek || 1,
        id: course.id,
      };

      this.dialogVisible = true;
      this.$nextTick(() => {
        // 确保表单已渲染
        if (this.$refs.formRef) {
          console.log("formRef is now available");
        } else {
          console.error("formRef is still undefined after nextTick");
        }
      });
    },
    // 提交表单
    submitForm() {
      this.$refs.formRef.validate((valid) => {
        if (valid && !this.isEdit) {
          createCourse(this.form.courseName, +this.form.startClass, +this.form.endClass, +this.form.startWeek, +this.form.endWeek, +this.form.dayOfWeek)
            .then(response => {
              console.log('msg:', response.msg);
              console.log('error:', response.error);
              if (response.msg == "ok") {
                this.$message({
                  message: '课程提交成功！',
                  type: 'success',
                  duration: 2000,
                });
                this.dialogVisible = false; // 关闭弹窗
                this.resetForm();          // 重置表单
                return this.fetchCourse();
              }
            })
            .catch(error => {
              console.error('提交失败:', error);
              this.$message({
                message: '课程提交失败，请重试！',
                type: 'error',
                duration: 2000,
              });
            });
        } else if (valid && this.isEdit) {
          editCourse(this.form.id, this.form.courseName, +this.form.startClass, +this.form.endClass, +this.form.startWeek, +this.form.endWeek, +this.form.dayOfWeek)
            .then(response => {
              console.log('msg:', response.msg);
              console.log('error:', response.error);
              if (response.msg == "ok") {
                this.$message({
                  message: '课程修改成功！',
                  type: 'success',
                  duration: 2000,
                });
                this.dialogVisible = false; // 关闭弹窗
                this.resetForm();          // 重置表单
                return this.fetchCourse();
              }
            })
            .catch(error => {
              console.error('修改失败:', error);
              this.$message({
                message: '课程修改失败，请重试！',
                type: 'error',
                duration: 2000,
              });
            });
        } else {
          console.log('表单验证失败');
        }
      });
    },
    // 删除课程
    deleteForm() {
      deleteCourse(this.form.id)
        .then(response => {
          console.log('msg:', response.msg);
          console.log('error:', response.error);
          if (response.msg == "ok") {
            this.$message({
              message: '课程删除成功！',
              type: 'success',
              duration: 2000,
            });
            this.dialogVisible = false; // 关闭弹窗
            this.resetForm();          // 重置表单
            return this.fetchCourse();
          }
        })
        .catch(error => {
          console.error('删除失败:', error);
          this.$message({
            message: '课程删除失败，请重试！',
            type: 'error',
            duration: 2000,
          });
        });
    },
    // 获取初始表单数据
    getInitialFormData() {
      return {
        courseName: '',
        startClass: 1,
        endClass: 1,
        startWeek: 1,
        endWeek: 1,
        dayOfWeek: 1,
      };
    },
    // 重置表单
    resetForm() {
      if (this.$refs.formRef) {
        this.$refs.formRef.resetFields(); // 重置表单数据和验证状态
        this.isEdit = false;
        console.log("Form has been reset.");
      }
    },
    // 验证结束周数
    validateEndWeek(rule, value, callback) {
      const startWeek = this.form.startWeek;
      if (value < startWeek) {
        console.log("error");
        callback(new Error('结束周数必须大于等于起始周数'));
      } else {
        console.log("access")
        callback();
      }
    },
    // 验证下课节数
    validateEndClass(rule, value, callback) {
      const startClass = this.form.startClass;
      if (value < startClass) {
        console.log("error");
        callback(new Error('下课节数必须大于等于上课节数'));
      } else {
        console.log("access")
        callback();
      }
    },
    // 获取课程
    fetchCourse() {
      getCourse(this.currentWeekNo)
        .then(response => {
          console.log('获取课程成功!', response.data);
          if (response.data) {
            if(!response.data.course_list){
              this.course = [];
            } else{
              this.course = this.convertKeysToCamel(response.data.course_list);
            }
          } else {
            console.error('响应数据格式不正确:', response.data.course_list);
            this.$message({
              message: '获取课程失败，数据格式错误！',
              type: 'error',
              duration: 2000,
            });
          }
        })
        .catch(error => {
          console.error('获取课程失败:', error);
          this.$message({
            message: '获取课程失败，请重试！',
            type: 'error',
            duration: 2000,
          });
        });
    },
    // 课程名替换
    snakeToCamel(str) {
      return str.replace(/_([a-z])/g, (match, p1) => p1.toUpperCase());
    },
    convertKeysToCamel(obj) {
      if (Array.isArray(obj)) {
        return obj.map(item => this.convertKeysToCamel(item));
      } else if (obj !== null && typeof obj === 'object') {
        return Object.keys(obj).reduce((acc, key) => {
          const camelKey = this.snakeToCamel(key);
          acc[camelKey] = this.convertKeysToCamel(obj[key]);
          return acc;
        }, {});
      }
      return obj;
    },
  },
  created() {
    // 定义验证规则
    this.rules = {
      courseName: [
        { required: true, message: '课程名称不能为空！', trigger: 'blur' }
      ],
      startClass: [
        { required: true, message: '课程上课节数不能为空！', trigger: 'blur' }
      ],
      endClass: [
        { required: true, message: '课程下课节数不能为空！', trigger: 'blur' },
        { validator: this.validateEndClass, trigger: 'change' }
      ],
      startWeek: [
        { required: true, message: '课程起始周数不能为空！', trigger: 'blur' }
      ],
      endWeek: [
        { required: true, message: '课程结束周数不能为空！', trigger: 'blur' },
        { validator: this.validateEndWeek, trigger: 'change' }
      ],
      dayOfWeek: [
        { required: true, message: '周几上课不能为空！', trigger: 'blur' }
      ],
    };
    this.initDate();
    this.fetchCourse();
  },
  computed: {}
};
</script>

<style lang='scss' scoped>
.container-schedule {
  height: 100%;
  width: 100%;
  overflow-y:auto;
  overflow-x: hidden; 
}

.index {
  width: 100%;
  box-sizing: border-box;
  padding: 10px;
  background-color: #fff;
}

table {
  width: 100%;
  background-color: #fff;
  border-color: #e3e8ee;
  text-align: center;
  border-radius: 4px;
  border-top: 1px solid #e3e8ee;
  border-left: 1px solid #e3e8ee;
  table-layout: fixed;
}

tr {
  border: none;
  text-align: center;
}

tr:nth-child(1) {
  height: 100%;
  background-color: #f3f3f3;
  border: none;
  color: #525252;
  font-weight: bold;
}

tr:first-child {
  width: 50px;
  height: 50px;
}

td .sepacialy {
  border-top: 1px solid #e3e8ee;
  border-left: 1px solid #e3e8ee;
  height: 50px;
  width: 100px;
}

.table-heard {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  box-sizing: border-box;
  padding: 10px 10px;

  .center {
    display: flex;
    justify-content: center;
    align-items: center;

    p {
      font-size: 20px;
    }
  }
}

.course-item {
  display: flex;
  justify-content: center;
  flex-direction: column;
  gap: 6px;
  align-items: center;
  background: white;
  min-height: 80px;
  border-radius: 4px;

  box-sizing: border-box;
  padding: 10px 0;

  p {
    color: black;
  }
}

.time {
  background: #f7f7f7;
  color: #707070;
}

.timeNo {
  font-weight: bold;
}

.class-block {
  border-radius: 5px;
  border: 2px solid #707070;
}

p {
  padding: 10px;
  margin: 0;
}

.coursename {
  font-weight: bolder;
}

td {
  position: relative;
  /* 使 td 可以定位子元素 */
  padding: 0;
  /* 移除内边距，确保 div 能填充满整个 td */
  height: 100%;
  /* 让 td 高度充满其容器 */
  border-bottom: 1px solid #e3e8ee;
  border-right: 1px solid #e3e8ee;
  text-align: center;
  box-sizing: border-box;
}

/* 让 outblock 填满 td */
.outblock {
  display: block;
  /* 确保它作为块级元素显示 */
  width: 100%;
  /* 填满整个 td 宽度 */
  height: 100%;
  /* 填满整个 td 高度 */
  position: absolute;
  /* 绝对定位，使它充满 td */
  top: 0;
  left: 0;
}

/* 内部 div 设置样式 */
.outblock>div {
  width: 100%;
  /* 使 div 的宽度填满 */
  height: 100%;
  /* 使 div 的高度填满 */
  padding: 10px;
  /* 给内部 div 添加一些内边距 */
  border-radius: 4px;
  /* 边角圆滑 */
  box-sizing: border-box;
  /* 确保内边距不影响宽高 */
  background-color: #f7f7f7;
  /* 示例背景色 */
}

.delete-button {
  background-color: red;
  color: white;
}
</style>
