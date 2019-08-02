## 3.  Fabric核心模块

Fabric是一个由五个核心模块组成的程序组. 在fabric在成功编译完成之后, 一共会有五个核心模块, 如下:

|    模块名称     |                     功能                     |
| :-------------: | :------------------------------------------: |
|     `peer`      | 主节点模块, 负责存储区块链数据, 运行维护链码 |
|    `orderer`    |              交易打包, 排序模块              |
|   `cryptogen`   |              组织和证书生成模块              |
|  `configtxgen`  |              区块和交易生成模块              |
| `configtxlator` |              区块和交易解析模块              |

> 五个模块中`peer`和`orderer`属于系统模块, `cryptogen`, `configtxgen`, `configtxlator`属于工具模块。工具模块负责证书文件、区块链创始块、通道创始块等相关文件和证书的生成工作，但是工具模块不参与系统的运行。peer模块和 orderer 模块作为系统模块是Fabric的核心模块，启动之后会以守护进程的方式在系统后台长期运行。
>
> Fabric的5个核心模块都是基于命令行的方式运行的，目前Fabric没有为这些模块提供相关的图形界面，因此想要熟练使用Fabric的这些核心模块，必须熟悉这些模块的命令选项。

### 3.1 cryptogen

> <font color="red">cryptogen模块主要用来生成组织结构和账号相关的文件</font>，任何Fabric系统的开发通常都是从cryptogen模块开始的。在Fabric项目中，当系统设计完成之后第一项工作就是根据系统的设计来编写cryptogen的配置文件，然后通过这些配置文件生成相关的证书文件。
>
> cryptogen模块所使用的配置文件是整个Fabric项目的基石。下面我们将介绍cryptogen模块命令行选项及其使用方式。

#### cryptogen 模块命令

> cryptogen模块是通过命令行的方式运行的，一个cryptogen命令由命令行参数和配置文件两部分组成，通过执行命令`cryptogen --help`可以显示cryptogen模块的命令行选项，执行结果如下所示：

```shell
$ cryptogen --help
usage: cryptogen [<flags>] <command> [<args> ...]
Utility for generating Hyperledger Fabric key material
Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).
Commands:
   # 显示帮助信息
  `help [<command>...]
   # 根据配置文件生成证书信息。
  `generate [<flags>]	
   # 显示系统默认的cryptogen模块配置文件信息
  `showtemplate
   # 显示当前模块版本号
  `version`
   # 扩展现有网络
  `extend [<flags>]
```

#### cryptogen 模块配置文件

> cryptogen模块的配置文件用来描述需要生成的证书文件的特性，比如：有多少个组织有多少个节点，需要多少个账号等。这里我们通过一个cryptogen模块配置文件的具体例子来初步了解配置文件的结构，该例子是Fabric源代码中自带的示例 -  crypto-config.yaml:
>
> <font color="red">这个配置文件的名字可以根据自己的意愿进行自定义, 即: xxx.yaml</font>

```yaml
OrdererOrgs:					# 排序节点的组织定义
  - Name: Orderer				# orderer节点的名称
 	Domain: example.com			# orderer节点的根域名 
 	Specs:
	    - Hostname: orderer		# orderer节点的主机名
PeerOrgs:						# peer节点的组织定义
  - Name: Org1					# 组织1的名称	1	1
	Domain: org1.example.com	# 组织1的根域名
 	EnableNodeOUs: true			# 是否支持node.js
 	Template:					
	    Count: 2				# 组织1中的节点(peer)数目
	Users:
 	    Count: 1				# 组织1中的用户数目
  - Name: Org2
    Domain: org2.example.com
    EnableNodeOUs: true
    Template:
        Count: 2
    Users:
        Count: 1
```

> 上述模板文件定义了一个orderer节点，这个orderer节点的名字为orderer，orderer节点的根域名为example.com，主机名为orderer。模板文件同时定义了两个组织，两个组织的名字分别为org1 和 org2，其中组织 org1 包含了2个节点和1个用户，组织 org2 包含2个点和1个用户。
>
> 除了Fabric源码中提供的例子，还可以通过命令`cryptogen showtemplate`获取默认的模板文件，在实际项目中稍加修改这些默认的模板文件即可使用。

#### 生成证书文件

> 在任意目录下创建一个新的目录如: MyTestNetWork， 在该目录下左如下操作： 
>
> - 编写yaml配置文件 - crypto-config.yaml
>   - 一个排序节点： orader
>     - 根域名：itcast.com
>   - 两个组织： java， go
>     - go
>       - peer节点数： 3个
>       - 用户个数： 2个
>       - 根域名：go.itcast.com
>     - java
>       - peer节点数： 3个
>       - 用户个数： 2个
>       - 根域名: java.itcast.com

