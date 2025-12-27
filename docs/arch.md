

## **Recommendation: Add an IR Transformation Pass**

Your compiler already has a clean pipeline:
```
Parser → IRVisitor (IR Builder) → Codegen → Object File
```

For async/await, you should add a **transformation pass** between IR building and codegen:

```
Parser → IRVisitor → Async Transform → Codegen → Object File
```