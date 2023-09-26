```rust
$ ./main
do
dropping
end
```

结论：和c++类似，这种临时对象，不会在块结束时销毁，而是在本行表达式结束时销毁。