- 命令

  ```shell
  # 根据默认模板在对应目录下生成证书
  $ cryptogen generate
  # 根据指定的模板在指定目录下生成证书
  $ cryptogen generate --config=./crycrypto-config.yaml --output ./crypto-config
  	--config: 指定配置文件
  	--output: 指定证书文件的存储位置, 可以不指定。会在对应路径生成目录，默认名字为：crypto-config
  ```

- Fabric证书文件结构

  - orderer节点

    ```shell
    # 查看TestNetWork目录
    itcast@ubuntu:~/TestNetWork$ tree -L 2
    .
    ├── crypto-config
    │   ├── ordererOrganizations	# orderer节点相关的证书文件
    │   └── peerOrganizations		# 组织相关的证书文件(组织的节点数, 用户数等证书文件)
    └── crypto-config.yaml			# 配置文件
    
    # 查看排序节点的证书目录, 进入到 ordererOrganizations 子目录中
    itcast@ubuntu:ordererOrganizations$ tree -L 4
    .
    └── itcast.com	# 根域名为itcast.com的orderer节点的相关证书文件
        ├── ca	# CA服务器的签名文件
        │   ├── 94db924d3be00c5adda6ac3c3cb7a5f8b80868681c3dd04b58c2920cdf56fdc7_sk
        │   └── ca.itcast.com-cert.pem
        ├── msp
        │   ├── admincerts	# orderer管理员的证书
        │   │   └── Admin@itcast.com-cert.pem
        │   ├── cacerts		# orderer根域名服务器的签名证书
        │   │   └── ca.itcast.com-cert.pem
        │   └── tlscacerts	# tls连接用的身份证书
        │       └── tlsca.itcast.com-cert.pem
        ├── orderers	# orderer节点需要的相关的证书文件
        │   └── ubuntu.itcast.com
        │       ├── msp	# orderer节点相关证书
        │       └── tls	# orderer节点和其他节点连接用的身份证书
        ├── tlsca
        │   ├── de45aeb112ee820197f7d4d475f2edbeb1705d53a690f3537dd794b66de1d6ba_sk
        │   └── tlsca.itcast.com-cert.pem
        └── users	# orderer节点用户相关的证书
            └── Admin@itcast.com
                ├── msp
                └── tls
    ```

    > 在实际开发中orderer节点这些证书其实不需要直接使用, 只是在orderer节点启动时指明项目的位置即可。

  - Peer节点

    > 进入到 peerOrganizations 子目录中,  我们详细介绍这些证书的种类和作用。由于每个组织的目录结构都是一样的， 所以我们只对其中一个组织的目录进行详细介绍。

    ```shell
    # 查看 peerOrganizations 子目录中内容
    itcast@ubuntu:peerOrganizations$ tree -L 1
    .
    ├── go.itcast.com		# go组织
    └── java.itcast.com		# java组织
    # 进入go.itcast.com 组织目录中
    itcast@ubuntu:go.itcast.com$ tree -L 4
    .
    ├── ca 	# 根节点签名证书
    │   ├── 4a367bf9e43142846e7c851830f69f72483ecb7a6def7c782278a9808bbb5fb0_sk
    │   └── ca.go.itcast.com-cert.pem
    ├── msp	
    │   ├── admincerts	# 组织管理员的证书
    │   │   └── Admin@go.itcast.com-cert.pem
    │   ├── cacerts		# 组织的根证书
    │   │   └── ca.go.itcast.com-cert.pem
    │   ├── config.yaml
    │   └── tlscacerts	# TLS连接身份证书
    │       └── tlsca.go.itcast.com-cert.pem
    ├── peers
    │   ├── peer0.go.itcast.com
    │   │   ├── msp
    │   │   │   ├── admincerts	# 组织的管理证书, 创建通道必须要有该证书
    │   │   │   ├── cacerts		# 组织根证书
    │   │   │   ├── config.yaml	
    │   │   │   ├── keystore	# 当前节点的私钥
    │   │   │   ├── signcerts	# 当前节点签名的数字证书
    │   │   │   └── tlscacerts	# tls连接的身份证书
    │   │   └── tls
    │   │       ├── ca.crt		# 组织的根证书
    │   │       ├── server.crt	# 验证本节点签名的证书
    │   │       └── server.key	# 当前节点的私钥
    │   ├── peer1.go.itcast.com
    │   │   ├── msp
    │   │   │   ├── admincerts
    │   │   │   ├── cacerts
    │   │   │   ├── config.yaml
    │   │   │   ├── keystore
    │   │   │   ├── signcerts
    │   │   │   └── tlscacerts
    │   │   └── tls
    │   │       ├── ca.crt
    │   │       ├── server.crt
    │   │       └── server.key
    │   └── peer2.go.itcast.com
    │       ├── msp
    │       │   ├── admincerts
    │       │   ├── cacerts
    │       │   ├── config.yaml
    │       │   ├── keystore
    │       │   ├── signcerts
    │       │   └── tlscacerts
    │       └── tls
    │           ├── ca.crt
    │           ├── server.crt
    │           └── server.key
    ├── tlsca
    │   ├── 3273887b1da2f27a6cbad3ac4acb0379df3d7858e0553a91fb9acb93da50b670_sk
    │   └── tlsca.go.itcast.com-cert.pem
    └── users
        ├── Admin@go.itcast.com
        │   ├── msp
        │   │   ├── admincerts	# 组织的根证书, 作为管理身份的验证
        │   │   ├── cacerts		# 用户所属组织的根证书
        │   │   ├── keystore	# 用户私钥
        │   │   ├── signcerts	# 用户的签名证书
        │   │   └── tlscacerts	# tls连接通信证书, sdk客户端使用
        │   └── tls
        │       ├── ca.crt		# 组织的根证书
        │       ├── client.crt	# 客户端身份的证书
        │       └── client.key	# 客户端的私钥
        ├── User1@go.itcast.com
        │   ├── msp
        │   │   ├── admincerts
        │   │   ├── cacerts
        │   │   ├── keystore
        │   │   ├── signcerts
        │   │   └── tlscacerts
        │   └── tls
        │       ├── ca.crt
        │       ├── client.crt
        │       └── client.key
        └── User2@go.itcast.com
            ├── msp
            │   ├── admincerts
            │   ├── cacerts
            │   ├── keystore
            │   ├── signcerts
            │   └── tlscacerts
            └── tls
                ├── ca.crt
                ├── client.crt
                └── client.key
    ```

### 3.2 configtxgen

> configtxgen 模块的功能一共有两个:
>
> - 生成 orderer 节点的初始化文件
> - 生成 channel 的初始化文件

#### configtxgen 模块命令

> configtxgen 模块是通过命令行的方式运行的，通过执行命令`configtxgen --help`可以显示 configtxgen 模块的命令行选项，执行结果如下所示：

```shell
$ configtxgen --help
Usage of ./configtxgen:
  # 指定所属的组织
  `-asOrg string`
        Performs the config generation as a particular organization (by name), only 
        including values in the write set that org (likely) has privilege to set
  # 指定创建的channel的名字, 如果没指定系统会提供一个默认的名字.
  `-channelID string`
        The channel ID to use in the configtx
  # 执行命令要加载的配置文件的路径, 不指定会在当前目录下查找
  -configPath string
        The path containing the configuration to use (if set)
  # 打印指定区块文件中的配置内容，string：查看的区块文件的名字
  -inspectBlock string
        Prints the configuration contained in the block at the specified path
  # 打印创建通道的交易的配置文件
  -inspectChannelCreateTx string
        Prints the configuration contained in the transaction at the specified path
  # 指定锚节点更新文件的路径和名字
  `-outputAnchorPeersUpdate string`
        Creates an config update to update an anchor peer (works only with the default 
        channel creation, and only for the first update)
  # 指定生成的创始区块文件的路径和名字
  `-outputBlock string`
        The path to write the genesis block to (if set)
  # 标示输出的通道文件路径和名字
  `-outputCreateChannelTx string`
        The path to write a channel creation configtx to (if set)
  #  将组织的定义打印为JSON(这对在组织中手动添加一个通道很有用)。
  -printOrg string
        Prints the definition of an organization as JSON. (useful for adding an org to
        a channel manually)
  # 指定配置文件中的节点 - configtx.yaml
  `-profile string`
        The profile from configtx.yaml to use for generation. (default
        "SampleInsecureSolo")
  # 显示版本信息
  -version
        Show version information
```

#### configtxgen模块配置文件

> configtxgen 模块的配置文件包含Fabric系统初始块、Channel初始块文件等信息。configtxgen 模块的配置文件样例如下所示，以下部分定义了整个系统的配置信息：

