## protocol buffer



![1538304025110](assets/1538304025110-1548584857308.png)

### 简介

**Google Protocol Buffer** (简称 Protobuf)是google旗下的一款轻便高效的结构化数据存储格式，平台无关、语言无关、可扩展，可用于通讯协议和数据存储等领域。所以很适合用做数据存储和作为不同应用，不同语言之间相互通信的数据交换格式，只要实现相同的协议格式即同一 proto文件被编译成不同的语言版本，加入到各自的工程中去。这样不同语言就可以解析其他语言通过 protobuf序列化的数据。目前官网提供了 C++,Python,JAVA,GO等语言的支持。google在2008年7月7号将其作为开源项目对外公布。

### **tips：**

1. 啥叫平台无关？Linux、mac和Windows都可以用，32位系统，64位系统通吃
2. 啥叫语言无关？C++、Java、Python、Golang语言编写的程序都可以用，而且可以相互通信
3. 那啥叫可扩展呢？就是这个数据格式可以方便的增删一部分字段啦~
4. 最后，啥叫序列化啊？解释得通俗点儿就是把复杂的结构体数据按照一定的规则编码成一个字节切片

### 数据交换格式

常用的数据交换格式有三种：

1. json: 一般的web项目中，最流行的主要还是 json。因为浏览器对于json 数据支持非常好，有很多内建的函数支持。 
2. xml: 在 webservice 中应用最为广泛，但是相比于 json，它的数据更加冗余，因为需要成对的闭合标签。json 使用了键值对的方式，不仅压缩了一定的数据空间，同时也具有可读性。
3. protobuf: 是后起之秀，是谷歌开源的一种数据格式，适合高性能，对响应速度有要求的数据传输场景。因为 profobuf 是二进制数据格式，需要编码和解码。数据本身不具有可读性。因此只能反序列化之后得到真正可读的数据。

#### protobuf的优势与劣势

#### 优势： 

1：序列化后体积相比Json和XML很小，适合网络传输 

2：支持跨平台多语言 

3：消息格式升级和兼容性还不错 

4：序列化反序列化速度很快，快于Json的处理速度

#### 劣势：

1：应用不够广(相比xml和json)

2：二进制格式导致可读性差

3：缺乏自描述

### protobuf环境安装

1. protobuf 编译工具安装

    ```shell
    1、下载 protoBuf：
    git clone https://github.com/protocolbuffers/protobuf.git
    2、或者直接将压缩包拖入后解压
    unzip protobuf.zip 
    3、安装依赖库
    sudo apt-get install autoconf  automake  libtool curl make  g++  unzip libffi-dev -y
    4、进入目录
    cd protobuf/ 
    5、自动生成configure配置文件：
    ./autogen.sh 
    6、配置环境：
    ./configure
    7、编译源代码(时间比较长)：
    make 
    8、安装
    sudo make install
    9、刷新共享库 （很重要的一步啊）
    sudo ldconfig 
    10、成功后需要使用命令测试
    protoc -h  
    ```

2. protobuf 的 go 语言插件安装

   由于protobuf并没直接支持go语言需要我们手动安装相关插件

   ```shell
   1获取 proto包
   Go语言的proto API接口
   go get  -v -u github.com/golang/protobuf/proto
   go get  -v -u github.com/golang/protobuf/protoc-gen-go
   
   2编译
   cd $GOPATH/src/github.com/golang/protobuf/protoc-gen-go/
   go build
   3将生成的 protoc-gen-go可执行文件，放在/bin目录下
   sudo cp protoc-gen-go /bin/
   ```

### protobuf语法

protobuf 通常会把用户定义的结构体类型叫做一个消息，这里我们遵循惯例，统一称为消息。protobuf 消息的定义（或者称为描述）通常都写在一个以 .proto 结尾的文件中。

#### 消息类型

```protobuf
syntax = "proto3"; 						//指定版本信息，不指定会报错
package pb;						//后期生成go文件的包名
//message为关键字，作用为定义一种消息类型
message Person{
	//    名字
    string name = 1;
	//    年龄
    int32  age = 2 ;
	//    邮箱
    repeated string emalis =3;
	//    手机
    repeated string phones =4;
    // repeated为关键字，作用为重复使用 一般在go语言中用切片表示
}
```



消息格式说明

消息由至少一个字段组合而成，类似于Go语言中的结构体，每个字段都有一定的格式：

```
//注释格式 注释尽量也写在内容上方
（字段修饰符）数据类型 字段名称 = 唯一的编号标签值;

```

- 字段名称：protobuf建议以下划线命名而非驼峰式
- 唯一的编号标签：代表每个字段的一个唯一的编号标签，在同一个消息里不可以重复。这些编号标签用与在消息二进制格式中标识你的字段，并且消息一旦定义就不能更改。需要说明的是标签在1到15范围的采用一个字节进行编码，所以通常将标签1到15用于频繁发生的消息字段。编号标签大小的范围是1到229
- 注释格式：向.proto文件添加注释，可以使用C/C++/java/Go风格的双斜杠（//） 语法格式

#### 数据类型 

