## drone-feishu
推送DroneCI的流水线构建信息到飞书

## 使用示例

```yaml
kind: pipeline
type: docker
name: build
steps:
  - name: build
    image: golang:1.21-alpine
    commands:
      - GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-feishu .
  - name: notification
    image: cruii/drone-feishu
    pull: always
    settings:
      user_id:
        from_secret: user_id
      chat_id:
        from_secret: chat_id
      app_id:
        from_secret: app_id
      app_secret:
        from_secret: app_secret
    when:
      status:
        - success
        - failure
trigger:
  event:
    include:
      - push
      - pull_request
  branch:
    - main
    - dev
```
