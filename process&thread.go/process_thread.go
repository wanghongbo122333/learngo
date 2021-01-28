import	os
print("当前进程%s启动中。。。",os.getpid())
pid = os.fork()
if pid == 0:
