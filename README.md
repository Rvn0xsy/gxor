# gxor

**gxor** 程序根据输入的二进制文件进行异或运算输出

```
[gxor] Xor Binary file

Usage : gxor -input-file payload.bin -output-file out.bin -xor-key 10 -appendFlag A
```


## 使用帮助

```
Usage of ./gxor:
  -append-flag string
    	flag for file head, 1 bytes (default "A")
  -input-file string
    	input bin file
  -output-file string
    	input bin file (default "output.bin")
  -xor-key int
    	input xor key (default 8)
```

- append-flag 主要是为了防止直接异或过的shellcode代码被查杀，因此可以在文件头添加一些字符串