```yaml
Profiles:
	# 组织定义标识符，可自定义，命令中的 -profile 参数对应该标识符， 二者要保持一致
    ItcastOrgsOrdererGenesis:
        Capabilities:
            <<: *ChannelCapabilities	# 引用下面为 ChannelCapabilities 的属性
        Orderer:						# 配置属性，系统关键字，不能修改
            <<: *OrdererDefaults		# 引用下面为 OrdererDefaults 的属性
            Organizations:
                - *OrdererOrg			# 引用下面为 OrdererOrg 的属性
            Capabilities:
                <<: *OrdererCapabilities # 引用下面为 OrdererCapabilities 的属性
        Consortiums:					# 定义了系统中包含的组织
            SampleConsortium:
                Organizations:			# 系统中包含的组织
                    - *OrgGo				# 引用了下文包含的配置
                    - *OrgJava
    # 通道定义标识符，可自定义
    TwoOrgsChannel:	
        Consortium: SampleConsortium
        Application:
            <<: *ApplicationDefaults
            Organizations:
                - *OrgGo
                - *OrgJava
            Capabilities:
                <<: *ApplicationCapabilities
                
# 所有的值使用默认的true即可， 不要修改                
Capabilities:
    Global: &ChannelCapabilities
        V1_1: true
    Orderer: &OrdererCapabilities
        V1_1: true
    Application: &ApplicationCapabilities
        V1_2: true
        
# 组织节点相关配置信息
Organizations:
	# orderer节点配置信息
    - &OrdererOrg
        Name: OrdererOrg	# orderer节点名称
        ID: OrdererMSP		# orderer节点编号
        MSPDir: ./crypto-config/ordererOrganizations/itcast.com/msp	# msp文件路径
	#orderer节点中包含的组织，如果有有多个需要配置多个
    - &OrgGo
        Name: OrgGoMSP		# 组织名称
        ID: OrgGoMSP		# 组织编号
        # 组织msp文件路径
        MSPDir: ./crypto-config/peerOrganizations/go.itcast.com/msp
        AnchorPeers:		# 组织的访问域名和端口
            - Host: peer0.go.itcast.com
              Port: 7051
    - &OrgJava
        Name: OrgJavaMSP
        ID: OrgJavaMSP
        MSPDir: ./crypto-config/peerOrganizations/java.itcast.com/msp
        AnchorPeers:
            - Host: peer0.java.itcast.com
              Port: 7051
              
# orderer节点的配置信息
Orderer: &OrdererDefaults
    # orderer节点共识算法，有效值："solo" 和 "kafka"
    OrdererType: solo
    Addresses:
        - ubuntu.itcast.com:7050	# orderer节点监听的地址
    BatchTimeout: 2s
    BatchSize:
        MaxMessageCount: 10
        AbsoluteMaxBytes: 99 MB
        PreferredMaxBytes: 512 KB
	# kafka相关配置
    Kafka:
        Brokers:
            - 127.0.0.1:9092
    Organizations:
    
Application: &ApplicationDefaults
    Organizations:
```

> 上述配置文件中的 Profiles节点定义了整个系统的结构和channel的结构, 配置文件中的`Profiles`关键字不允许修改，否则配置无效。系统配置信息中设置了系统中orderer节点的信息以及系统中包含的组织数。

#### configtxgen 的使用

> 为了统一管理，我们可以将生成的初始块文件放入指定目录中，如：channel-artifacts，我们在TestNetWork目录中创建该子目录。
>
> <font color="red">configtxgen  命令在执行的时候需要加载一个叫做configtx.yaml的配置文件, 如果没有指定默认重命令执行的当前目录查找，我们可以通过参数 `-configPath`进行指定，也可以将这个目录设置到环境变量`FABRIC_CFG_PATH`中。</font>
>
> `export FABRIC_CFG_PATH=$(pwd)/networks/config/`

- 创建 `orderer` 的初始块

  ```shell
  itcast@ubuntu:TestNetWork$ configtxgen -profile ItcastOrgOrdererGenesis -outputBlock ./channel-artifacts/genesis.block
  # ItcastOrgOrdererGenesis: 要和配置文件中的配置项对应, 可以由数字和字母构成.
  # orderer初始块文件为genesis.block，生成在channel-artifacts目录中
  ```

- 创建 `channel` 的初始块

  ```shell
  itcast@ubuntu:TestNetWork$ configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./channel-artifacts/channel.tx -channelID mychannel
  # TwoOrgsChannel: 要和配置文件中的配置项对应
  # channel.tx 为生成的频道文件, 在channel-artifacts目录中
  # 创建的频道名称为: mychannel
  ```

- 创建锚点更新文件 - 每个组织分别进行更新

  ```shell
  # 更新第一个组织 OrgGoMSP 的peer节点
  itcast@ubuntu:TestNetWork$ configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/GoMSPanchors.tx -channelID mychannel -asOrg OrgGoMSP
  # TwoOrgsChannel: 要和配置文件中的配置项对应
  # OrgGoMSP组织使用的频道为 mychannel 生成的配置信息文件名为 GoMSPanchors.tx
  #==============================================================================
  # 更新第2个组织 OrgJavaMSP 的peer节点
  itcast@ubuntu:TestNetWork$ configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./channel-artifacts/JavaMSPanchors.tx -channelID mychannel -asOrg OrgJavaMSP
  # TwoOrgsChannel: 要和配置文件中的配置项对应
  # OrgJavaMSP组织使用的频道为 mychannel 生成的配置信息文件名为 JavaMSPanchors.tx
  ```

