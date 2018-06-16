## Switch evaluation order

Switch cases evaluate values from top to bottom, stopping when a case succeeds.

(for example,

```
  switch i {
    case 0:
    case f():
  }
```

  does not call `f` if `i==0`)
