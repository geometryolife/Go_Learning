// 练习3.8: 生成高度放大的分形需要极高的数学精度。分别用以下4种类型
// (complex64、complex128、big.Float 和 big.Rat)表示数据实现同一个分
// 形(后面两种类型由 math/big 包给出。big.Float 类型随意选用
// float32/float64 浮点数，但精度有限；big.Rat 类型使用无限精度的有
// 理数。)它们在计算性能和内存消耗上相比如何? 放大到什么程度，渲染的
// 失真变得可见?
