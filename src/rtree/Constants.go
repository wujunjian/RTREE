package rtree


var  MAX_NUMBER_OF_ENTRIES_IN_NODE = 20// 结点中的最大条目数
var  MIN_NUMBER_OF_ENTRIES_IN_NODE = 8// 结点中的最小条目数

var  RTDataNode_Dimension = 2

/** Available RTree variants. */ // 树的类型常量
var  RTREE_LINEAR = 0 // 线性
var  RTREE_QUADRATIC = 1 // 二维
var  RTREE_EXPONENTIAL = 2 // 多维
var  RSTAR = 3 // 星型

var  NIL = -1