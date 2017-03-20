* 会自动把修改保存到db.json中，需要决定是否备份数据
* 支持rest的GET、POST、PUT、PATCH和DELETE，需要注意POST、PUT和PATCH的区别
* 需要增加头域Content-Type: application/json，否则返回200 OK但是数据不更新
* id字段只在POST增加时会设置，其他情况不可修改

