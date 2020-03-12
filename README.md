Service系统

自定义包简介：
web： 基于HTTP通信，通过引入gin实现
init：初始化数据库等
app： 功能实现接口和具体实现，功能扩展可在此添加
api： 定义web和app调用过程中的数据结构
config: 配置文件，通过viper解析
logger: 自定义log记录