### 3.3 orderer

> orderer 模块负责对交易进行排序, 并将排序好的交易打包成区块。orderer节点的配置信息通常放在环境变量或者配置文件中，在具体操作中，如果是通过docker镜像文件的方式启动orderer，推荐使用环境变量的方式，如果是采用命令的方式直接启动，推荐将所有的信息放到配置文件中。下面将会详细介绍其用到的环境变量。
>
> orader配置文件可参考实例程序中的`orderer.yaml`
>
> `~/hyperledger-fabric/fabric-samples/config/orderer.yaml`

#### orderer模块命令

```shell
$ orderer --help
usage: orderer [<flags>] <command> [<args> ...]
Hyperledger Fabric orderer node
Flags:
  --help  Show context-sensitive help (also try --help-long and --help-man).

Commands:
  # 显示帮助信息
  `help [<command>...]
     Show help.
  # 启动orderer节点
  `start*
     Start the orderer node
  # 显示版本信息
  `version
     Show version information
  # 采用基准模式运行orderer
  `benchmark
     Run orderer in benchmark mode
```

#### orderer模块的配置信息

```shell
# orderer节点运行需要配置一些环境变量
ORDERER_GENERAL_LOGLEVEL	# 日志级别
	- `critical | error | warning | notice | info | debug
ORDERER_GENERAL_LISTENADDRESS	# orderer服务器监听的地址
ORDERER_GENERAL_LISTENPORT		# orderer服务器监听的端口
ORDERER_GENERAL_GENESTSMETHOD	# 初始化块(Genesis)的来源方式, 一般赋值为 file 即可
ORDERER_GENERAL_GENESISFILE		# 存储初始块文件的路径
ORDERER_GENERAL_LOCALMSPID		# orderer节点的编号,在configtxgen模块配置文件中指定的
	- `configtx.yaml配置文件中排序节点的组织的ID
ORDERER_GENERAL_LOCALMSPDIR		# orderer节点msp文件路径
ORDERER_GENERAL_LEDGERTYPE		# 账本类型, ram, json, file
	- `ram: 账本数据存储在内存, 一般用于测试环境
	- `json/file: 账本数据保存在文件中, 生成环境中推荐使用file 
ORDERER_GENERAL_BATCHTIMEOUT	# 批处理超时, 创建批处理之前的等待时间
	- `每隔一个BATCHTIMEOUT时长, 就会生成一个新的区块
ORDERER_GENERAL_MAXMESSAGECOUNT	# 最大消息计数, 批处理的最大消息数量
	- `只要一个区块的消息达到MAXMESSAGECOUNT指定的数量, 就会生成一个新的区块
ORDERER_GENERAL_TLS_ENABLED		# 是否启用TLS, true/false
ORDERER_GENERAL_TLS_PRIVATEKEY	# orderer节点的私钥文件, 按照下边的示例目录找
	- `crypto-config/ordererOrganizations/xx.com/orderers/orderer.xx.com/tls/server.key
ORDERER_GENERAL_TLS_CERTIFICATE	# 证书文件
	- `crypto-config/ordererOrganizations/xx.com/orderers/orderer.xx.com/tls/server.crt
ORDERER—GENERAL_TLS_ROOTCAS		# 根证书文件
	- `crypto-config/ordererOrganizations/xx.com/orderers/orderer.xx.com/tls/ca.crt
```

### 3.4 peer

> peer模块是Fabric中最重要的模块，也是在Fabric系统使用最多的模块。peer模块在Fabric中被称为主节点模块，主要负责存储区块链数据、运行维护链码、提供对外服务接口等作用。

#### 命令行和常用参数

```shell
# 通过docker启动peer节点的镜像文件, 可查看相关操作命令
$ docker run -it hyperledger/fabric-peer bash
$ peer --help
Usage:
  peer [command]

Available Commands:
  `chaincode`   相关的子命令:
  		`install`
  		`instantiate`
  		`invoke`
  		`package`
  		`query`
  		`signpackage`
  		`upgrade`
  		`list`
  channel     通道操作: create|fetch|join|list|update|signconfigtx|getinfo.
  help        查看相关命令的帮助信息
  logging     日志级别: getlevel|setlevel|revertlevels.
  node        node节点操作: start|status.
  version     当前peer的版本.

Flags:
  -h, --help                   help for peer
      --logging-level string   Default logging level and overrides, see core.yaml for full syntax
