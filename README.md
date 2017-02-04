# minivm
This is my experimental repository of writing a stack-machine based interpreter language.

Example script:
```
func fib(n)
  if n <= 1
    return 1
  end
  return fib(n - 1) + fib(n - 2)
end

n = 0
while n < 15
  print fib(n)
  n = n + 1
end
```
