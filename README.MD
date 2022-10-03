## golang study notes

> The master has failed more times than the beginner has tried.

### go 语言的诞生背景

1. 多核硬件架构
2. 超大规模分布式计算集群
3. web 模式下的新开发规模和迭代速度

---

### 创始人

1. Rob Pike
2. Ken Thompson
3. Robert Griesemer

### 特点

1. 简单(关键字很少)
2. 高效(gc,指针,原生支持并发)
3. 云计算语言
4. 区块链语言

---

### main 函数

1. package 必须为 main 包
2. func 必须为 main 函数
3. main 函数无参数无返回值
4. 通过`os.Exit`返回状态
5. 通过`os.Args`获取命令行参数

---

### 类型转换

1. go 语言不允许隐式类型转换
2. 自定义类型别名也不能隐式转换

---

### 字符串

1. 字符串是数值类型
2. 字符串默认值为空字符,不是**nil**

---

### 切片

简单可以理解为可变长度的数组

1. len 表示切片内可访问的元素长度
2. cap 表示切片的可拓展容量,`cap = cap \* 2`

![切片的数组结构](./src/assets/images/slice.jpg)

---
