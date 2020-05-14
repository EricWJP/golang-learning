表4-1 Unix风格的ps命令参数 9
-A  显示所有进程
-N  显示与指定参数不符的所有进程
-a  显示除控制进程(session leader1)和无终端进程外的所有进程 
-d  显示除控制进程外的所有进程
-e  显示所有进程
-C cmdlist 显示包含在cmdlist列表中的进程 
-G grplist 显示组ID在grplist列表中的进程 
-U userlist	显示属主的用户ID在userlist列表中的进程 
-g grplist 显示会话或组ID在grplist列表中的进程
-p pidlist 显示PID在pidlist列表中的进程 
-s sesslist	显示会话ID在sesslist列表中的进程
-t ttylist	显示终端ID在ttylist列表中的进程
-u userlist 显示终端ID在ttylist列表中的进程
-F
-O format -M
-c
-f
-j
-l
-o format -y
-Z
-H
-n namelist -w
-L
-V
   显示更多额外输出(相对-f参数而言) 显示默认的输出列以及format列表指定的特定列 显示进程的安全信息 显示进程的额外调度器信息 显示完整格式的输出
显示任务信息
显示长列表
仅显示由format指定的列
不要显示进程标记(process flag，表明进程状态的标记) 显示安全标签(security context)1信息 用层级格式来显示进程(树状，用来显示父进程) 定义了WCHAN列显示的值 采用宽输出模式，不限宽度显示
显示进程中的线程 显示ps命令的版本号
  