```

#### peer channel子命令

> peer channel的子命令可以通过 `peer channel --help`进行查看. 这里介绍一个这些子命令可以共用的一些参数:
>
> - `--cafile `:  当前orderer节点pem格式的tls证书文件, <font color="red">要使用绝对路径</font>.
>
>   `crypto-config/ordererOrganizations/itcast.com/orderers/ubuntu.itcast.com/msp/tlscacerts/tlsca.itcast.com-cert.pem`
>
> - `-o, --orderer`: orderer节点的地址
>
> - `--tls`: 通信时是否使用tls加密

- **create** - 创建通道

  > 命令: `peer channel create [flags]`, 可用参数为:
  >
  > - ` -c, --channelID`: 要创建的通道的ID, 必须小写, 在250个字符以内
  > - `-f, --file`: 由configtxgen 生成的通道文件, 用于提交给orderer
  > - `-t, --timeout`: 创建通道的超时时长

  ```shell
  $ peer channel create -o orderer.itcast.com:7050 -c itcastchannel -f ./channel-artifacts/channel.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/itcast.com/msp/tlscacerts/tlsca.itcast.com-cert.pem
  ```

- **join** - 将peer加入到通道中

  > 命令: `peer channel join[flags]`, 可用参数为:
  >
  > - `-b, --blockpath`: 通道文件

  ```shell
  $ peer channel join -b itcastchannel.block
  ```

- **list** - 列出peer加入的通道

  ```shell
  $ peer channel list
  ```

- **update** - 更新

  > 命令: `peer channel update [flags]`, 可用参数为:
  >
  > - ` -c, --channelID`: 要创建的通道的ID, 必须小写, 在250个字符以内
  > - `-f, --file`: 由configtxgen 生成的组织锚节点文件, 用于提交给orderer

  ```shell
  $ peer channel update -o orderer.example.com:7050 -c itcastchannel -f ./channel-artifacts/Org1MSPanchors.tx --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
  ```

#### peer chaincode 子命令

> `chaincode`一共有四个公共参数选项, 这些选项所有的子命令都可以使用, 他们分别是:
>
> - `--cafile`: PEM格式证书的位置
> - `-o, --orderer`: orderer服务器的访问地址
> - `--tls`: 使用orderer的TLS证书位置
> - `--transient`: JSON参数的编码映射
>
> chaincode命令的运行需要一些参数，这些参数可以是配置文件也可以是环境变量，由于涉及的参数并不是很多，因此大多数时候都会采用环境变量的方式来设置参数。

- **install**

  > install命令负责安装chaincode，在这个过程中如果chaincode的源代码存在语法错误，install命令会报错。install命令的选项如下所示：
  >
  > - `-c, --ctor`: JSON格式的构造参数, 默认是`"{}"`
  > - `-l, --lang`: 编写chaincode的编程语言, 默认值是 `golang`
  > - `-n, --name`: chaincode的名字
  > - `-p, --path`: chaincode源代码的名字
  > - `-v, --version`: 当前操作的chaincode的版本, 适用这些命令`install/instantiate/upgrade`

  ```shell
  $ peer chaincode install -n mycc -v 1.0 -l golang -p github.com/chaincode/chaincode_example02/go/
  # 安装成功之后, 会在peer模块的数据文件中生成一个由 -n 参数和 -v 参数组成的文件,在本例中为:
  $ docker-compose -f docker-compose-cli.yaml images
        Container                  Repository            Tag      Image Id      Size  
  ------------------------------------------------------------------------------------
  cli                      hyperledger/fabric-tools     1.2.0   379602873003   1.41 GB
  orderer.example.com      hyperledger/fabric-orderer   1.2.0   4baf7789a8ec   145 MB 
  peer0.org1.example.com   hyperledger/fabric-peer      1.2.0   82c262e65984   151 MB 
  peer0.org2.example.com   hyperledger/fabric-peer      1.2.0   82c262e65984   151 MB 
  peer1.org1.example.com   hyperledger/fabric-peer      1.2.0   82c262e65984   151 MB 
  peer1.org2.example.com   hyperledger/fabric-peer      1.2.0   82c262e65984   151 MB 
  itcast@ubuntu:first-network$ docker exec -it peer0.org1.example.com bash
  root@661a44ad6677:/opt/gopath/src/github.com/hyperledger/fabric/peer# find / -name mycc.1.0
  /var/hyperledger/production/chaincodes/mycc.1.0	# 查询到的结果 - mycc.1.0
  # 通过find命令搜索到的 mycc.1.0 文件就是chaincode打包之后的文件
  ```

- **instantiate**

  > instantiate可以对已经执行过instanll命令的Chaincode进行实例化，instantiate命令执行完成之后会启动Chaincode运行的Docker镜像，同时instantiate命令还会对Chaincode进行初始化。instantiate命令的选项如下所示：
  >
  > - `-C，--channelID`：当前命令运行的通道，默认值是`“testchainid"`。
  > - `-c, --ctor`：JSON格式的构造参数，默认值是`“{}"`
  > - `-E ， --escc` ： 应用于当前Chaincode的系统背书Chaincode的名字。
  > - `-l，--lang`：编写Chaincode的编程语言，默认值是golang
  > - `-n，--name`：Chaincode的名字。
  > - `-P，--policy`：当前Chaincode的背书策略。
  > - `-v，--version`：当前操作的Chaincode的版本，适用于`install/instantiate/upgrade`等命令
  > - `-V，--vscc`：当前Chaincode调用的验证系统Chaincode的名字。

  ```shell
  $ peer chaincode instantiate -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C itcastchannel -n mycc -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P "AND ('Org1MSP.member', 'Org2MSP.member')"
  # instantiate命令成功执行之后，可以通过docker ps命令查看己经启动的运行Chaincode的docker容器。
  ```

