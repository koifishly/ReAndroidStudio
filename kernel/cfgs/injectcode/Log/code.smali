# 设置新的标签
const-string v0, "newTag"
invoke-static {v0}, Landroid/koifishly/log/Logger;->SetTag(Ljava/lang/String;)V

# 打印日志
const-string v0, "content"
invoke-static {v0}, Landroid/koifishly/log/Logger;->LogString(Ljava/lang/String;)V