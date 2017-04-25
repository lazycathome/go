package redisCliPool

// redis strings类型的命令定义
const (
	// K_GET 命令
	GET    string = "GET"
	SET    string = "SET"
	SETEX  string = "SETEX"
	SETNX  string = "SETNX" //只有当该键不存在时，用来实现分布式锁挺好
	INCR   string = "INCR"
	DECR   string = "DECR"
	DECRBY string = "DECRBY"
	MGET   string = "MGET"
	MSET   string = "MSET"//
	MSETNX string = "MSETNX" //设置多个key value，仅当key存在时
	PSETEX string = "PSETEX" //设置已毫秒为过期时间的key value
)

// KEYS 类型的命令定义
const (
	DEL       string = "DEL"       //删除一个key
	EXISTS    string = "EXISTS"    //查询一个key是否存在
	EXPIRE    string = "EXPIRE"    //设置一个key的过期秒数
	KEYS      string = "KEYS"      //查找所有匹配给定模式的键
	TTL       string = "TTL"       //获取过期时间，以s为单位
	SORT      string = "SORT"      //对队列、集合、有序集合排序
	RANDOMKEY string = "RANDOMKEY" //返回一个随机的key
	TYPE      string = "TYPE"      //返回key的存储类型
)

// LISTS 类型的命令定义
const (
	BLPOP      string = "BLPOP"      //删除，并或得该列表的第一个元素，或阻塞，知道有一个可用
	BRPOP      string = "BRPOP"      //删除、并或得该列表中的最后一个元素，或阻塞，知道有一个可用
	BRPOPLPUSH string = "BRPOPLPUSH" //弹出一个列表的值，将它推到另一个列表，并返回它或阻塞，知道有一个可用
	LINDEX     string = "LINDEX"     //获取一个元素，通过其索引列表
	LINSERT    string = "LINSERT"    //在列表中的另一个元素之前或之后插入一个元素
	LLEN       string = "LLEN"       //获得队列的长度
	LPUSH      string = "LPUSH"      //从队列的左边入队一个或多个元素
	LPOP       string = "LPOP"       //从队列的左边出队一个元素
	LRANGE     string = "LRANGE"     //从队列中获取指定返回的元素
	LREM       string = "LREM"       //从队列中删除元素
	LSET       string = "LSET"       //设置队列里面一个元素的值
	LTRIM      string = "LTRIM"      //修建到指定范围内的清单
	LPUSHX     string = "LPUSHX"     //当队列存在时，从队列左边入队一个元素
	RPOP       string = "RPOP"       //从队列的邮编出队一个元素
	RPOPLPUSH  string = "RPOPLPUSH"  //删除队列我中的最后一个元素，将其追缴到另一个队列
	RPUSH      string = "RPUSH"      //从队列的邮编入队一个元素
	RPUSHX     string = "RPUSHX"     //从队列的有队入队一个元素，仅队列存在时有效
)

// hashes 类型 命令集合
const (
	HDEL         string = "HDEL"         //删除一个或多个hash的field
	HEXISTS      string = "HEXISTS"      //判断field是否存在于hash中
	HGET         string = "HGET"         //获取hash中的field的值
	HGETALL      string = "HGETALL"      //从hash中读取全部的域和值
	HINCRBY      string = "HINCRBY"      //将hash中指定域的值增加给定的数字
	HINCRBYFLOAT string = "HINCRBYFLOAT" //将hash中指定域的值增加给定的浮点数
	HKEYS        string = "HKEYS"        //获取hash的所有字段
	HELN         string = "HELN"         //获取hash里所有字段的数量
	HMGET        string = "HMGET"        //设置hash字段值
	HMSET        string = "HMSET"        //设置hash字段值
	HSET         string = "HSET"         //设置hash里面一个字段的值
	HSETNX       string = "HSETNX"       //设置hash的一个字段，只有当这个字段不存在时有效
	HSTRLEN      string = "HSTRLEN"      //获取hash里面指定field的长度
	HVALS        string = "HVALS"        //获得hash的所有值
	HSCAN        string = "HSCAN"        //迭代hash里面的元素
)

//sets 类型命令集合
const (
	SADD        string = "SADD"        //添加一个或者多个元素到集合set里
	SCARD       string = "SCARD"       //获取集合里面的元素数量
	SDIFF       string = "SDIFF"       //或得队列不存在的元素
	SDIFFSTORE  string = "SDIFFSTORE"  //获得队列不存在的元素，并存储在一个关键的结果集
	SINTER      string = "SINTER"      //获得两个集合的交集
	SINTERSTORE string = "SINTERSTORE" //或得两个集合的交集，并存储在一个关键的结果集
	SISMEMBER   string = "SISMEMBER"   //确定一个给定的值是一个集合的成员
	SMOVE       string = "SMOVE"       //移动集合里面的一个key到另一个集合
	SPOP        string = "SPOP"        //删除并获取一个集合里面的元素
	SRANDMEMBER string = "SRANDMEMBER" //从集合里面随机获取一个key
	SRREM       string = "SRREM"       //从集合里删除一个或多个key
	SUNION      string = "SUNION"      //添加多个set元素
	SUNIONSTORE string = "SUNIONSTORE" //合并set元素。并将结果存入新的set里面
	SSCAN       string = "SSCAN"       //迭代set里面的元素
)

//connection 连接命令定义
const (
	AUTH   string = "AUTH"   //验证服务器命令
	ECHO   string = "ECHO"   //回显输入的字符串
	PING   string = "PING"   //ping服务器
	QUIT   string = "QUIT"   //关闭连接，退出
	SELECT string = "SELECT" //选择新数据库
)

//transactions 事务命令集合
const (
	DISCARD string = "DISCARD"
	EXEC    string = "EXEC"
	MULTI   string = "NULTI"
	UNWATCH string = "UNWATCH"
	WATCH   string = "WATCH"
)
