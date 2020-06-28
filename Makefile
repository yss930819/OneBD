# 查看版本
version:
	@awk -F '"' '/Version/ {print $$2;}' onebd.go

# 提交版本
tag:
	@awk -F '"' '/Version/ {print $$2;system("git tag "$$2);system("git push origin "$$2)}' onebd.go

# 删除远端版本 慎用
dropTag:
	@awk -F '"' '/Version/ {print $$2;system("git tag -d "$$2);system("git push origin :refs/tags/"$$2)}' onebd.go

