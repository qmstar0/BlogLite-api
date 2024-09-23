# 2024-9-23

## 内容更新

1. 重写领域层。领域层按照事件风暴分析结果进行分包，不再模糊不同领域之间的界限
2. 重新引入事件驱动模块。之前的事件驱动只能算是预留了事件驱动模块，并没有任何实际应用，这次更新引入的完整的事件驱动设计，随之而来的还有诸多困难和挑战，以及更多可能性
3. 重新划分包与模块。主要参考了这个项目[wild-workouts-go-ddd-example](https://github.com/ThreeDotsLabs/wild-workouts-go-ddd-example/)
4. 移除了第三方包`cobra`，相关功能改用golang标准库。
5. 优化config内容。