- **invoke**

  > invoke命令用来调用chaincode。invoke命令的选项如下所示：
  >
  > - `-C，--channelID`：当前命令运行的通道，默认值是``“testchainid"``
  > - `-c, --ctor`：JSON格式的构造参数，默认值是`“{}"`
  > - `-n，--name`：Chaincode的名字。

  ```shell
  # 调用示例
  $ peer chaincode invoke -o orderer.test.com:7050  -C testchannel -n testcc --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/test.com/orderers/orderer.test.com/msp/tlscacerts/tlsca.test.com-cert.pem --peerAddresses peer0.orgGo.test.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/orgGo.test.com/peers/peer0.orgGo.test.com/tls/ca.crt --peerAddresses peer0.orgcpp.test.com:7051 --tlsRootCertFiles /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/peerOrganizations/orgcpp.test.com/peers/peer0.orgcpp.test.com/tls/ca.crt -c '{"Args":["invoke","a","b","10"]}'
  ```

- **list**

  > list命令用来查询已经安装的Chaincode，list命令的选项如下所示：
  >
  > - `-C，--channelID`：当前命令运行的通道，默认值是`“testchainid"`
  > - `--installed`：获取当前Peer节点已经被安装的chaincode
  > - `--instantiated`：获取当前channel中已经被实例化的chaincode

  ```shell
  # 调用示例
  $ peer chaincode list --installed
  ```

- **package**

  > package用来将Chaincode打包。package命令的选项如下所示：
  >
  > - `-s，--cc-package`：对打包后的Chaincode进行签名。
  > - `-c, --ctor`：JSON格式的构造参数，默认值是`“{}"`
  > - `-i，--instantiate-policy`：Chaincode的权限
  > - `-l，--lang`：编写Chaincode的编程语言，默认值是golang
  > - `-n，--name`：Chaincode的名字。
  > - `-p，--path`：Chaincode源代码的路径。
  > - `-S，--sign`：对打包的文件用本地的MSP进行签名。
  > - `-v，--version`：当前操作的Chaincode的版本，适用于`install/instantiate/upgrade`等命令

  ```shell
  # 调用示例
  $ peer chaincode package — github.com/hyperledger/fabric/examples/chaincode/go/example  —n mycc —v 1.0 -s —S -i "OR ('Org1MSP.member'，'Org2MSP.member')" mycc.1.0.out 
  ```

- **query**

  > query命令用来执行chaincode代码中的query方法。query命令的选项如下：
  >
  > - `-C，--channelID`：当前命令运行的通道，默认值是`“testchainid"`
  > - `-c, --ctor`：JSON格式的构造参数，默认值是`“{}"`
  > - `-x，--hex`：是否对输出的内容进行编码处理
  > - `-n，--name`：Chaincode的名字。
  > - `-r，--raw`：是否输出二进制内容
  > - `-t, --tid`: 指定当前查询的编号

  ```shell
  # 调用示例
  $ peer chaincode query -C testchannel -n testcc -c '{"Args":["query","a"]}'
  ```

- **upgrade**

  > upgrade用来更新已经存在的chaincode。命令选项如下：
  >
  > - `-C，--channelID`：当前命令运行的通道，默认值是`“testchainid"`
  > - `-c, --ctor`：JSON格式的构造参数，默认值是`“{}"`
  > - `-E ， --escc` ： 应用于当前Chaincode的系统背书Chaincode的名字。
  > - `-l，--lang`：编写Chaincode的编程语言，默认值是golang
  > - `-n，--name`：Chaincode的名字。
  > - `-p, --path`: chaincode源代码的名字
  > - `-P，--policy`：当前Chaincode的背书策略。
  > - `-v，--version`：当前操作的Chaincode的版本，适用于`install/instantiate/upgrade`等命令
  > - `-V，--vscc`：当前Chaincode调用的验证系统Chaincode的名字。

  ```shell
  $ peer chaincode upgrade -o orderer.example.com:7050 -n mycc -v 1.1 -C mychannel -c '{"Args":["init","a","100","b","200"]}'
  ```