| .proto类型 | Go类型  | 介绍                                                         |
| ---------- | ------- | ------------------------------------------------------------ |
| double     | float64 | 64位浮点数                                                   |
| float      | float32 | 32位浮点数                                                   |
| int32      | int32   | 使用可变长度编码。编码负数效率低下——如果你的字段可能有负值，请改用sint32。 |
| int64      | int64   | 使用可变长度编码。编码负数效率低下——如果你的字段可能有负值，请改用sint64。 |
| uint32     | uint32  | 使用可变长度编码。                                           |
| uint64     | uint64  | 使用可变长度编码。                                           |
| sint32     | int32   | 使用可变长度编码。符号整型值。这些比常规int32s编码负数更有效。 |
| sint64     | int64   | 使用可变长度编码。符号整型值。这些比常规int64s编码负数更有效。 |
| fixed32    | uint32  | 总是四字节。如果值通常大于228，则比uint 32更有效             |
| fixed64    | uint64  | 总是八字节。如果值通常大于256，则比uint64更有效              |
| sfixed32   | int32   | 总是四字节。                                                 |
| sfixed64   | int64   | 总是八字节。                                                 |
| bool       | bool    | 布尔类型                                                     |
| string     | string  | 字符串必须始终包含UTF - 8编码或7位ASCII文本                  |
| bytes      | []byte  | 可以包含任意字节序列                                         |

 

更多详情请看：<https://developers.google.com/protocol-buffers/docs/encoding>

#### 结构体嵌套

```protobuf
syntax = "proto3"; 						//指定版本信息，不指定会报错
package pb;						//后期生成go文件的包名
//message为关键字，作用为定义一种消息类型
message Person{
	//    名字
    string name = 1;
	//    年龄
    int32  age = 2 ;
	//    邮箱
    repeated string emali =3;
	//    手机
    repeated string PhoneNumber =4;
    // repeated为关键字，作用为重复使用 一般在go语言中用切片表示
}

//message为关键字，作用为定义一种消息类型可以被另外的消息类型嵌套使用
message PhoneNumber {
    string number = 1;
    int64 type = 2;
}
```







#### 枚举类型

```protobuf
syntax = "proto3"; 						//指定版本信息，不指定会报错
package pb;						//后期生成go文件的包名

//message为关键字，作用为定义一种消息类型
message Person {
	string	name = 1;					//姓名
    int32	age = 2;					//年龄
	repeated string emails = 3; 		//电子邮件（repeated表示字段允许重复）
	repeated PhoneNumber phones = 4;	//手机号
}

//enum为关键字，作用为定义一种枚举类型
enum PhoneType {
	MOBILE = 0;
    HOME = 1;
    WORK = 2;
}

//message为关键字，作用为定义一种消息类型可以被另外的消息类型嵌套使用
message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
}
```









#### 默认缺省值

当一个消息被解析的时候，如果被编码的信息不包含一个特定的元素，被解析的对象锁对应的域被设置位一个默认值，对于不同类型指定如下：

- 对于strings，默认是一个空string
- 对于bytes，默认是一个空的bytes
- 对于bools，默认是false
- 对于数值类型，默认是0



### 基本编译

​	可以通过定义好的.proto文件来生成**go**,Java,Python,C++, Ruby, JavaNano, Objective-C,或者C# 代码，需要基于.proto文件运行protocolbuffer编译器protoc。

通过如下方式调用protocol编译器：

```shell
 protoc --proto_path=IMPORT_PATH --go_out=DST_DIR path/to/file.proto
```

其中：

1. --proto_path=IMPORT_PATH，IMPORT_PATH指定了 .proto 文件导包时的路径，如果忽略则默认当前目录。如果有多个目录则可以多次调用--proto_path，它们将会顺序的被访问并执行导入。
2. --go_out=DST_DIR， 指定了生成的go语言代码文件放入的文件夹
3. 允许使用 `protoc --go_out=./   *.proto` 的方式一次性编译多个 .proto 文件
4. 编译时，protobuf 编译器会把 .proto 文件编译成 .pd.go 文件



我们可以通过以下命令对刚写好的proto文件进行编译

```
protoc --go_out=./ *.proto
```

### 编译的时候发生了什么?

当用protocol buffer编译器来运行.proto文件时，编译器将生成所选择语言的代码，这些代码可以操作在.proto文件中定义的消息类型，包括获取、设置字段值，将消息序列化到一个输出流中，以及从一个输入流中解析消息。

 

​       对go来说，编译器会为每个消息类型生成了一个.pd.go文件。



### 利用protobuf生成的类来编码

```go
package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"protocolbuffer_excise/pb"
)

func main() {
	person := &pb.Person{
		Name:   "XiaoYuer",
		Age:    16,
		Emails: []string{"xiao_yu_er@sina.com", "yu_er@sina.cn"},
		Phones: []*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number: "13113111311",
				Type:   pb.PhoneType_MOBILE,
			},
			&pb.PhoneNumber{
				Number: "14141444144",
				Type:   pb.PhoneType_HOME,
			},
			&pb.PhoneNumber{
				Number: "19191919191",
				Type:   pb.PhoneType_WORK,
			},
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		fmt.Println("marshal err:", err)
	}

	newdata := &pb.Person{}
	err = proto.Unmarshal(data, newdata)
	if err != nil {
		fmt.Println("unmarshal err:", err)
	}

	fmt.Println(newdata)

}
```



