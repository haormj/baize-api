## baize-api
白泽是中国古代神话中地位崇高的神兽，祥瑞之象征，是令人逢凶化吉的吉祥之兽。白泽亦能说人话，通万物之情，晓天下万物状貌。《白泽图》曰：羊有一角当顶上，龙也，杀之震死。《三才图会》中白泽是狮子身姿，头有一角，山羊胡子。

### 概述

baize项目主要是人工智能相关，会将一些主流模型集成，并提供前端web上传图片进行验证。目前分为两个模块，[baize-api](https://github.com/haormj/baize-api)和[baize-web](https://github.com/haormj/baize-web)，目前都已开源

1. baize-api主要负责提供后端接口，集成训练好模型
    - 模型通过pytorch训练，目前训练代码未开源，后续整理好也会开源
    - 训练完成通过torchscript转化为libtorch可调用模型
    - [ccutil](https://github.com/haormj/ccutil)会对调用进行封装，并通过[swig](http://www.swig.org/)提供golang接口
    - [go-micro](https://github.com/micro/go-micro)暴露微服务
    - [micro](https://github.com/micro/micro)做http网关
2. baize-web主要负责展现前端页面，提供用户上传界面
    - 整体基于[react](https://reactjs.org/)
    - [ant design pro](https://pro.ant.design/)作为脚手架
    - [ant design](https://ant.design/)有丰富的react库

### 接口列表

|服务 | 方法 | 描述 |
|------| ------ | ---- |
|DogsVsCatsService | [GetUploadParam](#DogsVsCatsService.GetUploadParam)     |   获取上传图片参数，主要是阿里云sts返回值    |
|DogsVsCatsService | [Classify](#DogsVsCatsService.Classify)     |  进行猫狗分类|


### DogsVsCatsService.GetUploadParam

#### 接口概述

- 获取上传图片参数
- http路径：https://baize.haormj.xyz/DogsVsCatsService/GetUploadParam
- Content-Type：application/json

#### 请求参数

| 名称             | 类型   | 是否必选 | 示例值                           | 描述                                                                        |
|------------------|--------|----------|----------------------------------|-----------------------------------------------------------------------------|
|||||


#### 返回数据

| 名称    | 类型   | 示例值 | 描述                                                           |
|---------|--------|--------|----------------------------------------------------------------|
|region_id|string|oss-cn-shenzhen|主要用于前端使用sts时需传入，比阿里云region_id多了oss前缀|
|bucket|string|haormj|阿里云bucket名称|
|access_key_id|string| STS.NUap4fsGWBkAcibt****                  |访问密钥标识|
|access_key_secret|string| CSjbdvYU4AAR1yNk4cwiPasHAPendnyEepC8V**** |访问密钥|
|security_token|string| ********                                  |安全令牌|
|expiration|string| 2020-08-17T22:16:43Z                      |失效时间|


#### 示例

> request

```json
{

}
```

> response

```json
{
    "region_id":"oss-cn-shenzhen",
    "bucket":"haormj",
    "access_key_id":"STS.NUap4fsGWBkAcibt****",
    "access_key_secret":"CSjbdvYU4AAR1yNk4cwiPasHAPendnyEepC8V****",
    "security_token":"****",
    "expiration":"2020-08-17T22:16:43Z",
    "object_key":"xx/xxx/3be14af2e0d511eab1be00163e10a364"
}
```

#### 错误码

| http_code | code | code_msg | 描述 |
| --- | ---  | --- | --- |
|500|||服务器内部错误|

### DogsVsCatsService.Classify

#### 接口概述

- 获取上传图片参数
- http路径：https://baize.haormj.xyz/DogsVsCatsService/Classify
- Content-Type：application/json

#### 请求参数

| 名称             | 类型   | 是否必选 | 示例值                           | 描述                                                                        |
|------------------|--------|----------|----------------------------------|-----------------------------------------------------------------------------|
|object_key|string|是|xx/xxx/xxx|待分类图片的oss路径|


#### 返回数据

| 名称    | 类型   | 示例值 | 描述                                                           |
|---------|--------|--------|----------------------------------------------------------------|
|result|string| CAT                                             | 分类结果，比如猫或者狗       |
|image_url|string|http://haormj.oss-cn-shenzhen.aliyuncs.com/****|分类后图片地址，用于前端展示|

#### 示例

> request

```json
{
    "object_key":"xx/xxx/xxx"
}
```

> response

```json
{
    "result":"CAT",
    "image_url":"http://haormj.oss-cn-shenzhen.aliyuncs.com/****"
}
```

#### 错误码

| http_code | code | code_msg | 描述 |
| --- | ---  | --- | --- |
|500|||服务器内部错误|
|404|||未找到对应资源|

