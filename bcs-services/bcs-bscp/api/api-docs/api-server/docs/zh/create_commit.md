### 描述
该接口提供版本：v1.0.0+
 

创建提交记录，用于记录每一次配置文件的更改提交记录。在创建提交之前，需上传配置项元数据内容信息(upload_content)和创建配置项元数据(create_content)。

将指定的content_id配置项的元数据信息绑定到指定的config_item_id代表的配置项上，整体做为一个配置项的一次commit。

### 输入参数
| 参数名称     | 参数类型     | 必选   | 描述             |
| ------------ | ------------ | ------ | ---------------- |
| biz_id         | uint32       | 是     | 业务ID     |
| app_id         | uint32       | 是     | 应用ID     |
| config_item_id         | uint32       | 是     | 配置项ID     |
| content_id         | uint32       | 是     | 配置项元数据ID     |
| memo         | string       | 否     | 备注。最大长度256个字符，仅允许使用中文、英文、数字、下划线、中划线、空格，且必须以中文、英文、数字开头和结尾    | 

### 调用示例
```json
{
    "content_id": 1,
    "memo": "commit"
}
```

### 响应示例
```json
{
    "code": 0,
    "message": "ok",
    "data": {
        "id": 1
    }
}
```

### 响应参数说明
| 参数名称     | 参数类型   | 描述                           |
| ------------ | ---------- | ------------------------------ |
|      code        |      int32      |            错误码                   |
|      message        |      string      |             请求信息                  |
|       data       |      object      |            响应数据                  |

#### data
| 参数名称     | 参数类型   | 描述                           |
| ------------ | ---------- | ------------------------------ |
|      id        |      uint32      |            提交记录ID                    |