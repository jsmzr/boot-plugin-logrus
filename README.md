# boot-plugin-logrus

[![Build Status](https://github.com/jsmzr/boot-plugin-logrus/workflows/Run%20Tests/badge.svg?branch=main)](https://github.com/jsmzr/boot-plugin-logrus/actions?query=branch%3Amain)
[![codecov](https://codecov.io/gh/jsmzr/boot-plugin-logrus/branch/main/graph/badge.svg?token=HNQCAN3UVR)](https://codecov.io/gh/jsmzr/boot-plugin-logrus)

logrus 插件

## 使用说明

1. 拉取依赖，`go get -u github.com/jsmzr/boot-plugin-logurs`
2. 在 `main.go` 中引入，`import _ "github.com/jsmzr/boot-plugin-logrus"`
3. 在需要使用的地方引入，`import log "github.com/jsmzr/boot-log"`，并使用 `log.Info("hello world")`

当前暂只支持配置日志等级和格式化，输出固定使用 stdout

```yaml
boot:
  logging:
    level: info
    format:
      timestampFormat: 2006-01-02T15:04:05Z07:00
```