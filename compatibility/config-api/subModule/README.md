# 通过配置API生成实例化子模块

## 概述

subModule 目录展示了如何使用配置API独立获取和使用 Dubbo-Go 框架的特定组件实例，而无需启动完整的框架服务。这种方式提供了更高的灵活性和模块化能力。

## 目录结构

### registry/
使用配置API生成注册中心实例，展示如何独立使用注册中心进行服务注册和发现。

**示例功能：**
- 创建注册中心配置
- 获取注册中心实例
- 注册和反注册服务URL
- 独立使用注册中心功能

**适用场景：**
- 需要独立使用注册中心的项目
- 微服务架构中的服务发现
- 注册中心功能的测试和验证

### configcenter/
使用配置API生成配置中心实例，展示如何独立使用配置中心进行配置管理。

**示例功能：**
- 创建配置中心配置
- 获取配置中心实例
- 发布、获取、监听、删除配置
- 独立使用配置中心功能

**适用场景：**
- 需要独立使用配置中心的项目
- 配置管理的测试和验证
- 配置中心的集成开发

### service/
使用配置API生成服务实例，展示如何独立使用服务组件。

**示例功能：**
- 创建服务配置
- 获取服务实例
- 独立使用服务功能

**适用场景：**
- 需要独立使用服务组件的项目
- 服务功能的测试和验证
- 服务组件的集成开发

## 使用方式

### 基本模式
```go
// 1. 创建组件配置
componentConfig := config.NewComponentConfigBuilder().
    SetProtocol("protocol").
    SetAddress("address").
    Build()

// 2. 获取组件实例
instance, err := componentConfig.GetInstance()
if err != nil {
    panic(err)
}

// 3. 使用组件功能
// ... 具体的组件操作
```

## 相关链接

- [配置API使用指南](../README.md)
- [RPC服务示例](../rpc/README.md)
- [配置中心示例](../configcenter/README.md)
- [Dubbo-Go 官方文档](https://dubbo.apache.org/zh-cn/overview/mannual/golang-sdk/)