#### **peer 的环境变量**

```shell
# 配置文件和环境变量是设置peer启动参数的重要手段, 相关环境变量如下:
CORE_VM_ENDPOINT	# docker服务器的Deamon地址, 默认取端口的套接字, 如下:
	- `unix:///var/run/docker.sock
CORE_VM_DOCKER_HOSTCONFIG_NETWORKMODE	# chaincode容器的网络命名模式, 自己取名即可
	- `节点运行在同一个网络中才能相互通信, 不同网络中的节点相互隔离
CORE_PEER_PROFILE_ENABLED	# 使用peer内置的 profile server
	- `fabric的peer内置有profile server，默认时运行在6060端口上的，并且默认关闭。
CORE_LOGGING_LEVEL		# log日志的级别
	- `critical | error | warning | notice | info | debug
CORE_PEER_ID	# peer节点的编号, 自定义一个即可
CORE_PEER_GOSSIP_USELEADERELECTION	# 是否自动选举leader节点, 自动:true
CORE_PEER_GOSSIP_ORGLEADER			# 当前节点是否为leader节点, 是:true
CORE_PEER_ADDRESS					# 当前peer节点的访问地址
	- `格式: 域名:端口 / IP:端口
CORE_PEER_CHAINCODELISTENADDRESS	# chaincode的监听地址
CORE_PEER_GOSSIP_EXTERNALENDPOINT	# 节点被组织外节点感知时的地址
	- `默认为空, 代表不被其他组织节点所感知
CORE_PEER_GOSSIP_BOOTSTRAP	# 启动节点后向哪些节点发起gossip连接, 以加入网络
	- `这些节点与本地节点需要属于同一个网络
CORE_PEER_LOCALMSPID 	# peer节点所属的组织的编号, 在configtxgen.yaml中设置的
CORE_CHAINCODE_MODE		# chaincode的运行模式
	- `net: 网络模式
	- `dev: 开发模式, 该模式下可以在容器外运行chaincode
CORE_PEER_MSPCONFIGPATH	# 当前节点的msp文件路径
	- `启动peer的时候需要使用msp账号文件
	- `crypto-config/peerOrganizations/org1.x.com/peers/peer0.org1.x.com/msp
	- `创建channel的时候需要使用msp账号文件
	- `crypto-config/peerOrganizations/org1.x.com/users/Admin@org1.x.com/msp
CORE_PEER_TLS_ENABLED	# 是否激活tls, 激活:true, 不激活:false
CORE_PEER_TLS_CERT_FILE	# 服务器身份验证证书
	- `crypto-config/peerOrganizations/org1.x.com/peers/peer0.org1.x.com/tls/server.crt
CORE_PEER_TLS_KEY_FILE	# 服务器的私钥文件
	- `crypto-config/peerOrganizations/org1.x.com/peers/peer0.org1.x.com/tls/server.key
CORE_PEER_TLS_ROOTCERT_FILE	# 根服务器证书
	- `crypto-config/peerOrganizations/org1.x.com/peers/peer0.org1.x.com/tls/ca.crt
```

> <font color="red">每个 org 会选举出一个 **leader peer**（实际上可以存在多个），负责连接到 orderer。**leader peer**从orderer 拿到新块的信息后分发给其他 peer。</font>
>
> - 静态选择leader peer
>
>   ```shell
>   export CORE_PEER_GOSSIP_USELEADERELECTION=false
>   export CORE_PEER_GOSSIP_ORGLEADER=true #指定某一个peer为leader peer
>   # 1. 如果都配置为 false，那么 peer 不会尝试变成一个 leader
>   # 2. 如果都配置为 true，会引发异常
>   # 3. 静态配置的方式，需要自行保证 leader 的可用性
>   ```
>
> - 动态选择leader peer
>
>   ```shell
>   export CORE_PEER_GOSSIP_USELEADERELECTION=true
>   export CORE_PEER_GOSSIP_ORGLEADER=false
>   ```

#### **peer 默认监听的端口**

> 下面是Hyperledger中相关监听的服务端口（默认）
>
> - 7050: REST 服务端口
> - 7051：peer gRPC 服务监听端口
> - 7052：peer 代码调试模式使用的端口
> - 7053：peer 事件服务端口
> - 7054：eCAP
> - 7055：eCAA
> - 7056：tCAP
> - 7057：tCAA
> - 7058：tlsCAP
> - 7059：tlsCAA